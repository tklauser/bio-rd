// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocols/bgp/server/api/session.proto

package api // import "protocols/bgp/server/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import api "net/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Session_State int32

const (
	Session_Disabled      Session_State = 0
	Session_Idle          Session_State = 1
	Session_Connect       Session_State = 2
	Session_Active        Session_State = 3
	Session_OpenSent      Session_State = 4
	Session_OpenConfirmed Session_State = 5
	Session_Established   Session_State = 6
	Session_Ceased        Session_State = 7
)

var Session_State_name = map[int32]string{
	0: "Disabled",
	1: "Idle",
	2: "Connect",
	3: "Active",
	4: "OpenSent",
	5: "OpenConfirmed",
	6: "Established",
	7: "Ceased",
}
var Session_State_value = map[string]int32{
	"Disabled":      0,
	"Idle":          1,
	"Connect":       2,
	"Active":        3,
	"OpenSent":      4,
	"OpenConfirmed": 5,
	"Established":   6,
	"Ceased":        7,
}

func (x Session_State) String() string {
	return proto.EnumName(Session_State_name, int32(x))
}
func (Session_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_session_ac0c55083006f27c, []int{0, 0}
}

type Session struct {
	LocalAddress         *api.IP       `protobuf:"bytes,1,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	NeighborAddress      *api.IP       `protobuf:"bytes,2,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	LocalAsn             uint32        `protobuf:"varint,3,opt,name=local_asn,json=localAsn,proto3" json:"local_asn,omitempty"`
	PeerAsn              uint32        `protobuf:"varint,4,opt,name=peer_asn,json=peerAsn,proto3" json:"peer_asn,omitempty"`
	Status               Session_State `protobuf:"varint,5,opt,name=status,proto3,enum=bio.bgp.Session_State" json:"status,omitempty"`
	Stats                *SessionStats `protobuf:"bytes,6,opt,name=stats,proto3" json:"stats,omitempty"`
	EstablishedSince     uint64        `protobuf:"varint,7,opt,name=established_since,json=establishedSince,proto3" json:"established_since,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_ac0c55083006f27c, []int{0}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (dst *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(dst, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetLocalAddress() *api.IP {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *Session) GetNeighborAddress() *api.IP {
	if m != nil {
		return m.NeighborAddress
	}
	return nil
}

func (m *Session) GetLocalAsn() uint32 {
	if m != nil {
		return m.LocalAsn
	}
	return 0
}

func (m *Session) GetPeerAsn() uint32 {
	if m != nil {
		return m.PeerAsn
	}
	return 0
}

func (m *Session) GetStatus() Session_State {
	if m != nil {
		return m.Status
	}
	return Session_Disabled
}

func (m *Session) GetStats() *SessionStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Session) GetEstablishedSince() uint64 {
	if m != nil {
		return m.EstablishedSince
	}
	return 0
}

type SessionStats struct {
	PackagesIn           uint64   `protobuf:"varint,6,opt,name=packages_in,json=packagesIn,proto3" json:"packages_in,omitempty"`
	PackagesOut          uint64   `protobuf:"varint,7,opt,name=packages_out,json=packagesOut,proto3" json:"packages_out,omitempty"`
	Flaps                uint64   `protobuf:"varint,8,opt,name=flaps,proto3" json:"flaps,omitempty"`
	RoutesReceived       uint64   `protobuf:"varint,9,opt,name=routes_received,json=routesReceived,proto3" json:"routes_received,omitempty"`
	RoutesImported       uint64   `protobuf:"varint,10,opt,name=routes_imported,json=routesImported,proto3" json:"routes_imported,omitempty"`
	RoutesExported       uint64   `protobuf:"varint,11,opt,name=routes_exported,json=routesExported,proto3" json:"routes_exported,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionStats) Reset()         { *m = SessionStats{} }
func (m *SessionStats) String() string { return proto.CompactTextString(m) }
func (*SessionStats) ProtoMessage()    {}
func (*SessionStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_ac0c55083006f27c, []int{1}
}
func (m *SessionStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionStats.Unmarshal(m, b)
}
func (m *SessionStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionStats.Marshal(b, m, deterministic)
}
func (dst *SessionStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionStats.Merge(dst, src)
}
func (m *SessionStats) XXX_Size() int {
	return xxx_messageInfo_SessionStats.Size(m)
}
func (m *SessionStats) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionStats.DiscardUnknown(m)
}

