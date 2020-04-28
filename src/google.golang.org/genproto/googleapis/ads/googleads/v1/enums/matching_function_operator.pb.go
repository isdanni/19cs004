// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/matching_function_operator.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible operators in a matching function.
type MatchingFunctionOperatorEnum_MatchingFunctionOperator int32

const (
	// Not specified.
	MatchingFunctionOperatorEnum_UNSPECIFIED MatchingFunctionOperatorEnum_MatchingFunctionOperator = 0
	// Used for return value only. Represents value unknown in this version.
	MatchingFunctionOperatorEnum_UNKNOWN MatchingFunctionOperatorEnum_MatchingFunctionOperator = 1
	// The IN operator.
	MatchingFunctionOperatorEnum_IN MatchingFunctionOperatorEnum_MatchingFunctionOperator = 2
	// The IDENTITY operator.
	MatchingFunctionOperatorEnum_IDENTITY MatchingFunctionOperatorEnum_MatchingFunctionOperator = 3
	// The EQUALS operator
	MatchingFunctionOperatorEnum_EQUALS MatchingFunctionOperatorEnum_MatchingFunctionOperator = 4
	// Operator that takes two or more operands that are of type
	// FunctionOperand and checks that all the operands evaluate to true.
	// For functions related to ad formats, all the operands must be in
	// left_operands.
	MatchingFunctionOperatorEnum_AND MatchingFunctionOperatorEnum_MatchingFunctionOperator = 5
	// Operator that returns true if the elements in left_operands contain any
	// of the elements in right_operands. Otherwise, return false. The
	// right_operands must contain at least 1 and no more than 3
	// ConstantOperands.
	MatchingFunctionOperatorEnum_CONTAINS_ANY MatchingFunctionOperatorEnum_MatchingFunctionOperator = 6
)

var MatchingFunctionOperatorEnum_MatchingFunctionOperator_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "IN",
	3: "IDENTITY",
	4: "EQUALS",
	5: "AND",
	6: "CONTAINS_ANY",
}
var MatchingFunctionOperatorEnum_MatchingFunctionOperator_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"IN":           2,
	"IDENTITY":     3,
	"EQUALS":       4,
	"AND":          5,
	"CONTAINS_ANY": 6,
}

func (x MatchingFunctionOperatorEnum_MatchingFunctionOperator) String() string {
	return proto.EnumName(MatchingFunctionOperatorEnum_MatchingFunctionOperator_name, int32(x))
}
func (MatchingFunctionOperatorEnum_MatchingFunctionOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_matching_function_operator_bca31fb04efa2051, []int{0, 0}
}

// Container for enum describing matching function operator.
type MatchingFunctionOperatorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchingFunctionOperatorEnum) Reset()         { *m = MatchingFunctionOperatorEnum{} }
func (m *MatchingFunctionOperatorEnum) String() string { return proto.CompactTextString(m) }
func (*MatchingFunctionOperatorEnum) ProtoMessage()    {}
func (*MatchingFunctionOperatorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_matching_function_operator_bca31fb04efa2051, []int{0}
}
func (m *MatchingFunctionOperatorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchingFunctionOperatorEnum.Unmarshal(m, b)
}
func (m *MatchingFunctionOperatorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchingFunctionOperatorEnum.Marshal(b, m, deterministic)
}
func (dst *MatchingFunctionOperatorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchingFunctionOperatorEnum.Merge(dst, src)
}
func (m *MatchingFunctionOperatorEnum) XXX_Size() int {
	return xxx_messageInfo_MatchingFunctionOperatorEnum.Size(m)
}
func (m *MatchingFunctionOperatorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchingFunctionOperatorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MatchingFunctionOperatorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MatchingFunctionOperatorEnum)(nil), "google.ads.googleads.v1.enums.MatchingFunctionOperatorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.MatchingFunctionOperatorEnum_MatchingFunctionOperator", MatchingFunctionOperatorEnum_MatchingFunctionOperator_name, MatchingFunctionOperatorEnum_MatchingFunctionOperator_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/matching_function_operator.proto", fileDescriptor_matching_function_operator_bca31fb04efa2051)
}

