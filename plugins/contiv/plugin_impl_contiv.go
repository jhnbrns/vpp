// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate protoc -I ./model/cni --gogo_out=plugins=grpc:./model/cni ./model/cni/cni.proto
//go:generate protoc -I ./model/node --gogo_out=plugins=grpc:./model/node ./model/node/node.proto

package contiv

import (
	"context"
	"fmt"
	"net"

	"strings"

	"git.fd.io/govpp.git/api"
	"github.com/apparentlymart/go-cidr/cidr"

	nodeconfig "github.com/contiv/vpp/plugins/crd/handler/nodeconfig/model"
	protoNode "github.com/contiv/vpp/plugins/ksr/model/node"

	"github.com/contiv/vpp/plugins/contiv/containeridx"
	"github.com/contiv/vpp/plugins/contiv/containeridx/model"
	"github.com/contiv/vpp/plugins/contiv/model/cni"
	"github.com/contiv/vpp/plugins/contiv/model/node"
	"github.com/contiv/vpp/plugins/kvdbproxy"

	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/datasync/resync"
	"github.com/ligato/cn-infra/db/keyval"
	"github.com/ligato/cn-infra/db/keyval/etcd"
	"github.com/ligato/cn-infra/infra"
	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/cn-infra/rpc/grpc"
	"github.com/ligato/cn-infra/rpc/rest"
	"github.com/ligato/cn-infra/servicelabel"
	"github.com/ligato/cn-infra/utils/safeclose"

	"github.com/ligato/vpp-agent/clientv1/linux"
	linuxlocalclient "github.com/ligato/vpp-agent/clientv1/linux/localclient"
	"github.com/ligato/vpp-agent/plugins/govppmux"
	"github.com/ligato/vpp-agent/plugins/vpp"
)

// MgmtIPSeparator is a delimiter inserted between management IPs in nodeInfo structure
const MgmtIPSeparator = ","

// Plugin represents the instance of the Contiv network plugin, that transforms CNI requests received over
// GRPC into configuration for the vswitch VPP in order to connect/disconnect a container into/from the network.
type Plugin struct {
	Deps
	govppCh api.Channel

	configuredContainers *containeridx.ConfigIndex
	cniServer            *remoteCNIserver

	nodeIDAllocator   *idAllocator
	nodeIDsresyncChan chan datasync.ResyncEvent
	nodeIDSchangeChan chan datasync.ChangeEvent
	nodeIDwatchReg    datasync.WatchRegistration

	nodeConfigResyncChan chan datasync.ResyncEvent
	nodeConfigChangeChan chan datasync.ChangeEvent
	nodeConfigWatchReg   datasync.WatchRegistration

	watchReg datasync.WatchRegistration
	resyncCh chan datasync.ResyncEvent
	changeCh chan datasync.ChangeEvent

	ctx           context.Context
	ctxCancelFunc context.CancelFunc

	Config        *Config
	myNodeConfig  *NodeConfig
	nodeIPWatcher chan string
}

// Deps groups the dependencies of the Plugin.
type Deps struct {
	infra.PluginDeps
	ServiceLabel servicelabel.ReaderAPI
	GRPC         grpc.Server
	Proxy        *kvdbproxy.Plugin
	VPP          *vpp.Plugin
	GoVPP        govppmux.API
	Resync       resync.Subscriber
	ETCD         *etcd.Plugin
	Bolt         keyval.KvProtoPlugin
	Watcher      datasync.KeyValProtoWatcher
	HTTPHandlers rest.HTTPHandlers
}

