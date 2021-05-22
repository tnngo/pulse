// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: packet.proto

package packet

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Type int32

const (
	// Reserved.
	Type_Unknown Type = 0
	// Client request to connect to server.
	Type_Connect Type = 1
	// Connect Acknowledgment.
	Type_ConnAck Type = 2
	// Client sends a heartbeat to the server.
	Type_Ping Type = 3
	// Server responds to the client heartbeat
	Type_Pong Type = 4
	// Msg content.
	Type_Body Type = 99
)

var Type_name = map[int32]string{
	0:  "Unknown",
	1:  "Connect",
	2:  "ConnAck",
	3:  "Ping",
	4:  "Pong",
	99: "Body",
}

var Type_value = map[string]int32{
	"Unknown": 0,
	"Connect": 1,
	"ConnAck": 2,
	"Ping":    3,
	"Pong":    4,
	"Body":    99,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{0}
}

// AuthMode Type = Connect
type AuthMode int32

const (
	// PULSE default encryption method.
	// Base64(Hmac+Sha1).
	AuthMode_HmacSha1 AuthMode = 0
	AuthMode_NotSafe  AuthMode = 99
)

var AuthMode_name = map[int32]string{
	0:  "HmacSha1",
	99: "NotSafe",
}

var AuthMode_value = map[string]int32{
	"HmacSha1": 0,
	"NotSafe":  99,
}

func (x AuthMode) String() string {
	return proto.EnumName(AuthMode_name, int32(x))
}

func (AuthMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{1}
}

type RouteMode int32

const (
	RouteMode_Normal   RouteMode = 0
	RouteMode_Default  RouteMode = 1
	RouteMode_SelfProc RouteMode = 2
)

var RouteMode_name = map[int32]string{
	0: "Normal",
	1: "Default",
	2: "SelfProc",
}

var RouteMode_value = map[string]int32{
	"Normal":   0,
	"Default":  1,
	"SelfProc": 2,
}

func (x RouteMode) String() string {
	return proto.EnumName(RouteMode_name, int32(x))
}

func (RouteMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{2}
}

type Route struct {
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Id    int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{0}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Route.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return m.Size()
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Route) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Packet packet message structure.
type Packet struct {
	// Packet type.
	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=packet.Type" json:"type,omitempty"`
	// Unique identification of device terminal.
	Udid string `protobuf:"bytes,2,opt,name=udid,proto3" json:"udid,omitempty"`
	// Sometimes the same type of device will have different business scenarios,
	// so the applications in the device will be named differently.
	AppName string `protobuf:"bytes,3,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	//* Auth. Begin*
	AuthMode AuthMode `protobuf:"varint,50,opt,name=auth_mode,json=authMode,proto3,enum=packet.AuthMode" json:"auth_mode,omitempty"`
	Secret   string   `protobuf:"bytes,51,opt,name=secret,proto3" json:"secret,omitempty"`
	//* Other. Begin*
	LocalAddr string `protobuf:"bytes,70,opt,name=local_addr,json=localAddr,proto3" json:"local_addr,omitempty"`
	//* Route msg. Begin*
	RouteMode RouteMode `protobuf:"varint,90,opt,name=route_mode,json=routeMode,proto3,enum=packet.RouteMode" json:"route_mode,omitempty"`
	Route     *Route    `protobuf:"bytes,91,opt,name=route,proto3" json:"route,omitempty"`
	Msg       *Msg      `protobuf:"bytes,99,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{1}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return m.Size()
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_Unknown
}

func (m *Packet) GetUdid() string {
	if m != nil {
		return m.Udid
	}
	return ""
}

func (m *Packet) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Packet) GetAuthMode() AuthMode {
	if m != nil {
		return m.AuthMode
	}
	return AuthMode_HmacSha1
}

func (m *Packet) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Packet) GetLocalAddr() string {
	if m != nil {
		return m.LocalAddr
	}
	return ""
}

func (m *Packet) GetRouteMode() RouteMode {
	if m != nil {
		return m.RouteMode
	}
	return RouteMode_Normal
}

func (m *Packet) GetRoute() *Route {
	if m != nil {
		return m.Route
	}
	return nil
}

func (m *Packet) GetMsg() *Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

type Msg struct {
	RequestId string `protobuf:"bytes,98,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Body      []byte `protobuf:"bytes,99,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{2}
}
func (m *Msg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return m.Size()
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *Msg) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterEnum("packet.Type", Type_name, Type_value)
	proto.RegisterEnum("packet.AuthMode", AuthMode_name, AuthMode_value)
	proto.RegisterEnum("packet.RouteMode", RouteMode_name, RouteMode_value)
	proto.RegisterType((*Route)(nil), "packet.Route")
	proto.RegisterType((*Packet)(nil), "packet.Packet")
	proto.RegisterType((*Msg)(nil), "packet.Msg")
}

