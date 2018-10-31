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

package main

import (
	"os"
	"time"

	"github.com/ligato/cn-infra/agent"
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/datasync/kvdbsync"
	"github.com/ligato/cn-infra/datasync/kvdbsync/local"
	"github.com/ligato/cn-infra/datasync/resync"
	"github.com/ligato/cn-infra/db/keyval/etcd"
	"github.com/ligato/cn-infra/health/probe"
	"github.com/ligato/cn-infra/health/statuscheck"
	"github.com/ligato/cn-infra/logging/logmanager"
	"github.com/ligato/cn-infra/logging/logrus"
	"github.com/ligato/cn-infra/rpc/grpc"
	"github.com/ligato/cn-infra/rpc/prometheus"
	"github.com/ligato/cn-infra/rpc/rest"
	"github.com/ligato/cn-infra/servicelabel"

	"github.com/ligato/vpp-agent/plugins/govppmux"
	"github.com/ligato/vpp-agent/plugins/telemetry"
	"github.com/ligato/vpp-agent/plugins/kvscheduler"
	linux_ifplugin "github.com/ligato/vpp-agent/plugins/linuxv2/ifplugin"
	linux_l3plugin "github.com/ligato/vpp-agent/plugins/linuxv2/l3plugin"
	vpp_ifplugin "github.com/ligato/vpp-agent/plugins/vppv2/ifplugin"
	vpp_natplugin "github.com/ligato/vpp-agent/plugins/vppv2/natplugin"
	vpp_aclplugin "github.com/ligato/vpp-agent/plugins/vppv2/aclplugin"
	vpp_l3plugin "github.com/ligato/vpp-agent/plugins/vppv2/l3plugin"

	"github.com/contiv/vpp/plugins/contiv"
	"github.com/contiv/vpp/plugins/ksr"
	"github.com/contiv/vpp/plugins/kvdbproxy"
	"github.com/contiv/vpp/plugins/policy"
	"github.com/contiv/vpp/plugins/service"
	"github.com/contiv/vpp/plugins/statscollector"
)

const defaultStartupTimeout = 45 * time.Second

// ContivAgent manages vswitch in contiv/vpp solution
type ContivAgent struct {
	LogManager      *logmanager.Plugin
	HTTP            *rest.Plugin
	HealthProbe     *probe.Plugin
	Prometheus      *prometheus.Plugin

	ETCDDataSync    *kvdbsync.Plugin
	NodeIDDataSync  *kvdbsync.Plugin
	ServiceDataSync *kvdbsync.Plugin
	PolicyDataSync  *kvdbsync.Plugin

	KVScheduler     *kvscheduler.Scheduler
	KVProxy         *kvdbproxy.Plugin
	Stats           *statscollector.Plugin

	GoVPP            *govppmux.Plugin
	LinuxIfPlugin    *linux_ifplugin.IfPlugin
	LinuxL3Plugin    *linux_l3plugin.L3Plugin
	VPPIfPlugin      *vpp_ifplugin.IfPlugin
	VPPL3Plugin      *vpp_l3plugin.L3Plugin
	VPPNATPlugin     *vpp_natplugin.NATPlugin
	VPPACLPlugin     *vpp_aclplugin.ACLPlugin

	Telemetry        *telemetry.Plugin
	GRPC             *grpc.Plugin

	Contiv           *contiv.Plugin
	Policy           *policy.Plugin
	Service          *service.Plugin
}

func (c *ContivAgent) String() string {
	return "ContivAgent"
}

// Init is called in startup phase. Method added in order to implement Plugin interface.
func (c *ContivAgent) Init() error {
	return nil
}

// AfterInit triggers the first resync.
func (c *ContivAgent) AfterInit() error {
	resync.DefaultPlugin.DoResync()
	return nil
}

// Close is called in agent's cleanup phase. Method added in order to implement Plugin interface.
func (c *ContivAgent) Close() error {
	return nil
}