// Init initializes the Contiv plugin. Called automatically by plugin infra upon contiv-agent startup.
func (plugin *Plugin) Init() error {
	// init map with configured containers
	plugin.configuredContainers = containeridx.NewConfigIndex(plugin.Log, "containers",
		plugin.ETCD.NewBroker(plugin.ServiceLabel.GetAgentPrefix()))

	// load config file
	plugin.ctx, plugin.ctxCancelFunc = context.WithCancel(context.Background())
	if plugin.Config == nil {
		if err := plugin.loadExternalConfig(); err != nil {
			return err
		}
		plugin.myNodeConfig = plugin.loadNodeConfig()
	}

	var err error
	plugin.govppCh, err = plugin.GoVPP.NewAPIChannel()
	if err != nil {
		return err
	}

	// init node ID allocator
	nodeIP := ""
	if plugin.myNodeConfig != nil {
		nodeIP = plugin.myNodeConfig.MainVPPInterface.IP
	}
	plugin.nodeIDAllocator = newIDAllocator(plugin.ETCD, plugin.ServiceLabel.GetAgentLabel(), nodeIP)
	nodeID, err := plugin.nodeIDAllocator.getID()
	if err != nil {
		return err
	}
	plugin.Log.Infof("ID of the node is %v", nodeID)

	plugin.nodeIDsresyncChan = make(chan datasync.ResyncEvent)
	plugin.nodeIDSchangeChan = make(chan datasync.ChangeEvent)

	plugin.nodeConfigResyncChan = make(chan datasync.ResyncEvent)
	plugin.nodeConfigChangeChan = make(chan datasync.ChangeEvent)

	plugin.resyncCh = make(chan datasync.ResyncEvent)
	plugin.changeCh = make(chan datasync.ChangeEvent)

	plugin.nodeIDwatchReg, err = plugin.Watcher.Watch("contiv-plugin-ids",
		plugin.nodeIDSchangeChan, plugin.nodeIDsresyncChan, node.AllocatedIDsKeyPrefix)
	if err != nil {
		return err
	}

	plugin.nodeConfigWatchReg, err = plugin.Watcher.Watch("contiv-plugin-node-config",
		plugin.nodeConfigChangeChan, plugin.nodeConfigResyncChan, nodeconfig.Key(plugin.ServiceLabel.GetAgentLabel()))
	if err != nil {
		return err
	}

	plugin.watchReg, err = plugin.Watcher.Watch("contiv-plugin-node",
		plugin.changeCh, plugin.resyncCh, protoNode.KeyPrefix())
	if err != nil {
		return err
	}

	// start the GRPC server handling the CNI requests
	plugin.cniServer, err = newRemoteCNIServer(plugin.Log,
		func() linuxclient.DataChangeDSL {
			return linuxlocalclient.DataChangeRequest(plugin.String())
		},
		plugin.Proxy,
		plugin.configuredContainers,
		plugin.govppCh,
		plugin.VPP.GetSwIfIndexes(),
		plugin.VPP.GetDHCPIndices(),
		plugin.ServiceLabel.GetAgentLabel(),
		plugin.Config,
		plugin.myNodeConfig,
		nodeID,
		plugin.excludedIPsFromNodeCIDR(),
		plugin.ETCD.NewBroker(plugin.ServiceLabel.GetAgentPrefix()),
		plugin.HTTPHandlers)
	if err != nil {
		return fmt.Errorf("Can't create new remote CNI server due to error: %v ", err)
	}
	cni.RegisterRemoteCNIServer(plugin.GRPC.GetServer(), plugin.cniServer)

	plugin.nodeIPWatcher = make(chan string, 1)
	go plugin.watchEvents()
	plugin.cniServer.WatchNodeIP(plugin.nodeIPWatcher)

	// start goroutine handling changes in nodes within the k8s cluster
	go plugin.cniServer.handleNodeEvents(plugin.ctx, plugin.nodeIDsresyncChan, plugin.nodeIDSchangeChan)

	// start goroutine handling changes in the configuration specific to this node
	go plugin.cniServer.handleNodeConfigEvents(plugin.ctx, plugin.nodeConfigResyncChan, plugin.nodeConfigChangeChan)

	return nil
}

// AfterInit is called by the plugin infra after Init of all plugins is finished.
// It registers to the ResyncOrchestrator. The registration is done in this phase
// in order to trigger the resync for this plugin once the resync of VPP plugins is finished.
func (plugin *Plugin) AfterInit() error {
	if plugin.Resync != nil {
		reg := plugin.Resync.Register(string(plugin.PluginName))
		go plugin.handleResync(reg.StatusChan())
	}
	return nil
}