var fileDescriptor_matching_function_operator_bca31fb04efa2051 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4d, 0x6a, 0xf3, 0x30,
	0x14, 0xfc, 0xec, 0x7c, 0x75, 0x8a, 0x12, 0xa8, 0xd0, 0xaa, 0x94, 0x64, 0x91, 0x1c, 0x40, 0xc6,
	0x74, 0xa7, 0x42, 0x41, 0x49, 0x9c, 0x60, 0xda, 0x2a, 0x29, 0xf9, 0x29, 0x29, 0x86, 0xa0, 0xc6,
	0xae, 0x6a, 0x88, 0x25, 0x63, 0xd9, 0xb9, 0x49, 0x2f, 0xd0, 0x65, 0x8f, 0xd2, 0xa3, 0x74, 0xd9,
	0x13, 0x14, 0xff, 0x24, 0xbb, 0x74, 0x23, 0x06, 0xcd, 0x9b, 0x79, 0x6f, 0x06, 0xdc, 0x0a, 0xa5,
	0xc4, 0x2e, 0xb4, 0x79, 0xa0, 0xed, 0x0a, 0x16, 0x68, 0xef, 0xd8, 0xa1, 0xcc, 0x63, 0x6d, 0xc7,
	0x3c, 0xdb, 0xbe, 0x45, 0x52, 0x6c, 0x5e, 0x73, 0xb9, 0xcd, 0x22, 0x25, 0x37, 0x2a, 0x09, 0x53,
	0x9e, 0xa9, 0x14, 0x27, 0xa9, 0xca, 0x14, 0xea, 0x56, 0x22, 0xcc, 0x03, 0x8d, 0x8f, 0x7a, 0xbc,
	0x77, 0x70, 0xa9, 0xbf, 0xea, 0x1c, 0xec, 0x93, 0xc8, 0xe6, 0x52, 0xaa, 0x8c, 0x17, 0x26, 0xba,
	0x12, 0xf7, 0xdf, 0x0d, 0xd0, 0x79, 0xa8, 0x37, 0x8c, 0xeb, 0x05, 0xd3, 0xda, 0xdf, 0x95, 0x79,
	0xdc, 0xcf, 0xc1, 0xe5, 0x29, 0x1e, 0x5d, 0x80, 0xd6, 0x92, 0xcd, 0x67, 0xee, 0xd0, 0x1b, 0x7b,
	0xee, 0x08, 0xfe, 0x43, 0x2d, 0xd0, 0x5c, 0xb2, 0x3b, 0x36, 0x7d, 0x62, 0xd0, 0x40, 0x16, 0x30,
	0x3d, 0x06, 0x4d, 0xd4, 0x06, 0xe7, 0xde, 0xc8, 0x65, 0x0b, 0x6f, 0xb1, 0x86, 0x0d, 0x04, 0x80,
	0xe5, 0x3e, 0x2e, 0xe9, 0xfd, 0x1c, 0xfe, 0x47, 0x4d, 0xd0, 0xa0, 0x6c, 0x04, 0xcf, 0x10, 0x04,
	0xed, 0xe1, 0x94, 0x2d, 0xa8, 0xc7, 0xe6, 0x1b, 0xca, 0xd6, 0xd0, 0x1a, 0xfc, 0x18, 0xa0, 0xb7,
	0x55, 0x31, 0xfe, 0x33, 0xdb, 0xa0, 0x7b, 0xea, 0xb4, 0x59, 0x11, 0x6e, 0x66, 0x3c, 0x0f, 0x6a,
	0xbd, 0x50, 0x3b, 0x2e, 0x05, 0x56, 0xa9, 0xb0, 0x45, 0x28, 0xcb, 0xe8, 0x87, 0xae, 0x93, 0x48,
	0x9f, 0xa8, 0xfe, 0xa6, 0x7c, 0x3f, 0xcc, 0xc6, 0x84, 0xd2, 0x4f, 0xb3, 0x3b, 0xa9, 0xac, 0x68,
	0xa0, 0x71, 0x05, 0x0b, 0xb4, 0x72, 0x70, 0x51, 0x93, 0xfe, 0x3a, 0xf0, 0x3e, 0x0d, 0xb4, 0x7f,
	0xe4, 0xfd, 0x95, 0xe3, 0x97, 0xfc, 0xb7, 0xd9, 0xab, 0x3e, 0x09, 0xa1, 0x81, 0x26, 0xe4, 0x38,
	0x41, 0xc8, 0xca, 0x21, 0xa4, 0x9c, 0x79, 0xb1, 0xca, 0xc3, 0xae, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x5b, 0x67, 0x87, 0x8d, 0x12, 0x02, 0x00, 0x00,
}
