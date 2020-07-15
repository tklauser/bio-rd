// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bio-routing/bio-rd/cmd/ris/api/ris.proto

package api

import (
	context "context"
	fmt "fmt"
	api "github.com/bio-routing/bio-rd/net/api"
	api1 "github.com/bio-routing/bio-rd/route/api"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ObserveRIBRequest_AFISAFI int32

const (
	ObserveRIBRequest_IPv4Unicast ObserveRIBRequest_AFISAFI = 0
	ObserveRIBRequest_IPv6Unicast ObserveRIBRequest_AFISAFI = 1
)

var ObserveRIBRequest_AFISAFI_name = map[int32]string{
	0: "IPv4Unicast",
	1: "IPv6Unicast",
}

var ObserveRIBRequest_AFISAFI_value = map[string]int32{
	"IPv4Unicast": 0,
	"IPv6Unicast": 1,
}

func (x ObserveRIBRequest_AFISAFI) String() string {
	return proto.EnumName(ObserveRIBRequest_AFISAFI_name, int32(x))
}

func (ObserveRIBRequest_AFISAFI) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{6, 0}
}

type DumpRIBRequest_AFISAFI int32

const (
	DumpRIBRequest_IPv4Unicast DumpRIBRequest_AFISAFI = 0
	DumpRIBRequest_IPv6Unicast DumpRIBRequest_AFISAFI = 1
)

var DumpRIBRequest_AFISAFI_name = map[int32]string{
	0: "IPv4Unicast",
	1: "IPv6Unicast",
}

var DumpRIBRequest_AFISAFI_value = map[string]int32{
	"IPv4Unicast": 0,
	"IPv6Unicast": 1,
}

func (x DumpRIBRequest_AFISAFI) String() string {
	return proto.EnumName(DumpRIBRequest_AFISAFI_name, int32(x))
}

func (DumpRIBRequest_AFISAFI) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{8, 0}
}

type LPMRequest struct {
	Router               string      `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
	VrfId                uint64      `protobuf:"varint,2,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	Pfx                  *api.Prefix `protobuf:"bytes,3,opt,name=pfx,proto3" json:"pfx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LPMRequest) Reset()         { *m = LPMRequest{} }
func (m *LPMRequest) String() string { return proto.CompactTextString(m) }
func (*LPMRequest) ProtoMessage()    {}
func (*LPMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{0}
}

func (m *LPMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LPMRequest.Unmarshal(m, b)
}
func (m *LPMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LPMRequest.Marshal(b, m, deterministic)
}
func (m *LPMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LPMRequest.Merge(m, src)
}
func (m *LPMRequest) XXX_Size() int {
	return xxx_messageInfo_LPMRequest.Size(m)
}
func (m *LPMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LPMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LPMRequest proto.InternalMessageInfo

func (m *LPMRequest) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *LPMRequest) GetVrfId() uint64 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *LPMRequest) GetPfx() *api.Prefix {
	if m != nil {
		return m.Pfx
	}
	return nil
}

type LPMResponse struct {
	Routes               []*api1.Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LPMResponse) Reset()         { *m = LPMResponse{} }
func (m *LPMResponse) String() string { return proto.CompactTextString(m) }
func (*LPMResponse) ProtoMessage()    {}
func (*LPMResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{1}
}

func (m *LPMResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LPMResponse.Unmarshal(m, b)
}
func (m *LPMResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LPMResponse.Marshal(b, m, deterministic)
}
func (m *LPMResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LPMResponse.Merge(m, src)
}
func (m *LPMResponse) XXX_Size() int {
	return xxx_messageInfo_LPMResponse.Size(m)
}
func (m *LPMResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LPMResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LPMResponse proto.InternalMessageInfo

func (m *LPMResponse) GetRoutes() []*api1.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type GetRequest struct {
	Router               string      `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
	VrfId                uint64      `protobuf:"varint,2,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	Pfx                  *api.Prefix `protobuf:"bytes,3,opt,name=pfx,proto3" json:"pfx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *GetRequest) GetVrfId() uint64 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *GetRequest) GetPfx() *api.Prefix {
	if m != nil {
		return m.Pfx
	}
	return nil
}