func init() { proto.RegisterFile("packet.proto", fileDescriptor_e9ef1a6541f9f9e7) }

var fileDescriptor_e9ef1a6541f9f9e7 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x6d, 0xd2, 0x34, 0x9b, 0x7c, 0x8d, 0x4b, 0x1c, 0x44, 0xe2, 0x61, 0x43, 0xa8, 0x08, 0xa5,
	0xb0, 0x8b, 0x76, 0x2f, 0x5e, 0xbb, 0x8a, 0x28, 0xd2, 0x52, 0x52, 0xbd, 0xac, 0x87, 0x32, 0x9d,
	0x99, 0xa6, 0xa5, 0x49, 0x26, 0x4e, 0x26, 0x48, 0xff, 0x85, 0x7f, 0xc6, 0xff, 0xe0, 0x71, 0x8f,
	0x1e, 0xa5, 0xfd, 0x23, 0x92, 0x2f, 0x89, 0xb0, 0xb7, 0xf7, 0xbe, 0xef, 0xcd, 0x7b, 0x33, 0x8f,
	0x01, 0xaf, 0xa0, 0xec, 0x20, 0xf4, 0x4d, 0xa1, 0xa4, 0x96, 0xc4, 0x6e, 0xd8, 0xe8, 0x1a, 0x06,
	0xb1, 0xac, 0xb4, 0x20, 0xcf, 0x60, 0x90, 0x28, 0x59, 0x15, 0x81, 0x11, 0x19, 0x63, 0x37, 0x6e,
	0x08, 0xb9, 0x04, 0x73, 0xcf, 0x03, 0x33, 0x32, 0xc6, 0x83, 0xd8, 0xdc, 0xf3, 0xd1, 0x2f, 0x13,
	0xec, 0x25, 0x9e, 0x24, 0x11, 0x58, 0xfa, 0x58, 0x08, 0xd4, 0x5f, 0x4e, 0xbd, 0x9b, 0xd6, 0xfe,
	0xcb, 0xb1, 0x10, 0x31, 0x6e, 0x08, 0x01, 0xab, 0xe2, 0xed, 0x71, 0x37, 0x46, 0x4c, 0x5e, 0x80,
	0x43, 0x8b, 0x62, 0x9d, 0xd3, 0x4c, 0x04, 0x7d, 0x9c, 0x5f, 0xd0, 0xa2, 0x58, 0xd0, 0x4c, 0x90,
	0x6b, 0x70, 0x69, 0xa5, 0x77, 0xeb, 0x4c, 0x72, 0x11, 0x4c, 0xd1, 0xd5, 0xef, 0x5c, 0x67, 0x95,
	0xde, 0xcd, 0x25, 0x17, 0xb1, 0x43, 0x5b, 0x44, 0x9e, 0x83, 0x5d, 0x0a, 0xa6, 0x84, 0x0e, 0x6e,
	0xd1, 0xa7, 0x65, 0xe4, 0x0a, 0x20, 0x95, 0x8c, 0xa6, 0x6b, 0xca, 0xb9, 0x0a, 0x3e, 0xe0, 0xce,
	0xc5, 0xc9, 0x8c, 0x73, 0x45, 0x5e, 0x03, 0xa8, 0xfa, 0xc1, 0x4d, 0xcc, 0x3d, 0xc6, 0x3c, 0xed,
	0x62, 0xb0, 0x0a, 0xcc, 0x71, 0x55, 0x07, 0xc9, 0x4b, 0x18, 0x20, 0x09, 0xbe, 0x45, 0xc6, 0x78,
	0x38, 0x7d, 0xf2, 0x48, 0x1c, 0x37, 0x3b, 0x72, 0x05, 0xfd, 0xac, 0x4c, 0x02, 0x86, 0x92, 0x61,
	0x27, 0x99, 0x97, 0x49, 0x5c, 0xcf, 0x47, 0x6f, 0xa1, 0x3f, 0x2f, 0x93, 0xfa, 0x6e, 0x4a, 0x7c,
	0xaf, 0x44, 0xa9, 0xd7, 0x7b, 0x1e, 0x6c, 0x9a, 0xbb, 0xb5, 0x93, 0x4f, 0xbc, 0x2e, 0x6c, 0x23,
	0xf9, 0x11, 0x5d, 0xbc, 0x18, 0xf1, 0xe4, 0x33, 0x58, 0x75, 0xa5, 0x64, 0x08, 0x17, 0x5f, 0xf3,
	0x43, 0x2e, 0x7f, 0xe4, 0x7e, 0xaf, 0x26, 0xef, 0x64, 0x9e, 0x0b, 0xa6, 0x7d, 0xa3, 0x23, 0x33,
	0x76, 0xf0, 0x4d, 0xe2, 0x80, 0xb5, 0xdc, 0xe7, 0x89, 0xdf, 0x47, 0x24, 0xf3, 0xc4, 0xb7, 0x6a,
	0x74, 0x27, 0xf9, 0xd1, 0x67, 0x93, 0x57, 0xe0, 0x74, 0x4d, 0x12, 0x0f, 0x9c, 0x8f, 0x19, 0x65,
	0xab, 0x1d, 0x7d, 0xd3, 0x38, 0x2e, 0xa4, 0x5e, 0xd1, 0xad, 0xf0, 0xd9, 0x64, 0x0a, 0xee, 0xff,
	0x26, 0x08, 0x80, 0xbd, 0x90, 0x2a, 0xa3, 0x69, 0xa3, 0x7a, 0x2f, 0xb6, 0xb4, 0x4a, 0xeb, 0x5c,
	0x0f, 0x9c, 0x95, 0x48, 0xb7, 0x4b, 0x25, 0x99, 0x6f, 0xde, 0x45, 0xbf, 0x4f, 0xa1, 0xf1, 0x70,
	0x0a, 0x8d, 0xbf, 0xa7, 0xd0, 0xf8, 0x79, 0x0e, 0x7b, 0x0f, 0xe7, 0xb0, 0xf7, 0xe7, 0x1c, 0xf6,
	0xee, 0xdb, 0xaf, 0xb6, 0xb1, 0xf1, 0xe7, 0xdd, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xeb,
	0xc9, 0x78, 0x89, 0x02, 0x00, 0x00,
}

