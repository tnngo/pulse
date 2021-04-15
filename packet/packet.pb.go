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
	// Unknown reserved.
	Type_Unknown Type = 0
	// Connect client request to connect to Server.
	Type_Connect Type = 1
	// ConnAck connect Acknowledgment.
	Type_ConnAck Type = 2
	// Ping client sends a heartbeat to the server.
	Type_Ping Type = 3
	// Pong server responds to the client heartbeat
	Type_Pong Type = 4
	// Body msg content.
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
	AuthMode_Default      AuthMode = 0
	AuthMode_CustomSecret AuthMode = 1
	AuthMode_FullyCustom  AuthMode = 2
)

var AuthMode_name = map[int32]string{
	0: "Default",
	1: "CustomSecret",
	2: "FullyCustom",
}

var AuthMode_value = map[string]int32{
	"Default":      0,
	"CustomSecret": 1,
	"FullyCustom":  2,
}

func (x AuthMode) String() string {
	return proto.EnumName(AuthMode_name, int32(x))
}

func (AuthMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{1}
}

// Packet packet message structure。
type Packet struct {
	// Type packet type。
	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=packet.Type" json:"type,omitempty"`
	// udid Unique identification of device terminal.
	Udid string `protobuf:"bytes,2,opt,name=udid,proto3" json:"udid,omitempty"`
	// just a simple security authorization field,
	// if you want to implement it yourself,
	// you can give up this field, and encapsulate it in the Body when the type is Connect.
	AuthMode  AuthMode `protobuf:"varint,3,opt,name=auth_mode,json=authMode,proto3,enum=packet.AuthMode" json:"auth_mode,omitempty"`
	Secret    string   `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	RouteId   int32    `protobuf:"varint,5,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	RequestId string   `protobuf:"bytes,6,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Body      []byte   `protobuf:"bytes,99,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{0}
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

func (m *Packet) GetAuthMode() AuthMode {
	if m != nil {
		return m.AuthMode
	}
	return AuthMode_Default
}

func (m *Packet) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Packet) GetRouteId() int32 {
	if m != nil {
		return m.RouteId
	}
	return 0
}

func (m *Packet) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *Packet) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterEnum("packet.Type", Type_name, Type_value)
	proto.RegisterEnum("packet.AuthMode", AuthMode_name, AuthMode_value)
	proto.RegisterType((*Packet)(nil), "packet.Packet")
}

func init() { proto.RegisterFile("packet.proto", fileDescriptor_e9ef1a6541f9f9e7) }

var fileDescriptor_e9ef1a6541f9f9e7 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x91, 0xc1, 0x6a, 0xf2, 0x40,
	0x10, 0xc7, 0xb3, 0x1a, 0x63, 0x1c, 0xc3, 0xf7, 0x2d, 0x7b, 0x28, 0xe9, 0xa1, 0x21, 0xf4, 0x14,
	0x84, 0x7a, 0x68, 0x6f, 0xbd, 0xa9, 0xa5, 0x20, 0xa5, 0x20, 0x69, 0x7b, 0xe9, 0x45, 0x62, 0x76,
	0xaa, 0xa2, 0xee, 0xa6, 0x71, 0x97, 0x92, 0xb7, 0xe8, 0x63, 0xf5, 0x28, 0xf4, 0xd2, 0x63, 0xd1,
	0x17, 0x29, 0xbb, 0xc6, 0xdb, 0xef, 0x37, 0xb3, 0xff, 0x65, 0x86, 0x81, 0xa0, 0xc8, 0xf2, 0x15,
	0xaa, 0x7e, 0x51, 0x4a, 0x25, 0x99, 0x77, 0xb4, 0xcb, 0x6f, 0x02, 0xde, 0xc4, 0x22, 0x8b, 0xc1,
	0x55, 0x55, 0x81, 0x21, 0x89, 0x49, 0xf2, 0xef, 0x3a, 0xe8, 0xd7, 0xef, 0x9f, 0xab, 0x02, 0x53,
	0xdb, 0x61, 0x0c, 0x5c, 0xcd, 0x97, 0x3c, 0x6c, 0xc4, 0x24, 0xe9, 0xa4, 0x96, 0xd9, 0x15, 0x74,
	0x32, 0xad, 0x16, 0xd3, 0x8d, 0xe4, 0x18, 0x36, 0x6d, 0x94, 0x9e, 0xa2, 0x03, 0xad, 0x16, 0x8f,
	0x92, 0x63, 0xea, 0x67, 0x35, 0xb1, 0x33, 0xf0, 0xb6, 0x98, 0x97, 0xa8, 0x42, 0xd7, 0x7e, 0x52,
	0x1b, 0x3b, 0x07, 0xbf, 0x94, 0x5a, 0xe1, 0x74, 0xc9, 0xc3, 0x56, 0x4c, 0x92, 0x56, 0xda, 0xb6,
	0x3e, 0xe6, 0xec, 0x02, 0xa0, 0xc4, 0x77, 0x8d, 0x5b, 0x65, 0x9a, 0x9e, 0x8d, 0x75, 0xea, 0xca,
	0x98, 0x9b, 0xa1, 0x66, 0x92, 0x57, 0x61, 0x1e, 0x93, 0x24, 0x48, 0x2d, 0xf7, 0x1e, 0xc0, 0x35,
	0x63, 0xb3, 0x2e, 0xb4, 0x5f, 0xc4, 0x4a, 0xc8, 0x0f, 0x41, 0x1d, 0x23, 0x23, 0x29, 0x04, 0xe6,
	0x8a, 0x92, 0x93, 0x0c, 0xf2, 0x15, 0x6d, 0x30, 0x1f, 0xdc, 0xc9, 0x52, 0xcc, 0x69, 0xd3, 0x92,
	0x14, 0x73, 0xea, 0x1a, 0x1a, 0x4a, 0x5e, 0xd1, 0xbc, 0x77, 0x0b, 0xfe, 0x69, 0x11, 0x13, 0xbb,
	0xc3, 0xb7, 0x4c, 0xaf, 0x15, 0x75, 0x18, 0x85, 0x60, 0xa4, 0xb7, 0x4a, 0x6e, 0x9e, 0xec, 0x0e,
	0x94, 0xb0, 0xff, 0xd0, 0xbd, 0xd7, 0xeb, 0x75, 0x75, 0x2c, 0xd3, 0xc6, 0x30, 0xfe, 0xda, 0x47,
	0x64, 0xb7, 0x8f, 0xc8, 0xef, 0x3e, 0x22, 0x9f, 0x87, 0xc8, 0xd9, 0x1d, 0x22, 0xe7, 0xe7, 0x10,
	0x39, 0xaf, 0xf5, 0x01, 0x66, 0x9e, 0xbd, 0xc7, 0xcd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56,
	0xcf, 0xc5, 0xe0, 0x9f, 0x01, 0x00, 0x00,
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
		dAtA[i] = 0x32
	}
	if m.RouteId != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.RouteId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0x22
	}
	if m.AuthMode != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.AuthMode))
		i--
		dAtA[i] = 0x18
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
	if m.AuthMode != 0 {
		n += 1 + sovPacket(uint64(m.AuthMode))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.RouteId != 0 {
		n += 1 + sovPacket(uint64(m.RouteId))
	}
	l = len(m.RequestId)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
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
		case 4:
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteId", wireType)
			}
			m.RouteId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RouteId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
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