type GetResponse struct {
	Routes               []*api1.Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{3}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetRoutes() []*api1.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type GetLongerRequest struct {
	Router               string      `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
	VrfId                uint64      `protobuf:"varint,2,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	Pfx                  *api.Prefix `protobuf:"bytes,3,opt,name=pfx,proto3" json:"pfx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetLongerRequest) Reset()         { *m = GetLongerRequest{} }
func (m *GetLongerRequest) String() string { return proto.CompactTextString(m) }
func (*GetLongerRequest) ProtoMessage()    {}
func (*GetLongerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{4}
}

func (m *GetLongerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLongerRequest.Unmarshal(m, b)
}
func (m *GetLongerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLongerRequest.Marshal(b, m, deterministic)
}
func (m *GetLongerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLongerRequest.Merge(m, src)
}
func (m *GetLongerRequest) XXX_Size() int {
	return xxx_messageInfo_GetLongerRequest.Size(m)
}
func (m *GetLongerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLongerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLongerRequest proto.InternalMessageInfo

func (m *GetLongerRequest) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *GetLongerRequest) GetVrfId() uint64 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *GetLongerRequest) GetPfx() *api.Prefix {
	if m != nil {
		return m.Pfx
	}
	return nil
}

type GetLongerResponse struct {
	Routes               []*api1.Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetLongerResponse) Reset()         { *m = GetLongerResponse{} }
func (m *GetLongerResponse) String() string { return proto.CompactTextString(m) }
func (*GetLongerResponse) ProtoMessage()    {}
func (*GetLongerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{5}
}

func (m *GetLongerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLongerResponse.Unmarshal(m, b)
}
func (m *GetLongerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLongerResponse.Marshal(b, m, deterministic)
}
func (m *GetLongerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLongerResponse.Merge(m, src)
}
func (m *GetLongerResponse) XXX_Size() int {
	return xxx_messageInfo_GetLongerResponse.Size(m)
}
func (m *GetLongerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLongerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLongerResponse proto.InternalMessageInfo

func (m *GetLongerResponse) GetRoutes() []*api1.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type ObserveRIBRequest struct {
	Router               string                    `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
	VrfId                uint64                    `protobuf:"varint,2,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	Afisafi              ObserveRIBRequest_AFISAFI `protobuf:"varint,3,opt,name=afisafi,proto3,enum=bio.ris.ObserveRIBRequest_AFISAFI" json:"afisafi,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ObserveRIBRequest) Reset()         { *m = ObserveRIBRequest{} }
func (m *ObserveRIBRequest) String() string { return proto.CompactTextString(m) }
func (*ObserveRIBRequest) ProtoMessage()    {}
func (*ObserveRIBRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{6}
}

func (m *ObserveRIBRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObserveRIBRequest.Unmarshal(m, b)
}
func (m *ObserveRIBRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObserveRIBRequest.Marshal(b, m, deterministic)
}
func (m *ObserveRIBRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObserveRIBRequest.Merge(m, src)
}
func (m *ObserveRIBRequest) XXX_Size() int {
	return xxx_messageInfo_ObserveRIBRequest.Size(m)
}
func (m *ObserveRIBRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ObserveRIBRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ObserveRIBRequest proto.InternalMessageInfo

func (m *ObserveRIBRequest) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *ObserveRIBRequest) GetVrfId() uint64 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *ObserveRIBRequest) GetAfisafi() ObserveRIBRequest_AFISAFI {
	if m != nil {
		return m.Afisafi
	}
	return ObserveRIBRequest_IPv4Unicast
}