var xxx_messageInfo_SessionStats proto.InternalMessageInfo

func (m *SessionStats) GetPackagesIn() uint64 {
	if m != nil {
		return m.PackagesIn
	}
	return 0
}

func (m *SessionStats) GetPackagesOut() uint64 {
	if m != nil {
		return m.PackagesOut
	}
	return 0
}

func (m *SessionStats) GetFlaps() uint64 {
	if m != nil {
		return m.Flaps
	}
	return 0
}

func (m *SessionStats) GetRoutesReceived() uint64 {
	if m != nil {
		return m.RoutesReceived
	}
	return 0
}

func (m *SessionStats) GetRoutesImported() uint64 {
	if m != nil {
		return m.RoutesImported
	}
	return 0
}

func (m *SessionStats) GetRoutesExported() uint64 {
	if m != nil {
		return m.RoutesExported
	}
	return 0
}

func init() {
	proto.RegisterType((*Session)(nil), "bio.bgp.Session")
	proto.RegisterType((*SessionStats)(nil), "bio.bgp.SessionStats")
	proto.RegisterEnum("bio.bgp.Session_State", Session_State_name, Session_State_value)
}

func init() {
	proto.RegisterFile("protocols/bgp/server/api/session.proto", fileDescriptor_session_ac0c55083006f27c)
}