// Close is called by the plugin infra upon agent cleanup. It cleans up the resources allocated by the plugin.
func (plugin *Plugin) Close() error {
	plugin.ctxCancelFunc()
	plugin.cniServer.close()
	//plugin.nodeIDAllocator.releaseID()
	_, err := safeclose.CloseAll(plugin.govppCh, plugin.nodeIDwatchReg, plugin.nodeConfigWatchReg, plugin.watchReg)
	return err
}

// GetPodByIf looks up podName and podNamespace that is associated with logical interface name.
func (plugin *Plugin) GetPodByIf(ifname string) (podNamespace string, podName string, exists bool) {
	ids := plugin.configuredContainers.LookupPodIf(ifname)
	if len(ids) != 1 {
		return "", "", false
	}
	config, found := plugin.configuredContainers.LookupContainer(ids[0])
	if !found {
		return "", "", false
	}
	return config.PodNamespace, config.PodName, true
}

// GetPodByAppNsIndex looks up podName and podNamespace that is associated with the VPP application namespace.
func (plugin *Plugin) GetPodByAppNsIndex(nsIndex uint32) (podNamespace string, podName string, exists bool) {
	nsID, _, found := plugin.VPP.GetAppNsIndexes().LookupName(nsIndex)
	if !found {
		return "", "", false
	}
	ids := plugin.configuredContainers.LookupPodAppNs(nsID)
	if len(ids) != 1 {
		return "", "", false
	}
	config, found := plugin.configuredContainers.LookupContainer(ids[0])
	if !found {
		return "", "", false
	}
	return config.PodNamespace, config.PodName, true
}

// GetIfName looks up logical interface name that corresponds to the interface associated with the given POD name.
func (plugin *Plugin) GetIfName(podNamespace string, podName string) (name string, exists bool) {
	config := plugin.getContainerConfig(podNamespace, podName)
	if config != nil && config.VppIfName != "" {
		return config.VppIfName, true
	}
	plugin.Log.WithFields(logging.Fields{"podNamespace": podNamespace, "podName": podName}).Warn("No matching result found")
	return "", false
}

// GetNsIndex returns the index of the VPP session namespace associated with the given POD name.
func (plugin *Plugin) GetNsIndex(podNamespace string, podName string) (nsIndex uint32, exists bool) {
	config := plugin.getContainerConfig(podNamespace, podName)
	if config != nil {
		nsIndex, _, exists = plugin.VPP.GetAppNsIndexes().LookupIdx(config.AppNamespaceID)
		return nsIndex, exists
	}
	plugin.Log.WithFields(logging.Fields{"podNamespace": podNamespace, "podName": podName}).Warn("No matching result found")
	return 0, false
}

// GetPodSubnet provides subnet used for allocating pod IP addresses across all nodes.
func (plugin *Plugin) GetPodSubnet() *net.IPNet {
	return plugin.cniServer.ipam.PodSubnet()
}

// GetPodNetwork provides subnet used for allocating pod IP addresses on this node.
func (plugin *Plugin) GetPodNetwork() *net.IPNet {
	return plugin.cniServer.ipam.PodNetwork()
}

// GetContainerIndex returns the index of configured containers/pods
func (plugin *Plugin) GetContainerIndex() containeridx.Reader {
	return plugin.configuredContainers
}

// IsTCPstackDisabled returns true if the VPP TCP stack is disabled and only VETHs/TAPs are configured.
func (plugin *Plugin) IsTCPstackDisabled() bool {
	return plugin.Config.TCPstackDisabled
}

// InSTNMode returns true if Contiv operates in the STN mode (single interface for each node).
func (plugin *Plugin) InSTNMode() bool {
	return plugin.cniServer.UseSTN()
}

