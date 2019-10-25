// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enumdecl.proto

package enumdecl

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gracenoah/protobuf/gogoproto"
	proto "github.com/gracenoah/protobuf/proto"
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

var MyEnum_name = map[int32]string{
	0: "A",
	1: "B",
}

var MyEnum_value = map[string]int32{
	"A": 0,
	"B": 1,
}

func (x MyEnum) String() string {
	return proto.EnumName(MyEnum_name, int32(x))
}

func (MyEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_869e4edba7b1e4e3, []int{0}
}

type Message struct {
	EnumeratedField      MyEnum   `protobuf:"varint,1,opt,name=enumerated_field,json=enumeratedField,proto3,enum=enumdecl.MyEnum" json:"enumerated_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_869e4edba7b1e4e3, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetEnumeratedField() MyEnum {
	if m != nil {
		return m.EnumeratedField
	}
	return A
}

func init() {
	proto.RegisterEnum("enumdecl.MyEnum", MyEnum_name, MyEnum_value)
	proto.RegisterType((*Message)(nil), "enumdecl.Message")
}

func init() { proto.RegisterFile("enumdecl.proto", fileDescriptor_869e4edba7b1e4e3) }

var fileDescriptor_869e4edba7b1e4e3 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcd, 0x2b, 0xcd,
	0x4d, 0x49, 0x4d, 0xce, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x74,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5,
	0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x54, 0x72, 0xe3, 0x62,
	0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0xb2, 0xe6, 0x12, 0x00, 0x99, 0x92, 0x5a, 0x94,
	0x58, 0x92, 0x9a, 0x12, 0x9f, 0x96, 0x99, 0x9a, 0x93, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x67,
	0x24, 0xa0, 0x07, 0xb7, 0xce, 0xb7, 0xd2, 0x35, 0xaf, 0x34, 0x37, 0x88, 0x1f, 0xa1, 0xd2, 0x0d,
	0xa4, 0x50, 0x4b, 0x81, 0x8b, 0x0d, 0x22, 0x25, 0xc4, 0xca, 0xc5, 0xe8, 0x28, 0xc0, 0x00, 0xa2,
	0x9c, 0x04, 0x18, 0xa5, 0x38, 0x3a, 0x16, 0xcb, 0x31, 0x1c, 0x58, 0x22, 0xc7, 0xe0, 0xa4, 0xf1,
	0xe0, 0xa1, 0x1c, 0xe3, 0x8f, 0x87, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0xee, 0x78, 0x24, 0xc7,
	0x78, 0xe0, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0xf8, 0xe3, 0x91, 0x1c, 0x43, 0xc3, 0x63, 0x39, 0x86, 0x24, 0x36, 0xb0, 0xd3, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x76, 0x04, 0x55, 0xb7, 0xe5, 0x00, 0x00, 0x00,
}

func (this *Message) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Message)
	if !ok {
		that2, ok := that.(Message)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Message")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Message but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Message but is not nil && this == nil")
	}
	if this.EnumeratedField != that1.EnumeratedField {
		return fmt.Errorf("EnumeratedField this(%v) Not Equal that(%v)", this.EnumeratedField, that1.EnumeratedField)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Message) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Message)
	if !ok {
		that2, ok := that.(Message)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.EnumeratedField != that1.EnumeratedField {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.EnumeratedField != 0 {
		i = encodeVarintEnumdecl(dAtA, i, uint64(m.EnumeratedField))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEnumdecl(dAtA []byte, offset int, v uint64) int {
	offset -= sovEnumdecl(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedMessage(r randyEnumdecl, easy bool) *Message {
	this := &Message{}
	this.EnumeratedField = MyEnum([]int32{0, 1}[r.Intn(2)])
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEnumdecl(r, 2)
	}
	return this
}

type randyEnumdecl interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEnumdecl(r randyEnumdecl) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringEnumdecl(r randyEnumdecl) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneEnumdecl(r)
	}
	return string(tmps)
}
func randUnrecognizedEnumdecl(r randyEnumdecl, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldEnumdecl(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldEnumdecl(dAtA []byte, r randyEnumdecl, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateEnumdecl(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateEnumdecl(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateEnumdecl(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateEnumdecl(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateEnumdecl(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateEnumdecl(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateEnumdecl(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnumeratedField != 0 {
		n += 1 + sovEnumdecl(uint64(m.EnumeratedField))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEnumdecl(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEnumdecl(x uint64) (n int) {
	return sovEnumdecl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnumdecl
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnumeratedField", wireType)
			}
			m.EnumeratedField = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnumdecl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EnumeratedField |= MyEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEnumdecl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnumdecl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEnumdecl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEnumdecl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnumdecl
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
					return 0, ErrIntOverflowEnumdecl
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
					return 0, ErrIntOverflowEnumdecl
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
				return 0, ErrInvalidLengthEnumdecl
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEnumdecl
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEnumdecl
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEnumdecl        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnumdecl          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEnumdecl = fmt.Errorf("proto: unexpected end of group")
)
