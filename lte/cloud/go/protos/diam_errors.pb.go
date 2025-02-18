// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/diam_errors.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// ErrorCode reflects Experimental-Result values which are 3GPP failures
// to be processed by EPC. Diameter Base Protocol errors are reflected in gRPC status code
type ErrorCode int32

const (
	ErrorCode_UNDEFINED ErrorCode = 0
	// Default success code
	ErrorCode_MULTI_ROUND_AUTH        ErrorCode = 1001
	ErrorCode_SUCCESS                 ErrorCode = 2001
	ErrorCode_LIMITED_SUCCESS         ErrorCode = 2002
	ErrorCode_COMMAND_UNSUPORTED      ErrorCode = 3001
	ErrorCode_UNABLE_TO_DELIVER       ErrorCode = 3002
	ErrorCode_REALM_NOT_SERVED        ErrorCode = 3003
	ErrorCode_TOO_BUSY                ErrorCode = 3004
	ErrorCode_LOOP_DETECTED           ErrorCode = 3005
	ErrorCode_REDIRECT_INDICATION     ErrorCode = 3006
	ErrorCode_APPLICATION_UNSUPPORTED ErrorCode = 3007
	ErrorCode_INVALIDH_DR_BITS        ErrorCode = 3008
	ErrorCode_INVALID_AVP_BITS        ErrorCode = 3009
	ErrorCode_UNKNOWN_PEER            ErrorCode = 3010
	ErrorCode_AUTHENTICATION_REJECTED ErrorCode = 4001
	ErrorCode_OUT_OF_SPACE            ErrorCode = 4002
	ErrorCode_ELECTION_LOST           ErrorCode = 4003
	ErrorCode_AUTHORIZATION_REJECTED  ErrorCode = 5003
	// Permanent Failures 7.4.3
	ErrorCode_USER_UNKNOWN             ErrorCode = 5001
	ErrorCode_UNKNOWN_SESSION_ID       ErrorCode = 5002
	ErrorCode_UNKNOWN_EPS_SUBSCRIPTION ErrorCode = 5420
	ErrorCode_RAT_NOT_ALLOWED          ErrorCode = 5421
	ErrorCode_ROAMING_NOT_ALLOWED      ErrorCode = 5004
	ErrorCode_EQUIPMENT_UNKNOWN        ErrorCode = 5422
	ErrorCode_UNKNOWN_SERVING_NODE     ErrorCode = 5423
	// Transient Failures 7.4.4
	ErrorCode_AUTHENTICATION_DATA_UNAVAILABLE ErrorCode = 4181
)

var ErrorCode_name = map[int32]string{
	0:    "UNDEFINED",
	1001: "MULTI_ROUND_AUTH",
	2001: "SUCCESS",
	2002: "LIMITED_SUCCESS",
	3001: "COMMAND_UNSUPORTED",
	3002: "UNABLE_TO_DELIVER",
	3003: "REALM_NOT_SERVED",
	3004: "TOO_BUSY",
	3005: "LOOP_DETECTED",
	3006: "REDIRECT_INDICATION",
	3007: "APPLICATION_UNSUPPORTED",
	3008: "INVALIDH_DR_BITS",
	3009: "INVALID_AVP_BITS",
	3010: "UNKNOWN_PEER",
	4001: "AUTHENTICATION_REJECTED",
	4002: "OUT_OF_SPACE",
	4003: "ELECTION_LOST",
	5003: "AUTHORIZATION_REJECTED",
	5001: "USER_UNKNOWN",
	5002: "UNKNOWN_SESSION_ID",
	5420: "UNKNOWN_EPS_SUBSCRIPTION",
	5421: "RAT_NOT_ALLOWED",
	5004: "ROAMING_NOT_ALLOWED",
	5422: "EQUIPMENT_UNKNOWN",
	5423: "UNKNOWN_SERVING_NODE",
	4181: "AUTHENTICATION_DATA_UNAVAILABLE",
}

