// Code generated by protoc-gen-gogo.
// source: validator.proto
// DO NOT EDIT!

/*
Package validator is a generated protocol buffer package.

It is generated from these files:
	validator.proto

It has these top-level messages:
	FieldValidator
*/
package validator

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FieldValidator struct {
	// Uses a Golang RE2-syntax regex to match the field contents.
	Regex *string `protobuf:"bytes,1,opt,name=regex" json:"regex,omitempty"`
	// Field value of integer strictly greater than this value.
	IntGt *int64 `protobuf:"varint,2,opt,name=int_gt,json=intGt" json:"int_gt,omitempty"`
	// Field value of integer strictly smaller than this value.
	IntLt *int64 `protobuf:"varint,3,opt,name=int_lt,json=intLt" json:"int_lt,omitempty"`
	// Used for nested message types, requires that the message type exists.
	MsgExists *bool `protobuf:"varint,4,opt,name=msg_exists,json=msgExists" json:"msg_exists,omitempty"`
	// Human error specifies a user-customizable error that is visible to the user.
	HumanError *string `protobuf:"bytes,5,opt,name=human_error,json=humanError" json:"human_error,omitempty"`
	// Field value of float strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatGt *float64 `protobuf:"fixed64,6,opt,name=float_gt,json=floatGt" json:"float_gt,omitempty"`
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatLt *float64 `protobuf:"fixed64,7,opt,name=float_lt,json=floatLt" json:"float_lt,omitempty"`
	// Field value of double describing the epsilon within which
	// any comparison should be considered to be true. For example,
	// when using float_gt = 0.35, using a float_epsilon of 0.05
	// would mean that any value above 0.30 is acceptable. It can be
	// thought of as a {float_value_condition} +- {float_epsilon}.
	// If unset, no correction for floating point inaccuracies in
	// comparisons will be attempted.
	FloatEpsilon     *float64 `protobuf:"fixed64,8,opt,name=float_epsilon,json=floatEpsilon" json:"float_epsilon,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FieldValidator) Reset()                    { *m = FieldValidator{} }
func (m *FieldValidator) String() string            { return proto.CompactTextString(m) }
func (*FieldValidator) ProtoMessage()               {}
func (*FieldValidator) Descriptor() ([]byte, []int) { return fileDescriptorValidator, []int{0} }

func (m *FieldValidator) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *FieldValidator) GetIntGt() int64 {
	if m != nil && m.IntGt != nil {
		return *m.IntGt
	}
	return 0
}

func (m *FieldValidator) GetIntLt() int64 {
	if m != nil && m.IntLt != nil {
		return *m.IntLt
	}
	return 0
}

func (m *FieldValidator) GetMsgExists() bool {
	if m != nil && m.MsgExists != nil {
		return *m.MsgExists
	}
	return false
}

func (m *FieldValidator) GetHumanError() string {
	if m != nil && m.HumanError != nil {
		return *m.HumanError
	}
	return ""
}

func (m *FieldValidator) GetFloatGt() float64 {
	if m != nil && m.FloatGt != nil {
		return *m.FloatGt
	}
	return 0
}

func (m *FieldValidator) GetFloatLt() float64 {
	if m != nil && m.FloatLt != nil {
		return *m.FloatLt
	}
	return 0
}

func (m *FieldValidator) GetFloatEpsilon() float64 {
	if m != nil && m.FloatEpsilon != nil {
		return *m.FloatEpsilon
	}
	return 0
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*FieldValidator)(nil),
	Field:         65020,
	Name:          "validator.field",
	Tag:           "bytes,65020,opt,name=field",
}

func init() {
	proto.RegisterType((*FieldValidator)(nil), "validator.FieldValidator")
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("validator.proto", fileDescriptorValidator) }

var fileDescriptorValidator = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0x89, 0xb3, 0xdb, 0xfa, 0xea, 0x0f, 0x08, 0x0a, 0x99, 0x30, 0x2c, 0x7a, 0xe9, 0xa9,
	0x03, 0x8f, 0x1e, 0x85, 0xba, 0xcb, 0x40, 0xe9, 0xc1, 0x83, 0x97, 0x52, 0x6d, 0x1a, 0x03, 0x69,
	0x53, 0x92, 0x37, 0xd9, 0x7f, 0xed, 0x1f, 0xa0, 0x07, 0x69, 0x42, 0x5b, 0x3d, 0xbe, 0xef, 0x6b,
	0xf3, 0x78, 0x1f, 0x9c, 0x7f, 0x96, 0x4a, 0x56, 0x25, 0x6a, 0x93, 0x76, 0x46, 0xa3, 0xa6, 0xe1,
	0x08, 0xae, 0x62, 0xa1, 0xb5, 0x50, 0x7c, 0xe3, 0xc4, 0xdb, 0xbe, 0xde, 0x54, 0xdc, 0xbe, 0x1b,
	0xd9, 0x8d, 0x1f, 0xdf, 0x7c, 0x11, 0x38, 0x7b, 0x94, 0x5c, 0x55, 0x2f, 0xc3, 0x4f, 0xf4, 0x02,
	0x02, 0xc3, 0x05, 0x3f, 0x30, 0x12, 0x93, 0x24, 0xcc, 0xfd, 0x40, 0x2f, 0x61, 0x2e, 0x5b, 0x2c,
	0x04, 0xb2, 0xa3, 0x98, 0x24, 0xb3, 0x3c, 0x90, 0x2d, 0x6e, 0x71, 0xc0, 0x0a, 0xd9, 0x6c, 0xc4,
	0x3b, 0xa4, 0x6b, 0x80, 0xc6, 0x8a, 0x82, 0x1f, 0xa4, 0x45, 0xcb, 0x8e, 0x63, 0x92, 0x2c, 0xf3,
	0xb0, 0xb1, 0x22, 0x73, 0x80, 0x5e, 0x43, 0xf4, 0xb1, 0x6f, 0xca, 0xb6, 0xe0, 0xc6, 0x68, 0xc3,
	0x02, 0xb7, 0x08, 0x1c, 0xca, 0x7a, 0x42, 0x57, 0xb0, 0xac, 0x95, 0x2e, 0xdd, 0xbe, 0x79, 0x4c,
	0x12, 0x92, 0x2f, 0xdc, 0xbc, 0xc5, 0x49, 0x29, 0x64, 0x8b, 0x3f, 0x6a, 0x87, 0xf4, 0x16, 0x4e,
	0xbd, 0xe2, 0x9d, 0x95, 0x4a, 0xb7, 0x6c, 0xe9, 0xfc, 0x89, 0x83, 0x99, 0x67, 0xf7, 0xcf, 0x10,
	0xd4, 0xfd, 0xc1, 0x74, 0x9d, 0xfa, 0x3a, 0xe9, 0x50, 0x27, 0x75, 0x21, 0x9e, 0x3a, 0x94, 0xba,
	0xb5, 0xec, 0xe7, 0xbb, 0xbf, 0x28, 0xba, 0x5b, 0xa5, 0x53, 0xe0, 0xff, 0xa5, 0x72, 0xff, 0xd0,
	0x43, 0xf4, 0x3a, 0x25, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xbb, 0xaf, 0xbf, 0x8f, 0x01,
	0x00, 0x00,
}
