// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: route/api/route.proto

package api

import (
	api "github.com/bio-routing/bio-rd/net/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Path_Type int32

const (
	Path_Static Path_Type = 0
	Path_BGP    Path_Type = 1
	Path_GRP    Path_Type = 2
)

// Enum value maps for Path_Type.
var (
	Path_Type_name = map[int32]string{
		0: "Static",
		1: "BGP",
		2: "GRP",
	}
	Path_Type_value = map[string]int32{
		"Static": 0,
		"BGP":    1,
		"GRP":    2,
	}
)

func (x Path_Type) Enum() *Path_Type {
	p := new(Path_Type)
	*p = x
	return p
}

func (x Path_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Path_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_route_api_route_proto_enumTypes[0].Descriptor()
}

func (Path_Type) Type() protoreflect.EnumType {
	return &file_route_api_route_proto_enumTypes[0]
}

func (x Path_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Path_Type.Descriptor instead.
func (Path_Type) EnumDescriptor() ([]byte, []int) {
	return file_route_api_route_proto_rawDescGZIP(), []int{1, 0}
}

type Path_HiddenReason int32

const (
	Path_HiddenReasonNone               Path_HiddenReason = 0
	Path_HiddenReasonNextHopUnreachable Path_HiddenReason = 1
	Path_HiddenReasonFilteredByPolicy   Path_HiddenReason = 2
	Path_HiddenReasonASLoop             Path_HiddenReason = 3
	Path_HiddenReasonOurOriginatorID    Path_HiddenReason = 4
	Path_HiddenReasonClusterLoop        Path_HiddenReason = 5
	Path_HiddenReasonOTCMismatch        Path_HiddenReason = 6
)

// Enum value maps for Path_HiddenReason.
var (
	Path_HiddenReason_name = map[int32]string{
		0: "HiddenReasonNone",
		1: "HiddenReasonNextHopUnreachable",
		2: "HiddenReasonFilteredByPolicy",
		3: "HiddenReasonASLoop",
		4: "HiddenReasonOurOriginatorID",
		5: "HiddenReasonClusterLoop",
		6: "HiddenReasonOTCMismatch",
	}
	Path_HiddenReason_value = map[string]int32{
		"HiddenReasonNone":               0,
		"HiddenReasonNextHopUnreachable": 1,
		"HiddenReasonFilteredByPolicy":   2,
		"HiddenReasonASLoop":             3,
		"HiddenReasonOurOriginatorID":    4,
		"HiddenReasonClusterLoop":        5,
		"HiddenReasonOTCMismatch":        6,
	}
)

func (x Path_HiddenReason) Enum() *Path_HiddenReason {
	p := new(Path_HiddenReason)
	*p = x
	return p
}

func (x Path_HiddenReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Path_HiddenReason) Descriptor() protoreflect.EnumDescriptor {
	return file_route_api_route_proto_enumTypes[1].Descriptor()
}

func (Path_HiddenReason) Type() protoreflect.EnumType {
	return &file_route_api_route_proto_enumTypes[1]
}

func (x Path_HiddenReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Path_HiddenReason.Descriptor instead.
func (Path_HiddenReason) EnumDescriptor() ([]byte, []int) {
	return file_route_api_route_proto_rawDescGZIP(), []int{1, 1}
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pfx   *api.Prefix `protobuf:"bytes,1,opt,name=pfx,proto3" json:"pfx,omitempty"`
	Paths []*Path     `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_api_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_route_api_route_proto_rawDescGZIP(), []int{0}
}

func (x *Route) GetPfx() *api.Prefix {
	if x != nil {
		return x.Pfx
	}
	return nil
}

func (x *Route) GetPaths() []*Path {
	if x != nil {
		return x.Paths
	}
	return nil
}

type Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         Path_Type         `protobuf:"varint,1,opt,name=type,proto3,enum=bio.route.Path_Type" json:"type,omitempty"`
	StaticPath   *StaticPath       `protobuf:"bytes,2,opt,name=static_path,json=staticPath,proto3" json:"static_path,omitempty"`
	BgpPath      *BGPPath          `protobuf:"bytes,3,opt,name=bgp_path,json=bgpPath,proto3" json:"bgp_path,omitempty"`
	HiddenReason Path_HiddenReason `protobuf:"varint,4,opt,name=hidden_reason,json=hiddenReason,proto3,enum=bio.route.Path_HiddenReason" json:"hidden_reason,omitempty"`
	TimeLearned  uint32            `protobuf:"varint,5,opt,name=time_learned,json=timeLearned,proto3" json:"time_learned,omitempty"`
	GrpPath      *GRPPath          `protobuf:"bytes,6,opt,name=grp_path,json=grpPath,proto3" json:"grp_path,omitempty"`
}

func (x *Path) Reset() {
	*x = Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_api_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path) ProtoMessage() {}

func (x *Path) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path.ProtoReflect.Descriptor instead.
func (*Path) Descriptor() ([]byte, []int) {
	return file_route_api_route_proto_rawDescGZIP(), []int{1}
}

func (x *Path) GetType() Path_Type {
	if x != nil {
		return x.Type
	}
	return Path_Static
}

func (x *Path) GetStaticPath() *StaticPath {
	if x != nil {
		return x.StaticPath
	}
	return nil
}

func (x *Path) GetBgpPath() *BGPPath {
	if x != nil {
		return x.BgpPath
	}
	return nil
}

func (x *Path) GetHiddenReason() Path_HiddenReason {
	if x != nil {
		return x.HiddenReason
	}
	return Path_HiddenReasonNone
}

func (x *Path) GetTimeLearned() uint32 {
	if x != nil {
		return x.TimeLearned
	}
	return 0
}

func (x *Path) GetGrpPath() *GRPPath {
	if x != nil {
		return x.GrpPath
	}
	return nil
}

type StaticPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextHop *api.IP `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
}

func (x *StaticPath) Reset() {
	*x = StaticPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_api_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticPath) ProtoMessage() {}

func (x *StaticPath) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticPath.ProtoReflect.Descriptor instead.
func (*StaticPath) Descriptor() ([]byte, []int) {
	return file_route_api_route_proto_rawDescGZIP(), []int{2}
}

func (x *StaticPath) GetNextHop() *api.IP {
	if x != nil {
		return x.NextHop
	}
	return nil
}

type GRPPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextHop  *api.IP           `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	MetaData map[string]string `protobuf:"bytes,2,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GRPPath) Reset() {
	*x = GRPPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_api_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPPath) ProtoMessage() {}

func (x *GRPPath) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPPath.ProtoReflect.Descriptor instead.
func (*GRPPath) Descriptor() ([]byte, []int) {
	return file_route_api_route_proto_rawDescGZIP(), []int{3}
}

func (x *GRPPath) GetNextHop() *api.IP {
	if x != nil {
		return x.NextHop
	}
	return nil
}

func (x *GRPPath) GetMetaData() map[string]string {
	if x != nil {
		return x.MetaData
	}
	return nil
}

type BGPPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PathIdentifier    uint32                  `protobuf:"varint,1,opt,name=path_identifier,json=pathIdentifier,proto3" json:"path_identifier,omitempty"`
	NextHop           *api.IP                 `protobuf:"bytes,2,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	LocalPref         uint32                  `protobuf:"varint,3,opt,name=local_pref,json=localPref,proto3" json:"local_pref,omitempty"`
	AsPath            []*ASPathSegment        `protobuf:"bytes,4,rep,name=as_path,json=asPath,proto3" json:"as_path,omitempty"`
	Origin            uint32                  `protobuf:"varint,5,opt,name=origin,proto3" json:"origin,omitempty"`
	Med               uint32                  `protobuf:"varint,6,opt,name=med,proto3" json:"med,omitempty"`
	Ebgp              bool                    `protobuf:"varint,7,opt,name=ebgp,proto3" json:"ebgp,omitempty"`
	BgpIdentifier     uint32                  `protobuf:"varint,8,opt,name=bgp_identifier,json=bgpIdentifier,proto3" json:"bgp_identifier,omitempty"`
	Source            *api.IP                 `protobuf:"bytes,9,opt,name=source,proto3" json:"source,omitempty"`
	Communities       []uint32                `protobuf:"varint,10,rep,packed,name=communities,proto3" json:"communities,omitempty"`
	LargeCommunities  []*LargeCommunity       `protobuf:"bytes,11,rep,name=large_communities,json=largeCommunities,proto3" json:"large_communities,omitempty"`
	OriginatorId      uint32                  `protobuf:"varint,12,opt,name=originator_id,json=originatorId,proto3" json:"originator_id,omitempty"`
	ClusterList       []uint32                `protobuf:"varint,13,rep,packed,name=cluster_list,json=clusterList,proto3" json:"cluster_list,omitempty"`
	UnknownAttributes []*UnknownPathAttribute `protobuf:"bytes,14,rep,name=unknown_attributes,json=unknownAttributes,proto3" json:"unknown_attributes,omitempty"`
	BmpPostPolicy     bool                    `protobuf:"varint,15,opt,name=bmp_post_policy,json=bmpPostPolicy,proto3" json:"bmp_post_policy,omitempty"`
	OnlyToCustomer    uint32                  `protobuf:"varint,16,opt,name=only_to_customer,json=onlyToCustomer,proto3" json:"only_to_customer,omitempty"`
}

func (x *BGPPath) Reset() {
	*x = BGPPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_api_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BGPPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BGPPath) ProtoMessage() {}

func (x *BGPPath) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BGPPath.ProtoReflect.Descriptor instead.
func (*BGPPath) Descriptor() ([]byte, []int) {
	return file_route_api_route_proto_rawDescGZIP(), []int{4}
}

func (x *BGPPath) GetPathIdentifier() uint32 {
	if x != nil {
		return x.PathIdentifier
	}
	return 0
}

func (x *BGPPath) GetNextHop() *api.IP {
	if x != nil {
		return x.NextHop
	}
	return nil
}

func (x *BGPPath) GetLocalPref() uint32 {
	if x != nil {
		return x.LocalPref
	}
	return 0
}

func (x *BGPPath) GetAsPath() []*ASPathSegment {
	if x != nil {
		return x.AsPath
	}
	return nil
}

func (x *BGPPath) GetOrigin() uint32 {
	if x != nil {
		return x.Origin
	}
	return 0
}

func (x *BGPPath) GetMed() uint32 {
	if x != nil {
		return x.Med
	}
	return 0
}

func (x *BGPPath) GetEbgp() bool {
	if x != nil {
		return x.Ebgp
	}
	return false
}

func (x *BGPPath) GetBgpIdentifier() uint32 {
	if x != nil {
		return x.BgpIdentifier
	}
	return 0
}

func (x *BGPPath) GetSource() *api.IP {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *BGPPath) GetCommunities() []uint32 {
	if x != nil {
		return x.Communities
	}
	return nil
}

func (x *BGPPath) GetLargeCommunities() []*LargeCommunity {
	if x != nil {
		return x.LargeCommunities
	}
	return nil
}

func (x *BGPPath) GetOriginatorId() uint32 {
	if x != nil {
		return x.OriginatorId
	}
	return 0
}

func (x *BGPPath) GetClusterList() []uint32 {
	if x != nil {
		return x.ClusterList
	}
	return nil
}

func (x *BGPPath) GetUnknownAttributes() []*UnknownPathAttribute {
	if x != nil {
		return x.UnknownAttributes
	}
	return nil
}

func (x *BGPPath) GetBmpPostPolicy() bool {
	if x != nil {
		return x.BmpPostPolicy
	}
	return false
}

func (x *BGPPath) GetOnlyToCustomer() uint32 {
	if x != nil {
		return x.OnlyToCustomer
	}
	return 0
}

type ASPathSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AsSequence bool     `protobuf:"varint,1,opt,name=as_sequence,json=asSequence,proto3" json:"as_sequence,omitempty"`
	Asns       []uint32 `protobuf:"varint,2,rep,packed,name=asns,proto3" json:"asns,omitempty"`
}

func (x *ASPathSegment) Reset() {
	*x = ASPathSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_api_route_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ASPathSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ASPathSegment) ProtoMessage() {}

func (x *ASPathSegment) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ASPathSegment.ProtoReflect.Descriptor instead.
func (*ASPathSegment) Descriptor() ([]byte, []int) {
	return file_route_api_route_proto_rawDescGZIP(), []int{5}
}

func (x *ASPathSegment) GetAsSequence() bool {
	if x != nil {
		return x.AsSequence
	}
	return false
}

func (x *ASPathSegment) GetAsns() []uint32 {
	if x != nil {
		return x.Asns
	}
	return nil
}

type LargeCommunity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlobalAdministrator uint32 `protobuf:"varint,1,opt,name=global_administrator,json=globalAdministrator,proto3" json:"global_administrator,omitempty"`
	DataPart1           uint32 `protobuf:"varint,2,opt,name=data_part1,json=dataPart1,proto3" json:"data_part1,omitempty"`
	DataPart2           uint32 `protobuf:"varint,3,opt,name=data_part2,json=dataPart2,proto3" json:"data_part2,omitempty"`
}

func (x *LargeCommunity) Reset() {
	*x = LargeCommunity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_api_route_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LargeCommunity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LargeCommunity) ProtoMessage() {}

func (x *LargeCommunity) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LargeCommunity.ProtoReflect.Descriptor instead.
func (*LargeCommunity) Descriptor() ([]byte, []int) {
	return file_route_api_route_proto_rawDescGZIP(), []int{6}
}

func (x *LargeCommunity) GetGlobalAdministrator() uint32 {
	if x != nil {
		return x.GlobalAdministrator
	}
	return 0
}

func (x *LargeCommunity) GetDataPart1() uint32 {
	if x != nil {
		return x.DataPart1
	}
	return 0
}

func (x *LargeCommunity) GetDataPart2() uint32 {
	if x != nil {
		return x.DataPart2
	}
	return 0
}

type UnknownPathAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Optional   bool   `protobuf:"varint,1,opt,name=optional,proto3" json:"optional,omitempty"`
	Transitive bool   `protobuf:"varint,2,opt,name=transitive,proto3" json:"transitive,omitempty"`
	Partial    bool   `protobuf:"varint,3,opt,name=partial,proto3" json:"partial,omitempty"`
	TypeCode   uint32 `protobuf:"varint,4,opt,name=type_code,json=typeCode,proto3" json:"type_code,omitempty"`
	Value      []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UnknownPathAttribute) Reset() {
	*x = UnknownPathAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_api_route_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnknownPathAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnknownPathAttribute) ProtoMessage() {}

func (x *UnknownPathAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_route_api_route_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnknownPathAttribute.ProtoReflect.Descriptor instead.
func (*UnknownPathAttribute) Descriptor() ([]byte, []int) {
	return file_route_api_route_proto_rawDescGZIP(), []int{7}
}

func (x *UnknownPathAttribute) GetOptional() bool {
	if x != nil {
		return x.Optional
	}
	return false
}

func (x *UnknownPathAttribute) GetTransitive() bool {
	if x != nil {
		return x.Transitive
	}
	return false
}

func (x *UnknownPathAttribute) GetPartial() bool {
	if x != nil {
		return x.Partial
	}
	return false
}

func (x *UnknownPathAttribute) GetTypeCode() uint32 {
	if x != nil {
		return x.TypeCode
	}
	return 0
}

func (x *UnknownPathAttribute) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_route_api_route_proto protoreflect.FileDescriptor

var file_route_api_route_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x1a, 0x11, 0x6e, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x21,
	0x0a, 0x03, 0x70, 0x66, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x69,
	0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x03, 0x70, 0x66,
	0x78, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x74,
	0x68, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x22, 0xb2, 0x04, 0x0a, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x68,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x50, 0x61, 0x74, 0x68, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x2d, 0x0a, 0x08, 0x62, 0x67, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x42, 0x47, 0x50, 0x50, 0x61, 0x74, 0x68, 0x52, 0x07, 0x62, 0x67, 0x70, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x41, 0x0a, 0x0d, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x69, 0x6f, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x2e, 0x48, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x0c, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x69, 0x6d,
	0x65, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x08, 0x67, 0x72, 0x70, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x6f,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x47, 0x52, 0x50, 0x50, 0x61, 0x74, 0x68, 0x52, 0x07,
	0x67, 0x72, 0x70, 0x50, 0x61, 0x74, 0x68, 0x22, 0x24, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42,
	0x47, 0x50, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x52, 0x50, 0x10, 0x02, 0x22, 0xdd, 0x01,
	0x0a, 0x0c, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x10, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x6f,
	0x6e, 0x65, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x55, 0x6e, 0x72, 0x65, 0x61,
	0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x48, 0x69, 0x64, 0x64,
	0x65, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x42, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x41, 0x53, 0x4c, 0x6f, 0x6f, 0x70,
	0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x4f, 0x75, 0x72, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x44, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x6f, 0x70, 0x10, 0x05,
	0x12, 0x1b, 0x0a, 0x17, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x4f, 0x54, 0x43, 0x4d, 0x69, 0x73, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x06, 0x22, 0x34, 0x0a,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x61, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x08, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x62, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74,
	0x48, 0x6f, 0x70, 0x22, 0xad, 0x01, 0x0a, 0x07, 0x47, 0x52, 0x50, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x26, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x52, 0x07,
	0x6e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x12, 0x3d, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x69, 0x6f,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x47, 0x52, 0x50, 0x50, 0x61, 0x74, 0x68, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x8a, 0x05, 0x0a, 0x07, 0x42, 0x47, 0x50, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x61, 0x74, 0x68, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x68, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6f,
	0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x72, 0x65, 0x66, 0x12,
	0x31, 0x0a, 0x07, 0x61, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x41, 0x53, 0x50,
	0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x73, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6d, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x65, 0x62, 0x67, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x65, 0x62, 0x67, 0x70,
	0x12, 0x25, 0x0a, 0x0e, 0x62, 0x67, 0x70, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x62, 0x67, 0x70, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x6e, 0x65,
	0x74, 0x2e, 0x49, 0x50, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x46,
	0x0a, 0x11, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x69, 0x6f, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x52, 0x10, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4e,
	0x0a, 0x12, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x69, 0x6f,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x50, 0x61,
	0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x11, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x62, 0x6d, 0x70, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x62, 0x6d, 0x70, 0x50, 0x6f, 0x73, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x74,
	0x6f, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x6f, 0x6e, 0x6c, 0x79, 0x54, 0x6f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x22, 0x44, 0x0a, 0x0d, 0x41, 0x53, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x73, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x73, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x04, 0x61, 0x73, 0x6e, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x0e, 0x4c, 0x61, 0x72, 0x67, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x14, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x74, 0x31, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x74, 0x32, 0x22, 0x9f, 0x01, 0x0a, 0x14, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x79,
	0x70, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x29, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6f, 0x2d, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x69, 0x6f, 0x2d, 0x72, 0x64, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_route_api_route_proto_rawDescOnce sync.Once
	file_route_api_route_proto_rawDescData = file_route_api_route_proto_rawDesc
)

func file_route_api_route_proto_rawDescGZIP() []byte {
	file_route_api_route_proto_rawDescOnce.Do(func() {
		file_route_api_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_route_api_route_proto_rawDescData)
	})
	return file_route_api_route_proto_rawDescData
}

var file_route_api_route_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_route_api_route_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_route_api_route_proto_goTypes = []interface{}{
	(Path_Type)(0),               // 0: bio.route.Path.Type
	(Path_HiddenReason)(0),       // 1: bio.route.Path.HiddenReason
	(*Route)(nil),                // 2: bio.route.Route
	(*Path)(nil),                 // 3: bio.route.Path
	(*StaticPath)(nil),           // 4: bio.route.StaticPath
	(*GRPPath)(nil),              // 5: bio.route.GRPPath
	(*BGPPath)(nil),              // 6: bio.route.BGPPath
	(*ASPathSegment)(nil),        // 7: bio.route.ASPathSegment
	(*LargeCommunity)(nil),       // 8: bio.route.LargeCommunity
	(*UnknownPathAttribute)(nil), // 9: bio.route.UnknownPathAttribute
	nil,                          // 10: bio.route.GRPPath.MetaDataEntry
	(*api.Prefix)(nil),           // 11: bio.net.Prefix
	(*api.IP)(nil),               // 12: bio.net.IP
}
var file_route_api_route_proto_depIdxs = []int32{
	11, // 0: bio.route.Route.pfx:type_name -> bio.net.Prefix
	3,  // 1: bio.route.Route.paths:type_name -> bio.route.Path
	0,  // 2: bio.route.Path.type:type_name -> bio.route.Path.Type
	4,  // 3: bio.route.Path.static_path:type_name -> bio.route.StaticPath
	6,  // 4: bio.route.Path.bgp_path:type_name -> bio.route.BGPPath
	1,  // 5: bio.route.Path.hidden_reason:type_name -> bio.route.Path.HiddenReason
	5,  // 6: bio.route.Path.grp_path:type_name -> bio.route.GRPPath
	12, // 7: bio.route.StaticPath.next_hop:type_name -> bio.net.IP
	12, // 8: bio.route.GRPPath.next_hop:type_name -> bio.net.IP
	10, // 9: bio.route.GRPPath.meta_data:type_name -> bio.route.GRPPath.MetaDataEntry
	12, // 10: bio.route.BGPPath.next_hop:type_name -> bio.net.IP
	7,  // 11: bio.route.BGPPath.as_path:type_name -> bio.route.ASPathSegment
	12, // 12: bio.route.BGPPath.source:type_name -> bio.net.IP
	8,  // 13: bio.route.BGPPath.large_communities:type_name -> bio.route.LargeCommunity
	9,  // 14: bio.route.BGPPath.unknown_attributes:type_name -> bio.route.UnknownPathAttribute
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_route_api_route_proto_init() }
func file_route_api_route_proto_init() {
	if File_route_api_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_route_api_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_api_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_api_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticPath); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_api_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPPath); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_api_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BGPPath); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_api_route_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ASPathSegment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_api_route_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LargeCommunity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_api_route_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnknownPathAttribute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_route_api_route_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_route_api_route_proto_goTypes,
		DependencyIndexes: file_route_api_route_proto_depIdxs,
		EnumInfos:         file_route_api_route_proto_enumTypes,
		MessageInfos:      file_route_api_route_proto_msgTypes,
	}.Build()
	File_route_api_route_proto = out.File
	file_route_api_route_proto_rawDesc = nil
	file_route_api_route_proto_goTypes = nil
	file_route_api_route_proto_depIdxs = nil
}
