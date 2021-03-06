// Code generated by protoc-gen-go. DO NOT EDIT.
// source: processor_message.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CPU struct {
	Brand         string  `protobuf:"bytes,1,opt,name=brand" json:"brand,omitempty"`
	Name          string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	NumberCores   uint32  `protobuf:"varint,3,opt,name=number_cores,json=numberCores" json:"number_cores,omitempty"`
	NumberThreads uint32  `protobuf:"varint,4,opt,name=number_threads,json=numberThreads" json:"number_threads,omitempty"`
	MinGhz        float64 `protobuf:"fixed64,5,opt,name=min_ghz,json=minGhz" json:"min_ghz,omitempty"`
	MaxGhz        float64 `protobuf:"fixed64,6,opt,name=max_ghz,json=maxGhz" json:"max_ghz,omitempty"`
}

func (m *CPU) Reset()                    { *m = CPU{} }
func (m *CPU) String() string            { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()               {}
func (*CPU) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *CPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *CPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CPU) GetNumberCores() uint32 {
	if m != nil {
		return m.NumberCores
	}
	return 0
}

func (m *CPU) GetNumberThreads() uint32 {
	if m != nil {
		return m.NumberThreads
	}
	return 0
}

func (m *CPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *CPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

type GPU struct {
	Brand  string  `protobuf:"bytes,1,opt,name=brand" json:"brand,omitempty"`
	Name   string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	MinGhz float64 `protobuf:"fixed64,3,opt,name=min_ghz,json=minGhz" json:"min_ghz,omitempty"`
	MaxGhz float64 `protobuf:"fixed64,4,opt,name=max_ghz,json=maxGhz" json:"max_ghz,omitempty"`
	Memory *Memory `protobuf:"bytes,5,opt,name=memory" json:"memory,omitempty"`
}

func (m *GPU) Reset()                    { *m = GPU{} }
func (m *GPU) String() string            { return proto.CompactTextString(m) }
func (*GPU) ProtoMessage()               {}
func (*GPU) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *GPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *GPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *GPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

func (m *GPU) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterType((*CPU)(nil), "com.pcbook.CPU")
	proto.RegisterType((*GPU)(nil), "com.pcbook.GPU")
}

func init() { proto.RegisterFile("processor_message.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x99, 0x3f, 0x69, 0x7e, 0x9c, 0x5a, 0x17, 0x43, 0xa1, 0xc1, 0x55, 0x2c, 0x08, 0xc1,
	0x45, 0x22, 0xf5, 0x0d, 0xcc, 0x22, 0xab, 0x4a, 0x09, 0x75, 0xe3, 0x26, 0xcc, 0x24, 0x97, 0x4c,
	0xd0, 0xc9, 0x0d, 0x33, 0x29, 0xd4, 0xbc, 0x83, 0x4f, 0xe2, 0x4b, 0x4a, 0x67, 0x02, 0xd6, 0x85,
	0x0b, 0x77, 0x33, 0xdf, 0xf9, 0xe0, 0x5c, 0x0e, 0x5d, 0xf5, 0x1a, 0x2b, 0x30, 0x06, 0x75, 0xa9,
	0xc0, 0x18, 0xde, 0x40, 0xd2, 0x6b, 0x1c, 0x90, 0xd1, 0x0a, 0x55, 0xd2, 0x57, 0x02, 0xf1, 0xf5,
	0x7a, 0xa9, 0x40, 0xa1, 0x7e, 0xff, 0x69, 0xac, 0x3f, 0x09, 0xf5, 0xb2, 0xdd, 0x33, 0x5b, 0xd2,
	0x99, 0xd0, 0xbc, 0xab, 0x43, 0x12, 0x91, 0xf8, 0xa2, 0x70, 0x1f, 0xc6, 0xa8, 0xdf, 0x71, 0x05,
	0xe1, 0x3f, 0x0b, 0xed, 0x9b, 0xdd, 0xd0, 0xcb, 0xee, 0xa0, 0x04, 0xe8, 0xb2, 0x42, 0x0d, 0x26,
	0xf4, 0x22, 0x12, 0x2f, 0x8a, 0xb9, 0x63, 0xd9, 0x09, 0xb1, 0x5b, 0x7a, 0x35, 0x29, 0x83, 0xd4,
	0xc0, 0x6b, 0x13, 0xfa, 0x56, 0x5a, 0x38, 0xba, 0x77, 0x90, 0xad, 0xe8, 0x7f, 0xd5, 0x76, 0x65,
	0x23, 0xc7, 0x70, 0x16, 0x91, 0x98, 0x14, 0x81, 0x6a, 0xbb, 0x5c, 0x8e, 0x36, 0xe0, 0x47, 0x1b,
	0x04, 0x53, 0xc0, 0x8f, 0xb9, 0x1c, 0xd7, 0x1f, 0x84, 0x7a, 0xf9, 0x9f, 0xae, 0x3d, 0xeb, 0xf0,
	0x7e, 0xeb, 0xf0, 0xcf, 0x3b, 0xd8, 0x1d, 0x0d, 0xdc, 0x52, 0xf6, 0xa8, 0xf9, 0x86, 0x25, 0xdf,
	0x23, 0x26, 0x5b, 0x9b, 0x14, 0x93, 0xf1, 0xb8, 0x79, 0xb9, 0x6f, 0xda, 0x41, 0x1e, 0xc4, 0xc9,
	0x49, 0xb7, 0x28, 0xb9, 0x82, 0xfa, 0x89, 0x8f, 0xad, 0x4e, 0xf7, 0x12, 0x32, 0x54, 0xfd, 0x1b,
	0x0c, 0x90, 0x17, 0xbb, 0x2c, 0xb5, 0x73, 0xa7, 0xbd, 0x10, 0x81, 0x7d, 0x3d, 0x7c, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xc0, 0x78, 0x1f, 0x45, 0xb5, 0x01, 0x00, 0x00,
}