// NatExternalTraffic returns true if traffic with cluster-outside destination should be S-NATed
// with node IP before being sent out from the node.
func (plugin *Plugin) NatExternalTraffic() bool {
	if plugin.Config.NatExternalTraffic ||
		(plugin.myNodeConfig != nil && plugin.myNodeConfig.NatExternalTraffic) {
		return true
	}
	return false
}

// CleanupIdleNATSessions returns true if cleanup of idle NAT sessions is enabled.
func (plugin *Plugin) CleanupIdleNATSessions() bool {
	return plugin.Config.CleanupIdleNATSessions
}

// GetTCPNATSessionTimeout returns NAT session timeout (in minutes) for TCP connections, used in case that CleanupIdleNATSessions is turned on.
func (plugin *Plugin) GetTCPNATSessionTimeout() uint32 {
	return plugin.Config.TCPNATSessionTimeout
}

// GetOtherNATSessionTimeout returns NAT session timeout (in minutes) for non-TCP connections, used in case that CleanupIdleNATSessions is turned on.
func (plugin *Plugin) GetOtherNATSessionTimeout() uint32 {
	return plugin.Config.OtherNATSessionTimeout
}

// GetServiceLocalEndpointWeight returns the load-balancing weight assigned to locally deployed service endpoints.
func (plugin *Plugin) GetServiceLocalEndpointWeight() uint8 {
	return plugin.Config.ServiceLocalEndpointWeight
}

// GetNatLoopbackIP returns the IP address of a virtual loopback, used to route traffic
// between clients and services via VPP even if the source and destination are the same
// IP addresses and would otherwise be routed locally.
func (plugin *Plugin) GetNatLoopbackIP() net.IP {
	// Last unicast IP from the pod subnet is used as NAT-loopback.
	podNet := plugin.cniServer.ipam.PodNetwork()
	_, broadcastIP := cidr.AddressRange(podNet)
	return cidr.Dec(broadcastIP)
}

// GetNodeIP returns the IP address of this node.
func (plugin *Plugin) GetNodeIP() (ip net.IP, network *net.IPNet) {
	return plugin.cniServer.GetNodeIP()
}

// GetHostIPs returns all IP addresses of this node present in the host network namespace (Linux).
func (plugin *Plugin) GetHostIPs() []net.IP {
	return plugin.cniServer.GetHostIPs()
}

// WatchNodeIP adds given channel to the list of subscribers that are notified upon change
// of nodeIP address. If the channel is not ready to receive notification, the notification is dropped.
func (plugin *Plugin) WatchNodeIP(subscriber chan string) {
	plugin.cniServer.WatchNodeIP(subscriber)
}

// GetMainPhysicalIfName returns name of the "main" interface - i.e. physical interface connecting
// the node with the rest of the cluster.
func (plugin *Plugin) GetMainPhysicalIfName() string {
	return plugin.cniServer.GetMainPhysicalIfName()
}

// GetOtherPhysicalIfNames returns a slice of names of all physical interfaces configured additionally
// to the main interface.
func (plugin *Plugin) GetOtherPhysicalIfNames() []string {
	return plugin.cniServer.GetOtherPhysicalIfNames()
}

// GetHostInterconnectIfName returns the name of the TAP/AF_PACKET interface
// interconnecting VPP with the host stack.
func (plugin *Plugin) GetHostInterconnectIfName() string {
	return plugin.cniServer.GetHostInterconnectIfName()
}

// GetVxlanBVIIfName returns the name of an BVI interface facing towards VXLAN tunnels to other hosts.
// Returns an empty string if VXLAN is not used (in L2 interconnect mode).
func (plugin *Plugin) GetVxlanBVIIfName() string {
	return plugin.cniServer.GetVxlanBVIIfName()
}

// GetDefaultInterface returns the name and the IP address of the interface
// used by the default route to send packets out from VPP towards the default gateway.
// If the default GW is not configured, the function returns zero values.
func (plugin *Plugin) GetDefaultInterface() (ifName string, ifAddress net.IP) {
	return plugin.cniServer.GetDefaultInterface()
}