var ErrorCode_value = map[string]int32{
	"UNDEFINED":                       0,
	"MULTI_ROUND_AUTH":                1001,
	"SUCCESS":                         2001,
	"LIMITED_SUCCESS":                 2002,
	"COMMAND_UNSUPORTED":              3001,
	"UNABLE_TO_DELIVER":               3002,
	"REALM_NOT_SERVED":                3003,
	"TOO_BUSY":                        3004,
	"LOOP_DETECTED":                   3005,
	"REDIRECT_INDICATION":             3006,
	"APPLICATION_UNSUPPORTED":         3007,
	"INVALIDH_DR_BITS":                3008,
	"INVALID_AVP_BITS":                3009,
	"UNKNOWN_PEER":                    3010,
	"AUTHENTICATION_REJECTED":         4001,
	"OUT_OF_SPACE":                    4002,
	"ELECTION_LOST":                   4003,
	"AUTHORIZATION_REJECTED":          5003,
	"USER_UNKNOWN":                    5001,
	"UNKNOWN_SESSION_ID":              5002,
	"UNKNOWN_EPS_SUBSCRIPTION":        5420,
	"RAT_NOT_ALLOWED":                 5421,
	"ROAMING_NOT_ALLOWED":             5004,
	"EQUIPMENT_UNKNOWN":               5422,
	"UNKNOWN_SERVING_NODE":            5423,
	"AUTHENTICATION_DATA_UNAVAILABLE": 4181,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98654a741ff5fd83, []int{0}
}