var fileDescriptor_session_ac0c55083006f27c = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0x9b, 0xd8, 0xee, 0x38, 0x69, 0x36, 0x2b, 0x40, 0xa6, 0x1c, 0x08, 0x39, 0x40,
	0xa4, 0x4a, 0x0e, 0x2a, 0x12, 0xf7, 0x12, 0x7a, 0xc8, 0xa9, 0xc8, 0xb9, 0x71, 0x89, 0xd6, 0xf6,
	0x34, 0x5d, 0xe1, 0xee, 0x5a, 0x3b, 0x9b, 0x88, 0x97, 0xe2, 0xa9, 0x78, 0x11, 0xb4, 0x6b, 0x27,
	0x58, 0x20, 0x6e, 0x99, 0xff, 0xfb, 0xfe, 0xcc, 0x8e, 0x64, 0x78, 0xd7, 0x18, 0x6d, 0x75, 0xa9,
	0x6b, 0x5a, 0x16, 0xbb, 0x66, 0x49, 0x68, 0x0e, 0x68, 0x96, 0xa2, 0x91, 0x4b, 0x42, 0x22, 0xa9,
	0x55, 0xe6, 0x05, 0x1e, 0x15, 0x52, 0x67, 0xc5, 0xae, 0xb9, 0x9a, 0x2a, 0xb4, 0x9e, 0x2b, 0xb4,
	0x2d, 0x9b, 0xff, 0x3c, 0x87, 0x68, 0xd3, 0xda, 0xfc, 0x03, 0x8c, 0x6b, 0x5d, 0x8a, 0x7a, 0x2b,
	0xaa, 0xca, 0x20, 0x51, 0x1a, 0xcc, 0x82, 0x45, 0x72, 0x93, 0x64, 0xae, 0xef, 0x2a, 0xeb, 0xaf,
	0xf9, 0xc8, 0x1b, 0xb7, 0xad, 0xc0, 0x3f, 0x01, 0x53, 0x28, 0x77, 0x8f, 0x85, 0x36, 0xa7, 0xd2,
	0xd9, 0xbf, 0xa5, 0xc9, 0x51, 0x3a, 0xf6, 0x5e, 0xc3, 0x45, 0xb7, 0x89, 0x54, 0x7a, 0x3e, 0x0b,
	0x16, 0xe3, 0x3c, 0x6e, 0xff, 0x98, 0x14, 0x7f, 0x05, 0x71, 0x83, 0x68, 0x3c, 0x1b, 0x78, 0x16,
	0xb9, 0xd9, 0xa1, 0x0c, 0x42, 0xb2, 0xc2, 0xee, 0x29, 0x1d, 0xce, 0x82, 0xc5, 0xe5, 0xcd, 0xcb,
	0xac, 0x3b, 0x2d, 0xeb, 0x6e, 0xc8, 0x36, 0x56, 0x58, 0xcc, 0x3b, 0x8b, 0x5f, 0xc3, 0xd0, 0xfd,
	0xa2, 0x34, 0xf4, 0x8f, 0x7a, 0xf1, 0xb7, 0xee, 0x6c, 0xca, 0x5b, 0x87, 0x5f, 0xc3, 0x14, 0xc9,
	0x8a, 0xa2, 0x96, 0xf4, 0x88, 0xd5, 0x96, 0xa4, 0x2a, 0x31, 0x8d, 0x66, 0xc1, 0x62, 0x90, 0xb3,
	0x1e, 0xd8, 0xb8, 0x7c, 0x7e, 0x80, 0xa1, 0x5f, 0xc5, 0x47, 0x10, 0x7f, 0x91, 0x24, 0x8a, 0x1a,
	0x2b, 0xf6, 0x8c, 0xc7, 0x30, 0x58, 0x57, 0x35, 0xb2, 0x80, 0x27, 0x10, 0xad, 0xb4, 0x52, 0x58,
	0x5a, 0x76, 0xc6, 0x01, 0xc2, 0xdb, 0xd2, 0xca, 0x03, 0xb2, 0x73, 0x57, 0xb8, 0x6f, 0x50, 0x6d,
	0x50, 0x59, 0x36, 0xe0, 0x53, 0x18, 0xbb, 0x69, 0xa5, 0xd5, 0x83, 0x34, 0x4f, 0x58, 0xb1, 0x21,
	0x9f, 0x40, 0x72, 0xf7, 0x67, 0x1d, 0x0b, 0x5d, 0x7b, 0x85, 0x82, 0xb0, 0x62, 0xd1, 0xfc, 0x57,
	0x00, 0xa3, 0xfe, 0xe3, 0xf9, 0x1b, 0x48, 0x1a, 0x51, 0x7e, 0x17, 0x3b, 0xa4, 0xad, 0x54, 0xfe,
	0xd0, 0x41, 0x0e, 0xc7, 0x68, 0xad, 0xf8, 0x5b, 0x18, 0x9d, 0x04, 0xbd, 0xb7, 0xdd, 0x45, 0xa7,
	0xd2, 0xfd, 0xde, 0xf2, 0xe7, 0x30, 0x7c, 0xa8, 0x45, 0x43, 0x69, 0xec, 0x59, 0x3b, 0xf0, 0xf7,
	0x30, 0x31, 0x7a, 0x6f, 0x91, 0xb6, 0x06, 0x4b, 0x94, 0x07, 0xac, 0xd2, 0x0b, 0xcf, 0x2f, 0xdb,
	0x38, 0xef, 0xd2, 0x9e, 0x28, 0x9f, 0x1a, 0x6d, 0x2c, 0x56, 0x29, 0xf4, 0xc5, 0x75, 0x97, 0xf6,
	0x44, 0xfc, 0xd1, 0x89, 0x49, 0x5f, 0xbc, 0xeb, 0xd2, 0xcf, 0x57, 0xdf, 0xd2, 0xff, 0x7d, 0xdb,
	0x45, 0xe8, 0xc9, 0xc7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x0e, 0xc1, 0xc1, 0xfe, 0x02,
	0x00, 0x00,
}
