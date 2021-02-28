// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_message.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Laptop struct {
	Id       string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu" json:"cpu,omitempty"`
	Ram      *Memory    `protobuf:"bytes,5,opt,name=ram" json:"ram,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus" json:"gpus,omitempty"`
	Storages []*Storage `protobuf:"bytes,7,rep,name=storages" json:"storages,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard" json:"keyboard,omitempty"`
	// Types that are valid to be assigned to Weight:
	//	*Laptop_WeightKg
	//	*Laptop_WeightLb
	Weight      isLaptop_Weight            `protobuf_oneof:"weight"`
	PriceUsd    float64                    `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd" json:"price_usd,omitempty"`
	ReleaseYear uint32                     `protobuf:"varint,13,opt,name=release_year,json=releaseYear" json:"release_year,omitempty"`
	UpdatedAt   *google_protobuf.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Laptop) Reset()                    { *m = Laptop{} }
func (m *Laptop) String() string            { return proto.CompactTextString(m) }
func (*Laptop) ProtoMessage()               {}
func (*Laptop) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isLaptop_Weight interface{ isLaptop_Weight() }

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,oneof"`
}
type Laptop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,oneof"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}
func (*Laptop_WeightLb) isLaptop_Weight() {}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Laptop) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Laptop) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Laptop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Laptop) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Laptop) GetRam() *Memory {
	if m != nil {
		return m.Ram
	}
	return nil
}

func (m *Laptop) GetGpus() []*GPU {
	if m != nil {
		return m.Gpus
	}
	return nil
}

func (m *Laptop) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func (m *Laptop) GetScreen() *Screen {
	if m != nil {
		return m.Screen
	}
	return nil
}

func (m *Laptop) GetKeyboard() *Keyboard {
	if m != nil {
		return m.Keyboard
	}
	return nil
}

func (m *Laptop) GetWeightKg() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (m *Laptop) GetWeightLb() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightLb); ok {
		return x.WeightLb
	}
	return 0
}

func (m *Laptop) GetPriceUsd() float64 {
	if m != nil {
		return m.PriceUsd
	}
	return 0
}

func (m *Laptop) GetReleaseYear() uint32 {
	if m != nil {
		return m.ReleaseYear
	}
	return 0
}

func (m *Laptop) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Laptop) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Laptop_OneofMarshaler, _Laptop_OneofUnmarshaler, _Laptop_OneofSizer, []interface{}{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLb)(nil),
	}
}

func _Laptop_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Laptop)
	// weight
	switch x := m.Weight.(type) {
	case *Laptop_WeightKg:
		b.EncodeVarint(10<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.WeightKg))
	case *Laptop_WeightLb:
		b.EncodeVarint(11<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.WeightLb))
	case nil:
	default:
		return fmt.Errorf("Laptop.Weight has unexpected type %T", x)
	}
	return nil
}

func _Laptop_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Laptop)
	switch tag {
	case 10: // weight.weight_kg
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Weight = &Laptop_WeightKg{math.Float64frombits(x)}
		return true, err
	case 11: // weight.weight_lb
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Weight = &Laptop_WeightLb{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Laptop_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Laptop)
	// weight
	switch x := m.Weight.(type) {
	case *Laptop_WeightKg:
		n += proto.SizeVarint(10<<3 | proto.WireFixed64)
		n += 8
	case *Laptop_WeightLb:
		n += proto.SizeVarint(11<<3 | proto.WireFixed64)
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Laptop)(nil), "com.pcbook.Laptop")
}