// RegisterPodPreRemovalHook allows to register callback that will be run for each
// pod immediately before its removal.
func (plugin *Plugin) RegisterPodPreRemovalHook(hook PodActionHook) {
	plugin.cniServer.RegisterPodPreRemovalHook(hook)
}

// RegisterPodPostAddHook allows to register callback that will be run for each
// pod once it is added and before the CNI reply is sent.
func (plugin *Plugin) RegisterPodPostAddHook(hook PodActionHook) {
	plugin.cniServer.RegisterPodPostAddHook(hook)
}

// GetMainVrfID returns the ID of the main network connectivity VRF.
func (plugin *Plugin) GetMainVrfID() uint32 {
	return plugin.cniServer.GetMainVrfID()
}

// GetPodVrfID returns the ID of the POD VRF.
func (plugin *Plugin) GetPodVrfID() uint32 {
	return plugin.cniServer.GetPodVrfID()
}

// handleResync handles resync events of the plugin. Called automatically by the plugin infra.
func (plugin *Plugin) handleResync(resyncChan chan resync.StatusEvent) {
	for {
		select {
		case ev := <-resyncChan:
			status := ev.ResyncStatus()
			if status == resync.Started {
				err := plugin.cniServer.resync()
				if err != nil {
					plugin.Log.Error(err)
				}
			}
			ev.Ack()
		case <-plugin.ctx.Done():
			return
		}
	}
}

// loadExternalConfig attempts to load external configuration from a YAML file.
func (plugin *Plugin) loadExternalConfig() error {
	externalCfg := &Config{}
	found, err := plugin.Cfg.LoadValue(externalCfg) // It tries to lookup `PluginName + "-config"` in the executable arguments.
	if err != nil {
		return fmt.Errorf("External Contiv plugin configuration could not load or other problem happened: %v", err)
	}
	if !found {
		return fmt.Errorf("External Contiv plugin configuration was not found")
	}

	plugin.Config = externalCfg
	plugin.Log.Info("Contiv config: ", externalCfg)
	err = plugin.Config.ApplyIPAMConfig()
	if err != nil {
		return err
	}
	plugin.Config.ApplyDefaults()

	return nil
}

// loadNodeConfig loads config specific for this node (given by its agent label).
func (plugin *Plugin) loadNodeConfig() *NodeConfig {
	myNodeName := plugin.ServiceLabel.GetAgentLabel()
	// first try to get node config from CRD, reflected by contiv-crd into etcd
	// and mirrored into Bolt by us
	nodeConfig := LoadNodeConfigFromCRD(myNodeName, plugin.ETCD, plugin.Bolt, plugin.Log)
	if nodeConfig != nil {
		return nodeConfig
	}
	// try to find the node-specific configuration inside the config file
	return plugin.Config.GetNodeConfig(myNodeName)
}

// getContainerConfig returns the configuration of the container associated with the given POD name.
func (plugin *Plugin) getContainerConfig(podNamespace string, podName string) *container.Persisted {
	podNamesMatch := plugin.configuredContainers.LookupPodName(podName)
	podNamespacesMatch := plugin.configuredContainers.LookupPodNamespace(podNamespace)

	for _, pod1 := range podNamespacesMatch {
		for _, pod2 := range podNamesMatch {
			if pod1 == pod2 {
				data, found := plugin.configuredContainers.LookupContainer(pod1)
				if found {
					return data
				}
			}
		}
	}

	return nil
}