type RIBUpdate struct {
	Advertisement        bool        `protobuf:"varint,1,opt,name=advertisement,proto3" json:"advertisement,omitempty"`
	IsInitialDump        bool        `protobuf:"varint,3,opt,name=is_initial_dump,json=isInitialDump,proto3" json:"is_initial_dump,omitempty"`
	Route                *api1.Route `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RIBUpdate) Reset()         { *m = RIBUpdate{} }
func (m *RIBUpdate) String() string { return proto.CompactTextString(m) }
func (*RIBUpdate) ProtoMessage()    {}
func (*RIBUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{7}
}

func (m *RIBUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RIBUpdate.Unmarshal(m, b)
}
func (m *RIBUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RIBUpdate.Marshal(b, m, deterministic)
}
func (m *RIBUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RIBUpdate.Merge(m, src)
}
func (m *RIBUpdate) XXX_Size() int {
	return xxx_messageInfo_RIBUpdate.Size(m)
}
func (m *RIBUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_RIBUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_RIBUpdate proto.InternalMessageInfo

func (m *RIBUpdate) GetAdvertisement() bool {
	if m != nil {
		return m.Advertisement
	}
	return false
}

func (m *RIBUpdate) GetIsInitialDump() bool {
	if m != nil {
		return m.IsInitialDump
	}
	return false
}

func (m *RIBUpdate) GetRoute() *api1.Route {
	if m != nil {
		return m.Route
	}
	return nil
}

type DumpRIBRequest struct {
	Router               string                 `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
	VrfId                uint64                 `protobuf:"varint,2,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	Afisafi              DumpRIBRequest_AFISAFI `protobuf:"varint,3,opt,name=afisafi,proto3,enum=bio.ris.DumpRIBRequest_AFISAFI" json:"afisafi,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DumpRIBRequest) Reset()         { *m = DumpRIBRequest{} }
func (m *DumpRIBRequest) String() string { return proto.CompactTextString(m) }
func (*DumpRIBRequest) ProtoMessage()    {}
func (*DumpRIBRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{8}
}

func (m *DumpRIBRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpRIBRequest.Unmarshal(m, b)
}
func (m *DumpRIBRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpRIBRequest.Marshal(b, m, deterministic)
}
func (m *DumpRIBRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpRIBRequest.Merge(m, src)
}
func (m *DumpRIBRequest) XXX_Size() int {
	return xxx_messageInfo_DumpRIBRequest.Size(m)
}
func (m *DumpRIBRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpRIBRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DumpRIBRequest proto.InternalMessageInfo

func (m *DumpRIBRequest) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *DumpRIBRequest) GetVrfId() uint64 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *DumpRIBRequest) GetAfisafi() DumpRIBRequest_AFISAFI {
	if m != nil {
		return m.Afisafi
	}
	return DumpRIBRequest_IPv4Unicast
}

type DumpRIBReply struct {
	Route                *api1.Route `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DumpRIBReply) Reset()         { *m = DumpRIBReply{} }
func (m *DumpRIBReply) String() string { return proto.CompactTextString(m) }
func (*DumpRIBReply) ProtoMessage()    {}
func (*DumpRIBReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{9}
}

func (m *DumpRIBReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpRIBReply.Unmarshal(m, b)
}
func (m *DumpRIBReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpRIBReply.Marshal(b, m, deterministic)
}
func (m *DumpRIBReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpRIBReply.Merge(m, src)
}
func (m *DumpRIBReply) XXX_Size() int {
	return xxx_messageInfo_DumpRIBReply.Size(m)
}
func (m *DumpRIBReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpRIBReply.DiscardUnknown(m)
}

var xxx_messageInfo_DumpRIBReply proto.InternalMessageInfo

func (m *DumpRIBReply) GetRoute() *api1.Route {
	if m != nil {
		return m.Route
	}
	return nil
}

type GetRoutersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRoutersRequest) Reset()         { *m = GetRoutersRequest{} }
func (m *GetRoutersRequest) String() string { return proto.CompactTextString(m) }
func (*GetRoutersRequest) ProtoMessage()    {}
func (*GetRoutersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{10}
}

func (m *GetRoutersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoutersRequest.Unmarshal(m, b)
}
func (m *GetRoutersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoutersRequest.Marshal(b, m, deterministic)
}
func (m *GetRoutersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoutersRequest.Merge(m, src)
}
func (m *GetRoutersRequest) XXX_Size() int {
	return xxx_messageInfo_GetRoutersRequest.Size(m)
}
func (m *GetRoutersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoutersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoutersRequest proto.InternalMessageInfo

type Router struct {
	SysName              string   `protobuf:"bytes,1,opt,name=sys_name,json=sysName,proto3" json:"sys_name,omitempty"`
	VrfIds               []uint64 `protobuf:"varint,2,rep,packed,name=vrf_ids,json=vrfIds,proto3" json:"vrf_ids,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Router) Reset()         { *m = Router{} }
func (m *Router) String() string { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()    {}
func (*Router) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{11}
}

func (m *Router) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Router.Unmarshal(m, b)
}
func (m *Router) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Router.Marshal(b, m, deterministic)
}
func (m *Router) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Router.Merge(m, src)
}
func (m *Router) XXX_Size() int {
	return xxx_messageInfo_Router.Size(m)
}
func (m *Router) XXX_DiscardUnknown() {
	xxx_messageInfo_Router.DiscardUnknown(m)
}

