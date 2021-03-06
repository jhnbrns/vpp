// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import interfaces1 "github.com/ligato/vpp-agent/plugins/linux/model/interfaces"
import l31 "github.com/ligato/vpp-agent/plugins/linux/model/l3"
import acl "github.com/ligato/vpp-agent/plugins/vpp/model/acl"
import bfd "github.com/ligato/vpp-agent/plugins/vpp/model/bfd"
import interfaces "github.com/ligato/vpp-agent/plugins/vpp/model/interfaces"
import l2 "github.com/ligato/vpp-agent/plugins/vpp/model/l2"
import l3 "github.com/ligato/vpp-agent/plugins/vpp/model/l3"
import l4 "github.com/ligato/vpp-agent/plugins/vpp/model/l4"
import nat "github.com/ligato/vpp-agent/plugins/vpp/model/nat"
import stn "github.com/ligato/vpp-agent/plugins/vpp/model/stn"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Data request is an inventory of supported data types with one or multiple
// items of every type. Universal type for every data change/resync request
type DataRequest struct {
	// vppplugin
	AccessLists           []*acl.AccessLists_Acl                 `protobuf:"bytes,10,rep,name=AccessLists" json:"AccessLists,omitempty"`
	Interfaces            []*interfaces.Interfaces_Interface     `protobuf:"bytes,20,rep,name=Interfaces" json:"Interfaces,omitempty"`
	BfdSessions           []*bfd.SingleHopBFD_Session            `protobuf:"bytes,30,rep,name=BfdSessions" json:"BfdSessions,omitempty"`
	BfdAuthKeys           []*bfd.SingleHopBFD_Key                `protobuf:"bytes,31,rep,name=BfdAuthKeys" json:"BfdAuthKeys,omitempty"`
	BfdEchoFunction       *bfd.SingleHopBFD_EchoFunction         `protobuf:"bytes,32,opt,name=BfdEchoFunction" json:"BfdEchoFunction,omitempty"`
	BridgeDomains         []*l2.BridgeDomains_BridgeDomain       `protobuf:"bytes,40,rep,name=BridgeDomains" json:"BridgeDomains,omitempty"`
	FIBs                  []*l2.FibTable_FibEntry                `protobuf:"bytes,41,rep,name=FIBs" json:"FIBs,omitempty"`
	XCons                 []*l2.XConnectPairs_XConnectPair       `protobuf:"bytes,42,rep,name=XCons" json:"XCons,omitempty"`
	StaticRoutes          []*l3.StaticRoutes_Route               `protobuf:"bytes,50,rep,name=StaticRoutes" json:"StaticRoutes,omitempty"`
	ArpEntries            []*l3.ArpTable_ArpEntry                `protobuf:"bytes,51,rep,name=ArpEntries" json:"ArpEntries,omitempty"`
	ProxyArpInterfaces    []*l3.ProxyArpInterfaces_InterfaceList `protobuf:"bytes,52,rep,name=ProxyArpInterfaces" json:"ProxyArpInterfaces,omitempty"`
	ProxyArpRanges        []*l3.ProxyArpRanges_RangeList         `protobuf:"bytes,53,rep,name=ProxyArpRanges" json:"ProxyArpRanges,omitempty"`
	L4Feature             *l4.L4Features                         `protobuf:"bytes,60,opt,name=L4Feature" json:"L4Feature,omitempty"`
	ApplicationNamespaces []*l4.AppNamespaces_AppNamespace       `protobuf:"bytes,61,rep,name=ApplicationNamespaces" json:"ApplicationNamespaces,omitempty"`
	StnRules              []*stn.STN_Rule                        `protobuf:"bytes,70,rep,name=StnRules" json:"StnRules,omitempty"`
	NatGlobal             *nat.Nat44Global                       `protobuf:"bytes,71,opt,name=NatGlobal" json:"NatGlobal,omitempty"`
	DNATs                 []*nat.Nat44DNat_DNatConfig            `protobuf:"bytes,72,rep,name=DNATs" json:"DNATs,omitempty"`
	// Linuxplugin
	LinuxInterfaces      []*interfaces1.LinuxInterfaces_Interface `protobuf:"bytes,80,rep,name=LinuxInterfaces" json:"LinuxInterfaces,omitempty"`
	LinuxArpEntries      []*l31.LinuxStaticArpEntries_ArpEntry    `protobuf:"bytes,90,rep,name=LinuxArpEntries" json:"LinuxArpEntries,omitempty"`
	LinuxRoutes          []*l31.LinuxStaticRoutes_Route           `protobuf:"bytes,91,rep,name=LinuxRoutes" json:"LinuxRoutes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *DataRequest) Reset()         { *m = DataRequest{} }
func (m *DataRequest) String() string { return proto.CompactTextString(m) }
func (*DataRequest) ProtoMessage()    {}
func (*DataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_4797542838a6c8ba, []int{0}
}
func (m *DataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRequest.Unmarshal(m, b)
}
func (m *DataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRequest.Marshal(b, m, deterministic)
}
func (dst *DataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRequest.Merge(dst, src)
}
func (m *DataRequest) XXX_Size() int {
	return xxx_messageInfo_DataRequest.Size(m)
}
func (m *DataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataRequest proto.InternalMessageInfo

func (m *DataRequest) GetAccessLists() []*acl.AccessLists_Acl {
	if m != nil {
		return m.AccessLists
	}
	return nil
}

func (m *DataRequest) GetInterfaces() []*interfaces.Interfaces_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *DataRequest) GetBfdSessions() []*bfd.SingleHopBFD_Session {
	if m != nil {
		return m.BfdSessions
	}
	return nil
}

func (m *DataRequest) GetBfdAuthKeys() []*bfd.SingleHopBFD_Key {
	if m != nil {
		return m.BfdAuthKeys
	}
	return nil
}

func (m *DataRequest) GetBfdEchoFunction() *bfd.SingleHopBFD_EchoFunction {
	if m != nil {
		return m.BfdEchoFunction
	}
	return nil
}

func (m *DataRequest) GetBridgeDomains() []*l2.BridgeDomains_BridgeDomain {
	if m != nil {
		return m.BridgeDomains
	}
	return nil
}

func (m *DataRequest) GetFIBs() []*l2.FibTable_FibEntry {
	if m != nil {
		return m.FIBs
	}
	return nil
}

func (m *DataRequest) GetXCons() []*l2.XConnectPairs_XConnectPair {
	if m != nil {
		return m.XCons
	}
	return nil
}

func (m *DataRequest) GetStaticRoutes() []*l3.StaticRoutes_Route {
	if m != nil {
		return m.StaticRoutes
	}
	return nil
}

func (m *DataRequest) GetArpEntries() []*l3.ArpTable_ArpEntry {
	if m != nil {
		return m.ArpEntries
	}
	return nil
}

func (m *DataRequest) GetProxyArpInterfaces() []*l3.ProxyArpInterfaces_InterfaceList {
	if m != nil {
		return m.ProxyArpInterfaces
	}
	return nil
}

func (m *DataRequest) GetProxyArpRanges() []*l3.ProxyArpRanges_RangeList {
	if m != nil {
		return m.ProxyArpRanges
	}
	return nil
}

func (m *DataRequest) GetL4Feature() *l4.L4Features {
	if m != nil {
		return m.L4Feature
	}
	return nil
}

func (m *DataRequest) GetApplicationNamespaces() []*l4.AppNamespaces_AppNamespace {
	if m != nil {
		return m.ApplicationNamespaces
	}
	return nil
}

func (m *DataRequest) GetStnRules() []*stn.STN_Rule {
	if m != nil {
		return m.StnRules
	}
	return nil
}

func (m *DataRequest) GetNatGlobal() *nat.Nat44Global {
	if m != nil {
		return m.NatGlobal
	}
	return nil
}

func (m *DataRequest) GetDNATs() []*nat.Nat44DNat_DNatConfig {
	if m != nil {
		return m.DNATs
	}
	return nil
}

func (m *DataRequest) GetLinuxInterfaces() []*interfaces1.LinuxInterfaces_Interface {
	if m != nil {
		return m.LinuxInterfaces
	}
	return nil
}

func (m *DataRequest) GetLinuxArpEntries() []*l31.LinuxStaticArpEntries_ArpEntry {
	if m != nil {
		return m.LinuxArpEntries
	}
	return nil
}

func (m *DataRequest) GetLinuxRoutes() []*l31.LinuxStaticRoutes_Route {
	if m != nil {
		return m.LinuxRoutes
	}
	return nil
}

// NotificationRequest represent a notification request which contains index of next required
// message
type NotificationRequest struct {
	Idx                  uint32   `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationRequest) Reset()         { *m = NotificationRequest{} }