func main() {
	// vpp-agent configuration data sync
	etcdDataSync := kvdbsync.NewPlugin(kvdbsync.UseDeps(func(deps *kvdbsync.Deps) {
		deps.KvPlugin = &etcd.DefaultPlugin
		deps.ResyncOrch = &resync.DefaultPlugin
	}))

	kvdbproxy.DefaultPlugin.KVDB = etcdDataSync

	etcd.DefaultPlugin.StatusCheck = nil // disable status check for etcd

	// datasync of Kubernetes state data
	ksrServicelabel := servicelabel.NewPlugin(servicelabel.UseLabel(ksr.MicroserviceLabel))
	ksrServicelabel.SetName("ksrServiceLabel")
	newKSRprefixSync := func(name string) *kvdbsync.Plugin {
		return kvdbsync.NewPlugin(
			kvdbsync.UseDeps(func(deps *kvdbsync.Deps) {
				deps.KvPlugin = &etcd.DefaultPlugin
				deps.ResyncOrch = &resync.DefaultPlugin
				deps.ServiceLabel = ksrServicelabel
				deps.SetName(name)
			}))
	}

	nodeIDDataSync := newKSRprefixSync("nodeIdDataSync")
	serviceDataSync := newKSRprefixSync("serviceDataSync")
	policyDataSync := newKSRprefixSync("policyDataSync")

	// set sources for VPP configuration
	watcher := &datasync.KVProtoWatchers{&kvdbproxy.DefaultPlugin, local.Get()}
	kvscheduler.DefaultPlugin.Watcher = watcher

	// initialize vpp-agent plugins
	linux_ifplugin.DefaultPlugin.VppIfPlugin = &vpp_ifplugin.DefaultPlugin
	vpp_ifplugin.DefaultPlugin.LinuxIfPlugin = &linux_ifplugin.DefaultPlugin
	vpp_ifplugin.DefaultPlugin.PublishStatistics = &statscollector.DefaultPlugin
	vpp_aclplugin.DefaultPlugin.IfPlugin = &vpp_ifplugin.DefaultPlugin

	// we don't want to publish status to etcd
	statuscheck.DefaultPlugin.Transport = nil

	// initialize GRPC
	grpc.DefaultPlugin.HTTP = &rest.DefaultPlugin

	// initialize Contiv plugins
	contivPlugin := contiv.NewPlugin(contiv.UseDeps(func(deps *contiv.Deps) {
		//deps.VPP = vppPlugin
		deps.Watcher = nodeIDDataSync
	}))

	statscollector.DefaultPlugin.Contiv = contivPlugin

	policyPlugin := policy.NewPlugin(policy.UseDeps(func(deps *policy.Deps) {
		deps.Watcher = policyDataSync
		deps.Contiv = contivPlugin
	}))

	servicePlugin := service.NewPlugin(service.UseDeps(func(deps *service.Deps) {
		deps.Watcher = serviceDataSync
		deps.Contiv = contivPlugin
	}))

	// initialize the agent
	contivAgent := &ContivAgent{
		LogManager:      &logmanager.DefaultPlugin,
		HTTP:            &rest.DefaultPlugin,
		HealthProbe:     &probe.DefaultPlugin,
		Prometheus:      &prometheus.DefaultPlugin,
		ETCDDataSync:    etcdDataSync,
		NodeIDDataSync:  nodeIDDataSync,
		ServiceDataSync: serviceDataSync,
		PolicyDataSync:  policyDataSync,
		KVScheduler:     &kvscheduler.DefaultPlugin,
		KVProxy:         &kvdbproxy.DefaultPlugin,
		Stats:           &statscollector.DefaultPlugin,
		GoVPP:           &govppmux.DefaultPlugin,
		LinuxIfPlugin:   &linux_ifplugin.DefaultPlugin,
		LinuxL3Plugin:   &linux_l3plugin.DefaultPlugin,
		VPPIfPlugin:     &vpp_ifplugin.DefaultPlugin,
		VPPL3Plugin:     &vpp_l3plugin.DefaultPlugin,
		VPPNATPlugin:    &vpp_natplugin.DefaultPlugin,
		VPPACLPlugin:    &vpp_aclplugin.DefaultPlugin,
		Telemetry:       &telemetry.DefaultPlugin,
		GRPC:            &grpc.DefaultPlugin,
		Contiv:          contivPlugin,
		Policy:          policyPlugin,
		Service:         servicePlugin,
	}

	a := agent.NewAgent(agent.AllPlugins(contivAgent), agent.StartTimeout(getStartupTimeout()))
	if err := a.Run(); err != nil {
		logrus.DefaultLogger().Fatal(err)
	}

}

func getStartupTimeout() time.Duration {
	var err error
	var timeout time.Duration

	// valid env value must conform to duration format
	// e.g: 45s
	envVal := os.Getenv("STARTUPTIMEOUT")

	if timeout, err = time.ParseDuration(envVal); err != nil {
		timeout = defaultStartupTimeout
	}

	return timeout
}