var xxx_messageInfo_Router proto.InternalMessageInfo

func (m *Router) GetSysName() string {
	if m != nil {
		return m.SysName
	}
	return ""
}

func (m *Router) GetVrfIds() []uint64 {
	if m != nil {
		return m.VrfIds
	}
	return nil
}

func (m *Router) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type GetRoutersResponse struct {
	Routers              []*Router `protobuf:"bytes,1,rep,name=routers,proto3" json:"routers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetRoutersResponse) Reset()         { *m = GetRoutersResponse{} }
func (m *GetRoutersResponse) String() string { return proto.CompactTextString(m) }
func (*GetRoutersResponse) ProtoMessage()    {}
func (*GetRoutersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe1202aa518913f, []int{12}
}

func (m *GetRoutersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoutersResponse.Unmarshal(m, b)
}
func (m *GetRoutersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoutersResponse.Marshal(b, m, deterministic)
}
func (m *GetRoutersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoutersResponse.Merge(m, src)
}
func (m *GetRoutersResponse) XXX_Size() int {
	return xxx_messageInfo_GetRoutersResponse.Size(m)
}
func (m *GetRoutersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoutersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoutersResponse proto.InternalMessageInfo

func (m *GetRoutersResponse) GetRouters() []*Router {
	if m != nil {
		return m.Routers
	}
	return nil
}

func init() {
	proto.RegisterEnum("bio.ris.ObserveRIBRequest_AFISAFI", ObserveRIBRequest_AFISAFI_name, ObserveRIBRequest_AFISAFI_value)
	proto.RegisterEnum("bio.ris.DumpRIBRequest_AFISAFI", DumpRIBRequest_AFISAFI_name, DumpRIBRequest_AFISAFI_value)
	proto.RegisterType((*LPMRequest)(nil), "bio.ris.LPMRequest")
	proto.RegisterType((*LPMResponse)(nil), "bio.ris.LPMResponse")
	proto.RegisterType((*GetRequest)(nil), "bio.ris.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "bio.ris.GetResponse")
	proto.RegisterType((*GetLongerRequest)(nil), "bio.ris.GetLongerRequest")
	proto.RegisterType((*GetLongerResponse)(nil), "bio.ris.GetLongerResponse")
	proto.RegisterType((*ObserveRIBRequest)(nil), "bio.ris.ObserveRIBRequest")
	proto.RegisterType((*RIBUpdate)(nil), "bio.ris.RIBUpdate")
	proto.RegisterType((*DumpRIBRequest)(nil), "bio.ris.DumpRIBRequest")
	proto.RegisterType((*DumpRIBReply)(nil), "bio.ris.DumpRIBReply")
	proto.RegisterType((*GetRoutersRequest)(nil), "bio.ris.GetRoutersRequest")
	proto.RegisterType((*Router)(nil), "bio.ris.Router")
	proto.RegisterType((*GetRoutersResponse)(nil), "bio.ris.GetRoutersResponse")
}

func init() {
	proto.RegisterFile("github.com/bio-routing/bio-rd/cmd/ris/api/ris.proto", fileDescriptor_ffe1202aa518913f)
}

var fileDescriptor_ffe1202aa518913f = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x4f, 0xdb, 0x4c,
	0x10, 0x66, 0x09, 0xc4, 0x64, 0xf2, 0x42, 0x60, 0x79, 0x29, 0xc1, 0x3d, 0x34, 0xb5, 0x2a, 0x14,
	0x84, 0xea, 0xa0, 0x50, 0x51, 0x55, 0xfd, 0x12, 0x08, 0x11, 0x59, 0x82, 0x36, 0x5a, 0x44, 0x0f,
	0x3d, 0x34, 0x72, 0xe2, 0x35, 0x5d, 0x09, 0x7f, 0x74, 0x77, 0x13, 0x91, 0x63, 0x7f, 0x41, 0x7f,
	0x46, 0x4f, 0xfd, 0x8f, 0x95, 0x77, 0x1d, 0xc7, 0x2e, 0x94, 0x16, 0x2a, 0x2e, 0x89, 0x77, 0x66,
	0x9f, 0x99, 0x67, 0x1e, 0xcd, 0xcc, 0xc2, 0xee, 0x39, 0x93, 0x9f, 0x87, 0x7d, 0x7b, 0x10, 0x05,
	0xad, 0x3e, 0x8b, 0x9e, 0xf2, 0x68, 0x28, 0x59, 0x78, 0xae, 0xbf, 0xbd, 0xd6, 0x20, 0xf0, 0x5a,
	0x9c, 0x89, 0x96, 0x1b, 0xb3, 0xe4, 0xdf, 0x8e, 0x79, 0x24, 0x23, 0x6c, 0xf4, 0x59, 0x64, 0x73,
	0x26, 0xcc, 0xd6, 0xcd, 0xe8, 0x90, 0x4a, 0x85, 0x0c, 0xa9, 0xd4, 0x48, 0xf3, 0x0f, 0xe9, 0x92,
	0x23, 0xd5, 0xc9, 0x92, 0x2f, 0x0d, 0xb2, 0x3e, 0x01, 0x1c, 0x77, 0x4f, 0x08, 0xfd, 0x32, 0xa4,
	0x42, 0xe2, 0x07, 0x50, 0x56, 0x4e, 0x5e, 0x47, 0x0d, 0xd4, 0xac, 0x90, 0xf4, 0x84, 0xd7, 0xa0,
	0x3c, 0xe2, 0x7e, 0x8f, 0x79, 0xf5, 0xd9, 0x06, 0x6a, 0xce, 0x91, 0xf9, 0x11, 0xf7, 0x1d, 0x0f,
	0x3f, 0x86, 0x52, 0xec, 0x5f, 0xd6, 0x4b, 0x0d, 0xd4, 0xac, 0xb6, 0x6b, 0x76, 0xc2, 0x3c, 0xa1,
	0xd3, 0xe5, 0xd4, 0x67, 0x97, 0x24, 0xf1, 0x59, 0xcf, 0xa1, 0xaa, 0xe2, 0x8b, 0x38, 0x0a, 0x05,
	0xc5, 0xcd, 0x34, 0x81, 0xa8, 0xa3, 0x46, 0xa9, 0x59, 0x6d, 0x2f, 0x2b, 0x90, 0x26, 0x44, 0x92,
	0xdf, 0x34, 0xa5, 0x48, 0x88, 0x75, 0xa8, 0xbc, 0x57, 0x62, 0x2a, 0xfe, 0xad, 0x89, 0x79, 0xb0,
	0xdc, 0xa1, 0xf2, 0x38, 0x0a, 0xcf, 0x29, 0xbf, 0x3f, 0x7a, 0xaf, 0x61, 0x25, 0x97, 0xe5, 0xd6,
	0x24, 0x7f, 0x20, 0x58, 0x79, 0xdf, 0x17, 0x94, 0x8f, 0x28, 0x71, 0x0e, 0xee, 0x48, 0xf3, 0x15,
	0x18, 0xae, 0xcf, 0x84, 0xeb, 0x33, 0x45, 0x75, 0xa9, 0x6d, 0xd9, 0x69, 0x73, 0xda, 0x57, 0x62,
	0xdb, 0xfb, 0x47, 0xce, 0xe9, 0xfe, 0x91, 0x43, 0x26, 0x10, 0x6b, 0x1b, 0x8c, 0xd4, 0x86, 0x6b,
	0x50, 0x75, 0xba, 0xa3, 0x67, 0x67, 0x21, 0x1b, 0xb8, 0x42, 0x2e, 0xcf, 0xa4, 0x86, 0xbd, 0x89,
	0x01, 0x59, 0x5f, 0x11, 0x54, 0x88, 0x73, 0x70, 0x16, 0x7b, 0xae, 0xa4, 0xf8, 0x09, 0x2c, 0xba,
	0xde, 0x88, 0x72, 0xc9, 0x04, 0x0d, 0x68, 0x28, 0x15, 0xdd, 0x05, 0x52, 0x34, 0xe2, 0x4d, 0xa8,
	0x31, 0xd1, 0x63, 0x21, 0x93, 0xcc, 0xbd, 0xe8, 0x79, 0xc3, 0x20, 0x56, 0x34, 0x17, 0xc8, 0x22,
	0x13, 0x8e, 0xb6, 0x1e, 0x0e, 0x83, 0x18, 0x6f, 0xc2, 0xbc, 0xaa, 0x53, 0x15, 0x77, 0x9d, 0x68,
	0xda, 0x6d, 0x7d, 0x47, 0xb0, 0x94, 0x00, 0xee, 0x2e, 0xd8, 0x8b, 0x5f, 0x05, 0x7b, 0x94, 0x09,
	0x56, 0x0c, 0xfc, 0x8f, 0x6a, 0xed, 0xc1, 0x7f, 0x59, 0xbc, 0xf8, 0x62, 0x3c, 0xad, 0x10, 0xdd,
	0x5c, 0xe1, 0xaa, 0x6a, 0x2a, 0x65, 0xe2, 0x22, 0xa5, 0x62, 0x7d, 0x80, 0xb2, 0xb6, 0xe0, 0x0d,
	0x58, 0x10, 0x63, 0xd1, 0x0b, 0xdd, 0x80, 0xa6, 0xf5, 0x1a, 0x62, 0x2c, 0xde, 0xb9, 0x01, 0xc5,
	0xeb, 0x60, 0xe8, 0x82, 0x45, 0x7d, 0xb6, 0x51, 0x6a, 0xce, 0x91, 0xb2, 0xaa, 0x58, 0xe0, 0x3a,
	0x18, 0xae, 0xe7, 0x71, 0x2a, 0x84, 0x2a, 0xb9, 0x42, 0x26, 0x47, 0xeb, 0x2d, 0xe0, 0x7c, 0xb2,
	0xb4, 0x85, 0xb7, 0xc0, 0xd0, 0x1a, 0x4e, 0x7a, 0xb8, 0x96, 0x49, 0xa4, 0xaf, 0x92, 0x89, 0xbf,
	0xfd, 0xad, 0x04, 0x1b, 0x44, 0xef, 0x30, 0x27, 0xf4, 0x23, 0x1e, 0xb8, 0x92, 0x45, 0xe1, 0x29,
	0xe5, 0x23, 0x36, 0xa0, 0xb8, 0x0d, 0xa5, 0xe3, 0xee, 0x09, 0x5e, 0xcd, 0xe0, 0xd3, 0x35, 0x66,
	0xfe, 0x5f, 0x34, 0xea, 0xd4, 0xd6, 0x4c, 0x82, 0xe9, 0x50, 0x99, 0xc3, 0x4c, 0x37, 0x4c, 0x0e,
	0x93, 0x5b, 0x0b, 0xd6, 0x0c, 0xee, 0xe8, 0x3d, 0xa4, 0x39, 0x61, 0xb3, 0x70, 0xab, 0x20, 0xa4,
	0xf9, 0xf0, 0x5a, 0x5f, 0x16, 0xe8, 0x10, 0x2a, 0xd9, 0x44, 0xe3, 0x8d, 0xfc, 0xdd, 0xc2, 0x2e,
	0x31, 0xcd, 0xeb, 0x5c, 0x59, 0x94, 0x37, 0x00, 0xd3, 0xd9, 0xcb, 0xd1, 0xb9, 0x32, 0x90, 0x26,
	0x9e, 0x0a, 0x3b, 0x19, 0xac, 0x1d, 0x84, 0x5f, 0x82, 0x91, 0xb6, 0x0e, 0x5e, 0xff, 0x4d, 0x73,
	0x9a, 0x6b, 0x57, 0x1d, 0xf1, 0xc5, 0x78, 0x07, 0x1d, 0x6c, 0x7f, 0xdc, 0xfa, 0xeb, 0x27, 0xad,
	0x5f, 0x56, 0x0f, 0xcc, 0xee, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x19, 0x12, 0xf0, 0x06,
	0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoutingInformationServiceClient is the client API for RoutingInformationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoutingInformationServiceClient interface {
	LPM(ctx context.Context, in *LPMRequest, opts ...grpc.CallOption) (*LPMResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetRouters(ctx context.Context, in *GetRoutersRequest, opts ...grpc.CallOption) (*GetRoutersResponse, error)
	GetLonger(ctx context.Context, in *GetLongerRequest, opts ...grpc.CallOption) (*GetLongerResponse, error)
	ObserveRIB(ctx context.Context, in *ObserveRIBRequest, opts ...grpc.CallOption) (RoutingInformationService_ObserveRIBClient, error)
	DumpRIB(ctx context.Context, in *DumpRIBRequest, opts ...grpc.CallOption) (RoutingInformationService_DumpRIBClient, error)
}

type routingInformationServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoutingInformationServiceClient(cc *grpc.ClientConn) RoutingInformationServiceClient {
	return &routingInformationServiceClient{cc}
}

func (c *routingInformationServiceClient) LPM(ctx context.Context, in *LPMRequest, opts ...grpc.CallOption) (*LPMResponse, error) {
	out := new(LPMResponse)
	err := c.cc.Invoke(ctx, "/bio.ris.RoutingInformationService/LPM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingInformationServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/bio.ris.RoutingInformationService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingInformationServiceClient) GetRouters(ctx context.Context, in *GetRoutersRequest, opts ...grpc.CallOption) (*GetRoutersResponse, error) {
	out := new(GetRoutersResponse)
	err := c.cc.Invoke(ctx, "/bio.ris.RoutingInformationService/GetRouters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingInformationServiceClient) GetLonger(ctx context.Context, in *GetLongerRequest, opts ...grpc.CallOption) (*GetLongerResponse, error) {
	out := new(GetLongerResponse)
	err := c.cc.Invoke(ctx, "/bio.ris.RoutingInformationService/GetLonger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingInformationServiceClient) ObserveRIB(ctx context.Context, in *ObserveRIBRequest, opts ...grpc.CallOption) (RoutingInformationService_ObserveRIBClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoutingInformationService_serviceDesc.Streams[0], "/bio.ris.RoutingInformationService/ObserveRIB", opts...)
	if err != nil {
		return nil, err
	}
	x := &routingInformationServiceObserveRIBClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoutingInformationService_ObserveRIBClient interface {
	Recv() (*RIBUpdate, error)
	grpc.ClientStream
}

type routingInformationServiceObserveRIBClient struct {
	grpc.ClientStream
}

func (x *routingInformationServiceObserveRIBClient) Recv() (*RIBUpdate, error) {
	m := new(RIBUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routingInformationServiceClient) DumpRIB(ctx context.Context, in *DumpRIBRequest, opts ...grpc.CallOption) (RoutingInformationService_DumpRIBClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoutingInformationService_serviceDesc.Streams[1], "/bio.ris.RoutingInformationService/DumpRIB", opts...)
	if err != nil {
		return nil, err
	}
	x := &routingInformationServiceDumpRIBClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoutingInformationService_DumpRIBClient interface {
	Recv() (*DumpRIBReply, error)
	grpc.ClientStream
}

type routingInformationServiceDumpRIBClient struct {
	grpc.ClientStream
}

func (x *routingInformationServiceDumpRIBClient) Recv() (*DumpRIBReply, error) {
	m := new(DumpRIBReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RoutingInformationServiceServer is the server API for RoutingInformationService service.
type RoutingInformationServiceServer interface {
	LPM(context.Context, *LPMRequest) (*LPMResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetRouters(context.Context, *GetRoutersRequest) (*GetRoutersResponse, error)
	GetLonger(context.Context, *GetLongerRequest) (*GetLongerResponse, error)
	ObserveRIB(*ObserveRIBRequest, RoutingInformationService_ObserveRIBServer) error
	DumpRIB(*DumpRIBRequest, RoutingInformationService_DumpRIBServer) error
}

func RegisterRoutingInformationServiceServer(s *grpc.Server, srv RoutingInformationServiceServer) {
	s.RegisterService(&_RoutingInformationService_serviceDesc, srv)
}

func _RoutingInformationService_LPM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LPMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingInformationServiceServer).LPM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bio.ris.RoutingInformationService/LPM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingInformationServiceServer).LPM(ctx, req.(*LPMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingInformationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingInformationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bio.ris.RoutingInformationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingInformationServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingInformationService_GetRouters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoutersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingInformationServiceServer).GetRouters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bio.ris.RoutingInformationService/GetRouters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingInformationServiceServer).GetRouters(ctx, req.(*GetRoutersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingInformationService_GetLonger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLongerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingInformationServiceServer).GetLonger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bio.ris.RoutingInformationService/GetLonger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingInformationServiceServer).GetLonger(ctx, req.(*GetLongerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingInformationService_ObserveRIB_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ObserveRIBRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoutingInformationServiceServer).ObserveRIB(m, &routingInformationServiceObserveRIBServer{stream})
}

type RoutingInformationService_ObserveRIBServer interface {
	Send(*RIBUpdate) error
	grpc.ServerStream
}

type routingInformationServiceObserveRIBServer struct {
	grpc.ServerStream
}

func (x *routingInformationServiceObserveRIBServer) Send(m *RIBUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _RoutingInformationService_DumpRIB_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DumpRIBRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoutingInformationServiceServer).DumpRIB(m, &routingInformationServiceDumpRIBServer{stream})
}

type RoutingInformationService_DumpRIBServer interface {
	Send(*DumpRIBReply) error
	grpc.ServerStream
}

type routingInformationServiceDumpRIBServer struct {
	grpc.ServerStream
}

func (x *routingInformationServiceDumpRIBServer) Send(m *DumpRIBReply) error {
	return x.ServerStream.SendMsg(m)
}

var _RoutingInformationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bio.ris.RoutingInformationService",
	HandlerType: (*RoutingInformationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LPM",
			Handler:    _RoutingInformationService_LPM_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RoutingInformationService_Get_Handler,
		},
		{
			MethodName: "GetRouters",
			Handler:    _RoutingInformationService_GetRouters_Handler,
		},
		{
			MethodName: "GetLonger",
			Handler:    _RoutingInformationService_GetLonger_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ObserveRIB",
			Handler:       _RoutingInformationService_ObserveRIB_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DumpRIB",
			Handler:       _RoutingInformationService_DumpRIB_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/bio-routing/bio-rd/cmd/ris/api/ris.proto",
}