func (m *NotificationRequest) String() string { return proto.CompactTextString(m) }
func (*NotificationRequest) ProtoMessage()    {}
func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_4797542838a6c8ba, []int{1}
}
func (m *NotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationRequest.Unmarshal(m, b)
}
func (m *NotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationRequest.Marshal(b, m, deterministic)
}
func (dst *NotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationRequest.Merge(dst, src)
}
func (m *NotificationRequest) XXX_Size() int {
	return xxx_messageInfo_NotificationRequest.Size(m)
}
func (m *NotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationRequest proto.InternalMessageInfo

func (m *NotificationRequest) GetIdx() uint32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

// Response to data change 'put'
type PutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResponse) Reset()         { *m = PutResponse{} }
func (m *PutResponse) String() string { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()    {}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_4797542838a6c8ba, []int{2}
}
func (m *PutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResponse.Unmarshal(m, b)
}
func (m *PutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResponse.Marshal(b, m, deterministic)
}
func (dst *PutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResponse.Merge(dst, src)
}
func (m *PutResponse) XXX_Size() int {
	return xxx_messageInfo_PutResponse.Size(m)
}
func (m *PutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutResponse proto.InternalMessageInfo

// Response to data change 'del'
type DelResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelResponse) Reset()         { *m = DelResponse{} }
func (m *DelResponse) String() string { return proto.CompactTextString(m) }
func (*DelResponse) ProtoMessage()    {}
func (*DelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_4797542838a6c8ba, []int{3}
}
func (m *DelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelResponse.Unmarshal(m, b)
}
func (m *DelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelResponse.Marshal(b, m, deterministic)
}
func (dst *DelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelResponse.Merge(dst, src)
}
func (m *DelResponse) XXX_Size() int {
	return xxx_messageInfo_DelResponse.Size(m)
}
func (m *DelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DelResponse proto.InternalMessageInfo

// Response to data resync
type ResyncResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResyncResponse) Reset()         { *m = ResyncResponse{} }
func (m *ResyncResponse) String() string { return proto.CompactTextString(m) }
func (*ResyncResponse) ProtoMessage()    {}
func (*ResyncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_4797542838a6c8ba, []int{4}
}
func (m *ResyncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResyncResponse.Unmarshal(m, b)
}
func (m *ResyncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResyncResponse.Marshal(b, m, deterministic)
}
func (dst *ResyncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResyncResponse.Merge(dst, src)
}
func (m *ResyncResponse) XXX_Size() int {
	return xxx_messageInfo_ResyncResponse.Size(m)
}
func (m *ResyncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResyncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResyncResponse proto.InternalMessageInfo

// Response to notification request 'get'. Returns indexed notification.
type NotificationsResponse struct {
	// Index of following notification
	NextIdx uint32 `protobuf:"varint,1,opt,name=nextIdx,proto3" json:"nextIdx,omitempty"`
	// Notification data
	NIf                  *interfaces.InterfaceNotification `protobuf:"bytes,2,opt,name=nIf" json:"nIf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *NotificationsResponse) Reset()         { *m = NotificationsResponse{} }
func (m *NotificationsResponse) String() string { return proto.CompactTextString(m) }
func (*NotificationsResponse) ProtoMessage()    {}
func (*NotificationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_4797542838a6c8ba, []int{5}
}
func (m *NotificationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationsResponse.Unmarshal(m, b)
}
func (m *NotificationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationsResponse.Marshal(b, m, deterministic)
}
func (dst *NotificationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationsResponse.Merge(dst, src)
}
func (m *NotificationsResponse) XXX_Size() int {
	return xxx_messageInfo_NotificationsResponse.Size(m)
}
func (m *NotificationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationsResponse proto.InternalMessageInfo

func (m *NotificationsResponse) GetNextIdx() uint32 {
	if m != nil {
		return m.NextIdx
	}
	return 0
}

func (m *NotificationsResponse) GetNIf() *interfaces.InterfaceNotification {
	if m != nil {
		return m.NIf
	}
	return nil
}

func init() {
	proto.RegisterType((*DataRequest)(nil), "rpc.DataRequest")
	proto.RegisterType((*NotificationRequest)(nil), "rpc.NotificationRequest")
	proto.RegisterType((*PutResponse)(nil), "rpc.PutResponse")
	proto.RegisterType((*DelResponse)(nil), "rpc.DelResponse")
	proto.RegisterType((*ResyncResponse)(nil), "rpc.ResyncResponse")
	proto.RegisterType((*NotificationsResponse)(nil), "rpc.NotificationsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DataChangeService service

type DataChangeServiceClient interface {
	// Creates or updates one or multiple configuration items
	Put(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// Removes one or multiple configuration items
	Del(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DelResponse, error)
}

type dataChangeServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataChangeServiceClient(cc *grpc.ClientConn) DataChangeServiceClient {
	return &dataChangeServiceClient{cc}
}

func (c *dataChangeServiceClient) Put(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/rpc.DataChangeService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataChangeServiceClient) Del(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := c.cc.Invoke(ctx, "/rpc.DataChangeService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataChangeService service

type DataChangeServiceServer interface {
	// Creates or updates one or multiple configuration items
	Put(context.Context, *DataRequest) (*PutResponse, error)
	// Removes one or multiple configuration items
	Del(context.Context, *DataRequest) (*DelResponse, error)
}

func RegisterDataChangeServiceServer(s *grpc.Server, srv DataChangeServiceServer) {
	s.RegisterService(&_DataChangeService_serviceDesc, srv)
}

func _DataChangeService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataChangeServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DataChangeService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataChangeServiceServer).Put(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataChangeService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataChangeServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DataChangeService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataChangeServiceServer).Del(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataChangeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.DataChangeService",
	HandlerType: (*DataChangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _DataChangeService_Put_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _DataChangeService_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// Client API for DataResyncService service

type DataResyncServiceClient interface {
	// Calls vpp-agent resync
	Resync(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*ResyncResponse, error)
}

type dataResyncServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataResyncServiceClient(cc *grpc.ClientConn) DataResyncServiceClient {
	return &dataResyncServiceClient{cc}
}

func (c *dataResyncServiceClient) Resync(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*ResyncResponse, error) {
	out := new(ResyncResponse)
	err := c.cc.Invoke(ctx, "/rpc.DataResyncService/Resync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataResyncService service

type DataResyncServiceServer interface {
	// Calls vpp-agent resync
	Resync(context.Context, *DataRequest) (*ResyncResponse, error)
}

func RegisterDataResyncServiceServer(s *grpc.Server, srv DataResyncServiceServer) {
	s.RegisterService(&_DataResyncService_serviceDesc, srv)
}

func _DataResyncService_Resync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataResyncServiceServer).Resync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DataResyncService/Resync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataResyncServiceServer).Resync(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataResyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.DataResyncService",
	HandlerType: (*DataResyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resync",
			Handler:    _DataResyncService_Resync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// Client API for NotificationService service

type NotificationServiceClient interface {
	// Get notification stack
	Get(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (NotificationService_GetClient, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Get(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (NotificationService_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NotificationService_serviceDesc.Streams[0], "/rpc.NotificationService/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_GetClient interface {
	Recv() (*NotificationsResponse, error)
	grpc.ClientStream
}

type notificationServiceGetClient struct {
	grpc.ClientStream
}

func (x *notificationServiceGetClient) Recv() (*NotificationsResponse, error) {
	m := new(NotificationsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for NotificationService service

type NotificationServiceServer interface {
	// Get notification stack
	Get(*NotificationRequest, NotificationService_GetServer) error
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotificationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).Get(m, &notificationServiceGetServer{stream})
}

type NotificationService_GetServer interface {
	Send(*NotificationsResponse) error
	grpc.ServerStream
}

type notificationServiceGetServer struct {
	grpc.ServerStream
}

func (x *notificationServiceGetServer) Send(m *NotificationsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _NotificationService_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_rpc_4797542838a6c8ba) }

var fileDescriptor_rpc_4797542838a6c8ba = []byte{
	// 879 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x51, 0x6f, 0x23, 0x35,
	0x10, 0x56, 0x09, 0x77, 0xd0, 0x09, 0xed, 0x15, 0xdf, 0x15, 0x99, 0x82, 0x4a, 0xa9, 0x40, 0xb4,
	0x08, 0xbc, 0x68, 0x93, 0x03, 0x41, 0xa9, 0xc4, 0xa6, 0x7b, 0x69, 0xa3, 0x56, 0x21, 0x72, 0xf2,
	0x80, 0xe0, 0xc9, 0xd9, 0x38, 0xa9, 0x25, 0xd7, 0xbb, 0xac, 0xbd, 0xa7, 0xe6, 0x37, 0xf3, 0x27,
	0x90, 0xbd, 0x9b, 0xc4, 0x9b, 0xe6, 0xa4, 0x6b, 0x1f, 0xdc, 0xb5, 0x67, 0xbe, 0xef, 0xdb, 0xf1,
	0xcc, 0xec, 0x34, 0xb0, 0x9d, 0x67, 0x09, 0xc9, 0xf2, 0xd4, 0xa4, 0xa8, 0x91, 0x67, 0xc9, 0xc1,
	0xd9, 0x4c, 0x98, 0xdb, 0x62, 0x4c, 0x92, 0xf4, 0x2e, 0x90, 0x62, 0xc6, 0x4c, 0x1a, 0xbc, 0xcd,
	0xb2, 0x1f, 0xd9, 0x8c, 0x2b, 0x13, 0x64, 0xb2, 0x98, 0x09, 0xa5, 0xad, 0x25, 0xb8, 0x4b, 0x27,
	0x5c, 0x06, 0x2c, 0x71, 0xab, 0x54, 0x78, 0x2c, 0x79, 0x3c, 0x9d, 0xd8, 0x55, 0x91, 0x7b, 0x8f,
	0x23, 0x0b, 0x65, 0x78, 0x3e, 0x65, 0x09, 0xd7, 0xde, 0xb6, 0x92, 0xfa, 0xf5, 0x71, 0x52, 0x32,
	0x0c, 0x64, 0xf8, 0x44, 0x6a, 0x2b, 0x90, 0xad, 0x27, 0x52, 0xdb, 0x81, 0x6c, 0x3f, 0x2d, 0x71,
	0x8a, 0x19, 0xbb, 0x9e, 0x46, 0xd6, 0x46, 0xd9, 0x55, 0x91, 0xaf, 0xdf, 0x87, 0x2c, 0x85, 0x2a,
	0xee, 0xdf, 0x23, 0xef, 0x67, 0x8f, 0x15, 0xf3, 0xd2, 0x77, 0xfc, 0xdf, 0x36, 0x34, 0x63, 0x66,
	0x18, 0xe5, 0xff, 0x16, 0x5c, 0x1b, 0xf4, 0x33, 0x34, 0xa3, 0x24, 0xe1, 0x5a, 0xdf, 0x08, 0x6d,
	0x34, 0x86, 0xa3, 0xc6, 0x49, 0x33, 0x7c, 0x45, 0x6c, 0xb7, 0x79, 0x76, 0x12, 0x25, 0x92, 0xfa,
	0x40, 0xf4, 0x07, 0x40, 0x6f, 0x19, 0x18, 0x7e, 0xe5, 0x68, 0x47, 0xc4, 0x8b, 0xb5, 0xb7, 0x61,
	0x4b, 0x3d, 0x0e, 0x3a, 0x83, 0x66, 0x67, 0x3a, 0x19, 0x72, 0xad, 0x45, 0xaa, 0x34, 0x3e, 0x74,
	0x12, 0x9f, 0x13, 0xdb, 0xaa, 0x43, 0xa1, 0x66, 0x92, 0x5f, 0xa5, 0x59, 0xa7, 0x1b, 0x93, 0x0a,
	0x41, 0x7d, 0x34, 0xfa, 0xc5, 0x91, 0xa3, 0xc2, 0xdc, 0x5e, 0xf3, 0xb9, 0xc6, 0x5f, 0x39, 0xf2,
	0xfe, 0x43, 0xf2, 0x35, 0x9f, 0x53, 0x1f, 0x89, 0xae, 0xe0, 0x45, 0x67, 0x3a, 0x79, 0x93, 0xdc,
	0xa6, 0xdd, 0x42, 0x25, 0x46, 0xa4, 0x0a, 0x1f, 0x1d, 0x6d, 0x9d, 0x34, 0xc3, 0xc3, 0x87, 0x64,
	0x1f, 0x45, 0xd7, 0x69, 0x28, 0x86, 0x9d, 0x4e, 0x2e, 0x26, 0x33, 0x1e, 0xa7, 0x77, 0x4c, 0x28,
	0x8d, 0x4f, 0x5c, 0x10, 0x87, 0x44, 0x86, 0xa4, 0xe6, 0xa8, 0x9d, 0x68, 0x9d, 0x84, 0x4e, 0xe1,
	0xc3, 0x6e, 0xaf, 0xa3, 0xf1, 0x69, 0x75, 0x03, 0x19, 0x92, 0xae, 0x18, 0x8f, 0xd8, 0x58, 0x72,
	0xbb, 0x79, 0xa3, 0x4c, 0x3e, 0xa7, 0x0e, 0x82, 0xda, 0xf0, 0xec, 0xaf, 0x0b, 0x9b, 0xaa, 0xef,
	0x57, 0x2f, 0xb2, 0x06, 0xc5, 0x13, 0x33, 0x60, 0x22, 0xd7, 0xb5, 0x13, 0x2d, 0xc1, 0xe8, 0x37,
	0xf8, 0x64, 0x68, 0x98, 0x11, 0x09, 0x4d, 0x0b, 0xc3, 0x35, 0x0e, 0x1d, 0xf9, 0x33, 0x22, 0x5b,
	0xc4, 0xb7, 0x13, 0xf7, 0xa0, 0x35, 0x2c, 0x7a, 0x0d, 0x10, 0xe5, 0x99, 0x8d, 0x41, 0x70, 0x8d,
	0x5b, 0x8b, 0x10, 0x5b, 0x24, 0xca, 0xb3, 0x32, 0xc4, 0xca, 0x3d, 0xa7, 0x1e, 0x10, 0x8d, 0x00,
	0x0d, 0xf2, 0xf4, 0x7e, 0x1e, 0xe5, 0x99, 0xd7, 0x23, 0x6d, 0x47, 0xff, 0xc6, 0xd2, 0x1f, 0x7a,
	0x57, 0x3d, 0x62, 0xdb, 0x8b, 0x6e, 0xe0, 0xa3, 0x18, 0x76, 0x17, 0x56, 0xca, 0xd4, 0x8c, 0x6b,
	0xfc, 0xda, 0x29, 0x7e, 0xe9, 0x2b, 0x96, 0x1e, 0xe2, 0x1e, 0x4e, 0x69, 0x8d, 0x83, 0x7e, 0x80,
	0xed, 0x9b, 0x76, 0x97, 0x33, 0x53, 0xe4, 0x1c, 0xff, 0xee, 0x2a, 0xbf, 0x4b, 0x64, 0x9b, 0x2c,
	0x8d, 0x9a, 0xae, 0x00, 0x68, 0x04, 0xfb, 0x51, 0x96, 0x49, 0x91, 0x30, 0x5b, 0xf2, 0x3e, 0xbb,
	0xe3, 0x3a, 0x73, 0x97, 0x39, 0x5f, 0x94, 0xa0, 0x4d, 0xa2, 0x2c, 0x5b, 0x39, 0x6a, 0x27, 0xba,
	0x99, 0x8c, 0x4e, 0xe1, 0xe3, 0xa1, 0x51, 0xb4, 0x90, 0x5c, 0xe3, 0xae, 0x13, 0xda, 0x21, 0x76,
	0x56, 0x0c, 0x47, 0x7d, 0x62, 0xad, 0x74, 0xe9, 0x46, 0x04, 0xb6, 0xfb, 0xcc, 0x5c, 0xca, 0x74,
	0xcc, 0x24, 0xbe, 0x74, 0xe1, 0xee, 0x11, 0x3b, 0x94, 0xfa, 0xcc, 0xb4, 0xdb, 0xa5, 0x9d, 0xae,
	0x20, 0x28, 0x80, 0x67, 0x71, 0x3f, 0x1a, 0x69, 0x7c, 0x55, 0x7d, 0x4e, 0x4b, 0x6c, 0xdc, 0x67,
	0x86, 0xd8, 0x3f, 0x17, 0xa9, 0x9a, 0x8a, 0x19, 0x2d, 0x71, 0xe8, 0x4f, 0x78, 0x71, 0x63, 0x47,
	0x85, 0x57, 0xa8, 0x81, 0xa3, 0x7e, 0xeb, 0x7f, 0xcc, 0x6b, 0x10, 0xef, 0x8b, 0x5e, 0x67, 0xa3,
	0x9b, 0x4a, 0xd0, 0x6b, 0x9c, 0xbf, 0x9d, 0xe0, 0xb1, 0xad, 0x93, 0x73, 0x95, 0x3d, 0xb6, 0x02,
	0xac, 0xba, 0x68, 0x9d, 0x8a, 0xce, 0xa1, 0xe9, 0x4c, 0x55, 0xf3, 0xfe, 0xe3, 0x94, 0xbe, 0x58,
	0x53, 0xaa, 0x75, 0xb0, 0x8f, 0x3f, 0xfe, 0x0e, 0x5e, 0xf6, 0x53, 0x23, 0xa6, 0x55, 0x0d, 0x16,
	0x43, 0x6f, 0x0f, 0x1a, 0x62, 0x72, 0x8f, 0xb7, 0x8e, 0xb6, 0x4e, 0x76, 0xa8, 0xdd, 0x1e, 0xef,
	0x40, 0x73, 0x50, 0x18, 0xca, 0x75, 0x96, 0x2a, 0xcd, 0xed, 0x31, 0xe6, 0x72, 0x79, 0xdc, 0x83,
	0x5d, 0xca, 0xf5, 0x5c, 0x25, 0x4b, 0xcb, 0x14, 0xf6, 0x7d, 0x61, 0xbd, 0x70, 0x20, 0x0c, 0x1f,
	0x29, 0x7e, 0x6f, 0x7a, 0x4b, 0xf9, 0xc5, 0x11, 0xb5, 0xa0, 0xa1, 0x7a, 0x53, 0xfc, 0x81, 0x2b,
	0xe2, 0xd7, 0x1b, 0x47, 0x65, 0x2d, 0x56, 0x8b, 0x0e, 0x05, 0x7c, 0x6a, 0xa7, 0xf5, 0xc5, 0xad,
	0xed, 0xde, 0x21, 0xcf, 0xdf, 0x8a, 0x84, 0xa3, 0x53, 0x68, 0x0c, 0x0a, 0x83, 0xf6, 0x88, 0xfd,
	0x55, 0xe1, 0x0d, 0xf3, 0x83, 0xd2, 0xe2, 0x5d, 0xc4, 0x42, 0x63, 0x2e, 0xdf, 0x09, 0xf5, 0x2e,
	0x19, 0xc6, 0xe5, 0xab, 0xca, 0x8b, 0x2e, 0x5e, 0x15, 0xc0, 0xf3, 0xd2, 0xb0, 0x41, 0xe2, 0xa5,
	0xb3, 0xd4, 0x13, 0x13, 0x8e, 0xea, 0x19, 0x5f, 0xe8, 0x9c, 0x43, 0xe3, 0x92, 0x1b, 0x84, 0x1d,
	0x65, 0x43, 0x49, 0x0e, 0x0e, 0x1e, 0x78, 0x96, 0x39, 0xfd, 0x69, 0x6b, 0xfc, 0xdc, 0xfd, 0xf3,
	0x6a, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x78, 0x23, 0x1c, 0x57, 0x48, 0x09, 0x00, 0x00,
}