func (m *Route) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Route) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Route) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Group) > 0 {
		i -= len(m.Group)
		copy(dAtA[i:], m.Group)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Group)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Packet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Packet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Packet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0x9a
	}
	if m.Route != nil {
		{
			size, err := m.Route.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5
		i--
		dAtA[i] = 0xda
	}
	if m.RouteMode != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.RouteMode))
		i--
		dAtA[i] = 0x5
		i--
		dAtA[i] = 0xd0
	}
	if len(m.LocalAddr) > 0 {
		i -= len(m.LocalAddr)
		copy(dAtA[i:], m.LocalAddr)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.LocalAddr)))
		i--
		dAtA[i] = 0x4
		i--
		dAtA[i] = 0xb2
	}
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0x3
		i--
		dAtA[i] = 0x9a
	}
	if m.AuthMode != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.AuthMode))
		i--
		dAtA[i] = 0x3
		i--
		dAtA[i] = 0x90
	}
	if len(m.AppName) > 0 {
		i -= len(m.AppName)
		copy(dAtA[i:], m.AppName)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.AppName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Udid) > 0 {
		i -= len(m.Udid)
		copy(dAtA[i:], m.Udid)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Udid)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Msg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Msg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Msg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0x9a
	}
	if len(m.RequestId) > 0 {
		i -= len(m.RequestId)
		copy(dAtA[i:], m.RequestId)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.RequestId)))
		i--
		dAtA[i] = 0x6
		i--
		dAtA[i] = 0x92
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Route) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Group)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPacket(uint64(m.Id))
	}
	return n
}

func (m *Packet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovPacket(uint64(m.Type))
	}
	l = len(m.Udid)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.AppName)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.AuthMode != 0 {
		n += 2 + sovPacket(uint64(m.AuthMode))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 2 + l + sovPacket(uint64(l))
	}
	l = len(m.LocalAddr)
	if l > 0 {
		n += 2 + l + sovPacket(uint64(l))
	}
	if m.RouteMode != 0 {
		n += 2 + sovPacket(uint64(m.RouteMode))
	}
	if m.Route != nil {
		l = m.Route.Size()
		n += 2 + l + sovPacket(uint64(l))
	}
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 2 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *Msg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RequestId)
	if l > 0 {
		n += 2 + l + sovPacket(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 2 + l + sovPacket(uint64(l))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Route) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Route: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Route: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Packet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Packet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Packet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Udid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Udid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 50:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthMode", wireType)
			}
			m.AuthMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthMode |= AuthMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 51:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 70:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocalAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 90:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteMode", wireType)
			}
			m.RouteMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RouteMode |= RouteMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 91:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Route", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Route == nil {
				m.Route = &Route{}
			}
			if err := m.Route.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 99:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &Msg{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Msg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Msg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Msg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 98:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 99:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