func init() {
	proto.RegisterEnum("magma.lte.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() { proto.RegisterFile("lte/protos/diam_errors.proto", fileDescriptor_98654a741ff5fd83) }

var fileDescriptor_98654a741ff5fd83 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcb, 0x6e, 0x53, 0x31,
	0x10, 0x86, 0x59, 0x20, 0x4a, 0xac, 0x46, 0x71, 0x4d, 0xe9, 0x69, 0xd5, 0xa2, 0xb0, 0x40, 0x8a,
	0x94, 0x45, 0xb3, 0xe0, 0x09, 0x9c, 0xe3, 0x29, 0x31, 0xf8, 0xd8, 0xc6, 0x97, 0x53, 0xd1, 0xcd,
	0x28, 0x90, 0xa8, 0x42, 0x4a, 0x74, 0x50, 0x1a, 0x1e, 0x82, 0xcb, 0x4b, 0x00, 0x5b, 0x2e, 0x6b,
	0xee, 0x97, 0x37, 0x00, 0x89, 0x87, 0x60, 0xc5, 0x2b, 0x20, 0x9f, 0x24, 0x40, 0xbb, 0xb2, 0xf4,
	0x8f, 0x67, 0xbe, 0x5f, 0xff, 0x0c, 0xd9, 0x9b, 0xcc, 0xc7, 0xbd, 0x07, 0xb3, 0x6a, 0x5e, 0x9d,
	0xf4, 0x46, 0xf7, 0x87, 0x53, 0x1c, 0xcf, 0x66, 0xd5, 0xec, 0x64, 0xbf, 0x96, 0x58, 0x63, 0x3a,
	0x3c, 0x9e, 0x0e, 0xf7, 0x27, 0xf3, 0x71, 0xf7, 0xf7, 0x79, 0xd2, 0x80, 0x54, 0xcb, 0xab, 0xd1,
	0x98, 0x35, 0x49, 0x23, 0x6a, 0x01, 0x07, 0x52, 0x83, 0xa0, 0xe7, 0xd8, 0x65, 0x42, 0x8b, 0xa8,
	0x82, 0x44, 0x67, 0xa2, 0x16, 0xc8, 0x63, 0x18, 0xd0, 0x5f, 0x6b, 0x6c, 0x9d, 0xac, 0xf9, 0x98,
	0xe7, 0xe0, 0x3d, 0xfd, 0xde, 0x62, 0x9b, 0xa4, 0xa5, 0x64, 0x21, 0x03, 0x08, 0x5c, 0xa9, 0x3f,
	0x5a, 0x2c, 0x23, 0x2c, 0x37, 0x45, 0xc1, 0xb5, 0xc0, 0xa8, 0x7d, 0xb4, 0xc6, 0x05, 0x10, 0xf4,
	0x6d, 0xc6, 0xb6, 0xc8, 0x46, 0xd4, 0xbc, 0xaf, 0x00, 0x83, 0x41, 0x01, 0x4a, 0x96, 0xe0, 0xe8,
	0xbb, 0x2c, 0xb1, 0x1c, 0x70, 0x55, 0xa0, 0x36, 0x01, 0x3d, 0xb8, 0x12, 0x04, 0x7d, 0x9f, 0xb1,
	0x26, 0xb9, 0x18, 0x8c, 0xc1, 0x7e, 0xf4, 0x77, 0xe8, 0x87, 0x8c, 0x31, 0xd2, 0x54, 0xc6, 0x58,
	0x14, 0x10, 0x20, 0x4f, 0x13, 0x3f, 0x66, 0x6c, 0x9b, 0x5c, 0x72, 0x20, 0xa4, 0x83, 0x3c, 0xa0,
	0xd4, 0x42, 0xe6, 0x3c, 0x48, 0xa3, 0xe9, 0xa7, 0x8c, 0xed, 0x91, 0x8c, 0x5b, 0xab, 0x96, 0xca,
	0xc2, 0xc8, 0xd2, 0xc9, 0xe7, 0x9a, 0x28, 0x75, 0xc9, 0x95, 0x14, 0x03, 0x14, 0x0e, 0xfb, 0x32,
	0x78, 0xfa, 0xe5, 0x7f, 0x19, 0x79, 0x69, 0x17, 0xf2, 0xd7, 0x8c, 0x6d, 0x90, 0xf5, 0xa8, 0x6f,
	0x69, 0x73, 0xa8, 0xd1, 0x02, 0x38, 0xfa, 0x6d, 0x31, 0x3e, 0x86, 0x01, 0xe8, 0xb0, 0x22, 0x38,
	0xb8, 0xb9, 0xb0, 0xf5, 0xac, 0x9d, 0x1a, 0x4c, 0x0c, 0x68, 0x0e, 0xd0, 0x5b, 0x9e, 0x03, 0x7d,
	0xde, 0x4e, 0xee, 0x41, 0x41, 0x5e, 0x7f, 0x55, 0xc6, 0x07, 0xfa, 0xa2, 0xcd, 0x76, 0xc9, 0x56,
	0x1a, 0x62, 0x9c, 0x3c, 0x3a, 0x33, 0xe3, 0x49, 0xa7, 0x86, 0x7a, 0x70, 0xb8, 0x24, 0xd3, 0x47,
	0x9d, 0x14, 0xec, 0xca, 0x87, 0x07, 0xef, 0x53, 0x87, 0x14, 0xf4, 0x71, 0x87, 0x5d, 0x21, 0xdb,
	0xab, 0x02, 0x58, 0x8f, 0x3e, 0xf6, 0x7d, 0xee, 0xa4, 0xad, 0xb3, 0x78, 0xd9, 0x4d, 0x6b, 0x72,
	0x3c, 0xd4, 0xe9, 0x72, 0xa5, 0xcc, 0x21, 0x08, 0xfa, 0xaa, 0x5b, 0x67, 0x67, 0x78, 0x21, 0xf5,
	0x8d, 0x53, 0x95, 0xa7, 0x9d, 0xb4, 0x27, 0xb8, 0x1d, 0xa5, 0x2d, 0x40, 0x87, 0xbf, 0xfc, 0xd7,
	0x5d, 0xb6, 0x43, 0x36, 0xff, 0xf1, 0x5d, 0xb9, 0xe8, 0x14, 0x40, 0xdf, 0x74, 0xd9, 0x35, 0xd2,
	0x3e, 0x93, 0x87, 0xe0, 0x81, 0x63, 0xd4, 0xbc, 0xe4, 0x52, 0xa5, 0x9d, 0xd3, 0x9f, 0x57, 0xfb,
	0xbb, 0x47, 0x3b, 0xf5, 0xf9, 0xf5, 0xd2, 0x89, 0xde, 0x9b, 0x54, 0x0f, 0x47, 0xbd, 0xe3, 0x6a,
	0x79, 0xab, 0x77, 0x2f, 0xd4, 0xef, 0xf5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x0a, 0x98,
	0xe7, 0xc0, 0x02, 0x00, 0x00,
}