func init() { proto.RegisterFile("laptop_message.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0xdb, 0x85, 0xe4, 0x75, 0x1b, 0x92, 0x29, 0x60, 0x15, 0x21, 0x32, 0xe0, 0x50,
	0x71, 0x48, 0xa6, 0x71, 0xe2, 0xc8, 0x7a, 0x18, 0xd2, 0x36, 0x54, 0x79, 0xeb, 0x01, 0x2e, 0x91,
	0x93, 0x3c, 0xd2, 0xa8, 0x71, 0x6d, 0xd9, 0x8e, 0x50, 0x39, 0xf2, 0x97, 0xa3, 0x3a, 0xe9, 0xe8,
	0xc2, 0xcd, 0xfe, 0x7e, 0x3e, 0x4a, 0xde, 0x0f, 0xc3, 0xa4, 0xe6, 0xca, 0x4a, 0x95, 0x0a, 0x34,
	0x86, 0x97, 0x18, 0x2b, 0x2d, 0xad, 0x24, 0x90, 0x4b, 0x11, 0xab, 0x3c, 0x93, 0x72, 0x3d, 0x7d,
	0xa5, 0xb4, 0xcc, 0xd1, 0x18, 0xa9, 0x1f, 0x4b, 0xd3, 0x89, 0x40, 0x21, 0xf5, 0xb6, 0x97, 0xbe,
	0x30, 0x56, 0x6a, 0x5e, 0x62, 0x5f, 0x36, 0xb9, 0x46, 0xdc, 0xf4, 0xd2, 0x97, 0x6b, 0xdc, 0x66,
	0x92, 0xeb, 0xa2, 0x97, 0xbf, 0x2d, 0xa5, 0x2c, 0x6b, 0x4c, 0xdc, 0x2d, 0x6b, 0x7e, 0x26, 0xb6,
	0x12, 0x68, 0x2c, 0x17, 0xaa, 0x15, 0xde, 0xfd, 0x19, 0x81, 0x7f, 0xe3, 0x2a, 0x27, 0xa7, 0x30,
	0xa8, 0x0a, 0xea, 0x45, 0xde, 0x2c, 0x64, 0x83, 0xaa, 0x20, 0x13, 0x38, 0xca, 0x34, 0xdf, 0x14,
	0x74, 0xe0, 0xa2, 0xf6, 0x42, 0x08, 0x8c, 0x36, 0x5c, 0x20, 0x1d, 0xba, 0xd0, 0x9d, 0xc9, 0x19,
	0x0c, 0x73, 0xd5, 0xd0, 0x51, 0xe4, 0xcd, 0xc6, 0x17, 0xcf, 0xe2, 0x7f, 0x3d, 0xc7, 0xf3, 0xc5,
	0x92, 0xed, 0x18, 0xf9, 0x00, 0x43, 0xcd, 0x05, 0x3d, 0x72, 0x0a, 0x39, 0x54, 0x6e, 0x5d, 0xf3,
	0x6c, 0x87, 0xc9, 0x7b, 0x18, 0x95, 0xaa, 0x31, 0xd4, 0x8f, 0x86, 0xfd, 0x2f, 0x5d, 0x2d, 0x96,
	0xcc, 0x41, 0x92, 0x40, 0xd0, 0x8d, 0xc6, 0xd0, 0xa7, 0x4e, 0x7c, 0x7e, 0x28, 0xde, 0xb5, 0x8c,
	0x3d, 0x48, 0xe4, 0x23, 0xf8, 0xed, 0xd0, 0x68, 0xf0, 0xff, 0xef, 0xef, 0x1c, 0x61, 0x9d, 0x41,
	0xce, 0x21, 0xd8, 0x8f, 0x92, 0x86, 0xce, 0x9e, 0x1c, 0xda, 0xd7, 0x1d, 0x63, 0x0f, 0x16, 0x79,
	0x03, 0xe1, 0x2f, 0xac, 0xca, 0x95, 0x4d, 0xd7, 0x25, 0x85, 0xc8, 0x9b, 0x79, 0x5f, 0x9f, 0xb0,
	0xa0, 0x8d, 0xae, 0xcb, 0x03, 0x5c, 0x67, 0x74, 0xfc, 0x18, 0xdf, 0x64, 0xe4, 0x35, 0x84, 0x4a,
	0x57, 0x39, 0xa6, 0x8d, 0x29, 0xe8, 0xf1, 0x0e, 0xb3, 0xc0, 0x05, 0x4b, 0x53, 0x90, 0x33, 0x38,
	0xd6, 0x58, 0x23, 0x37, 0x98, 0x6e, 0x91, 0x6b, 0x7a, 0x12, 0x79, 0xb3, 0x13, 0x36, 0xee, 0xb2,
	0xef, 0xc8, 0x35, 0xf9, 0x0c, 0xd0, 0xa8, 0x82, 0x5b, 0x2c, 0x52, 0x6e, 0xe9, 0xa9, 0xab, 0x78,
	0x1a, 0xb7, 0x5b, 0x8f, 0xf7, 0x5b, 0x8f, 0xef, 0xf7, 0x5b, 0x67, 0x61, 0x67, 0x7f, 0xb1, 0x97,
	0x01, 0xf8, 0x6d, 0x19, 0x97, 0x17, 0x3f, 0xce, 0xcb, 0xca, 0xae, 0x9a, 0x6c, 0xd7, 0x6a, 0x72,
	0x2b, 0x57, 0x5c, 0x60, 0xf1, 0x8d, 0xff, 0xae, 0x74, 0x72, 0xbf, 0xc2, 0xb9, 0x14, 0xaa, 0x46,
	0x8b, 0x57, 0x6c, 0x31, 0x6f, 0x1f, 0x52, 0xa2, 0xb2, 0xcc, 0x77, 0xa7, 0x4f, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x4a, 0x01, 0x0b, 0xf0, 0xf8, 0x02, 0x00, 0x00,
}