func (plugin *Plugin) watchEvents() {
	for {
		select {
		case newIP := <-plugin.nodeIPWatcher:
			if newIP != "" {
				err := plugin.nodeIDAllocator.updateIP(newIP)
				if err != nil {
					plugin.Log.Error(err)
				}
			}
		case changeEv := <-plugin.changeCh:
			var err error
			key := changeEv.GetKey()
			if strings.HasPrefix(key, protoNode.KeyPrefix()) {
				err = plugin.handleKsrNodeChange(changeEv)
			} else {
				plugin.Log.Warn("Change for unknown key %v received", key)
			}
			changeEv.Done(err)
		case resyncEv := <-plugin.resyncCh:
			var err error
			data := resyncEv.GetValues()

			for prefix, it := range data {
				if prefix == protoNode.KeyPrefix() {
					err = plugin.handleKsrNodeResync(it)
				}
			}
			resyncEv.Done(err)
		case <-plugin.ctx.Done():
		}
	}
}

// handleKsrNodeChange handles change event for the prefix where node data
// is stored by ksr. The aim is to extract node Internal IP - ip address
// that k8s use to access node(management IP). This IP is used as an endpoint
// for services where backends use host networking.
func (plugin *Plugin) handleKsrNodeChange(change datasync.ChangeEvent) error {
	var err error
	// look for our InternalIP skip the others
	if change.GetKey() != protoNode.Key(plugin.ServiceLabel.GetAgentLabel()) {
		return nil
	}
	if change.GetChangeType() == datasync.Delete {
		plugin.Log.Warn("Unexpected delete for node data received")
		return nil
	}
	value := &protoNode.Node{}
	err = change.GetValue(value)
	if err != nil {
		plugin.Log.Error(err)
		return err
	}
	var k8sIPs []string
	for i := range value.Addresses {
		if value.Addresses[i].Type == protoNode.NodeAddress_NodeInternalIP ||
			value.Addresses[i].Type == protoNode.NodeAddress_NodeExternalIP {
			k8sIPs = appendIfMissing(k8sIPs, value.Addresses[i].Address)
		}
	}
	if len(k8sIPs) > 0 {
		ips := strings.Join(k8sIPs, MgmtIPSeparator)
		plugin.Log.Info("Management IP of the node is ", ips)
		return plugin.nodeIDAllocator.updateManagementIP(ips)
	}

	plugin.Log.Warn("Internal IP of the node is missing in ETCD.")

	return err
}

// handleKsrNodeResync handles resync event for the prefix where node data
// is stored by ksr. The aim is to extract node Internal IP - ip address
// that k8s use to access node(management IP). This IP is used as an endpoint
// for services where backends use host networking.
func (plugin *Plugin) handleKsrNodeResync(it datasync.KeyValIterator) error {
	var err error
	for {
		kv, stop := it.GetNext()
		if stop {
			break
		}
		value := &protoNode.Node{}
		err = kv.GetValue(value)
		if err != nil {
			return err
		}

		if value.Name == plugin.ServiceLabel.GetAgentLabel() {
			var k8sIPs []string
			for i := range value.Addresses {
				if value.Addresses[i].Type == protoNode.NodeAddress_NodeInternalIP ||
					value.Addresses[i].Type == protoNode.NodeAddress_NodeExternalIP {
					k8sIPs = appendIfMissing(k8sIPs, value.Addresses[i].Address)
				}
			}
			if len(k8sIPs) > 0 {
				ips := strings.Join(k8sIPs, MgmtIPSeparator)
				plugin.Log.Info("Internal IP of the node is ", ips)
				return plugin.nodeIDAllocator.updateManagementIP(ips)
			}
		}
		plugin.Log.Debug("Internal IP of the node is not in ETCD yet.")
	}
	return err
}

func (plugin *Plugin) excludedIPsFromNodeCIDR() []net.IP {
	if plugin.Config == nil {
		return nil
	}
	var excludedIPs []string
	for _, oneNodeConfig := range plugin.Config.NodeConfig {
		if oneNodeConfig.Gateway == "" {
			continue
		}
		excludedIPs = appendIfMissing(excludedIPs, oneNodeConfig.Gateway)
	}
	var res []net.IP
	for _, ip := range excludedIPs {
		res = append(res, net.ParseIP(ip))
	}
	return res

}

func appendIfMissing(slice []string, s string) []string {
	for _, el := range slice {
		if el == s {
			return slice
		}
	}
	return append(slice, s)
}
