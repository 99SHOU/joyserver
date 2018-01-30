// Code generated by protoc-gen-go. DO NOT EDIT.
// source: define.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	define.proto
	game_msg.proto
	framework_msg.proto

It has these top-level messages:
	NodeRegisterReq
	NodeRegisterAck
	SetNodeStatu
	NodeInfo
	GameNodeListReq
	GameNodeListAck
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MsgID int32

const (
	MsgID_FWM_NODE_REGISTER_REQ  MsgID = 0
	MsgID_FWM_NODE_REGISTER_ACK  MsgID = 1
	MsgID_FWM_SET_NODE_STATU     MsgID = 2
	MsgID_FWM_GAME_NODE_LIST_REQ MsgID = 3
	MsgID_FWM_GAME_NODE_LIST_ACK MsgID = 4
)

var MsgID_name = map[int32]string{
	0: "FWM_NODE_REGISTER_REQ",
	1: "FWM_NODE_REGISTER_ACK",
	2: "FWM_SET_NODE_STATU",
	3: "FWM_GAME_NODE_LIST_REQ",
	4: "FWM_GAME_NODE_LIST_ACK",
}
var MsgID_value = map[string]int32{
	"FWM_NODE_REGISTER_REQ":  0,
	"FWM_NODE_REGISTER_ACK":  1,
	"FWM_SET_NODE_STATU":     2,
	"FWM_GAME_NODE_LIST_REQ": 3,
	"FWM_GAME_NODE_LIST_ACK": 4,
}

func (x MsgID) String() string {
	return proto.EnumName(MsgID_name, int32(x))
}
func (MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NODE_TYPE int32

const (
	NODE_TYPE_INVALID NODE_TYPE = 0
	NODE_TYPE_MACHINE NODE_TYPE = 1
	NODE_TYPE_CENTER  NODE_TYPE = 2
	NODE_TYPE_GAME    NODE_TYPE = 3
	NODE_TYPE_GATE    NODE_TYPE = 4
	NODE_TYPE_LOGIN   NODE_TYPE = 5
	NODE_TYPE_CLIENT  NODE_TYPE = 6
)

var NODE_TYPE_name = map[int32]string{
	0: "INVALID",
	1: "MACHINE",
	2: "CENTER",
	3: "GAME",
	4: "GATE",
	5: "LOGIN",
	6: "CLIENT",
}
var NODE_TYPE_value = map[string]int32{
	"INVALID": 0,
	"MACHINE": 1,
	"CENTER":  2,
	"GAME":    3,
	"GATE":    4,
	"LOGIN":   5,
	"CLIENT":  6,
}

func (x NODE_TYPE) String() string {
	return proto.EnumName(NODE_TYPE_name, int32(x))
}
func (NODE_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NODE_STATU int32

const (
	NODE_STATU_NOT_READY NODE_STATU = 0
	NODE_STATU_READY     NODE_STATU = 1
	NODE_STATU_OFF_LINE  NODE_STATU = 2
)

var NODE_STATU_name = map[int32]string{
	0: "NOT_READY",
	1: "READY",
	2: "OFF_LINE",
}
var NODE_STATU_value = map[string]int32{
	"NOT_READY": 0,
	"READY":     1,
	"OFF_LINE":  2,
}

func (x NODE_STATU) String() string {
	return proto.EnumName(NODE_STATU_name, int32(x))
}
func (NODE_STATU) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type NODE_INFO int32

const (
	NODE_INFO_SERVER_PORT NODE_INFO = 0
)

var NODE_INFO_name = map[int32]string{
	0: "SERVER_PORT",
}
var NODE_INFO_value = map[string]int32{
	"SERVER_PORT": 0,
}

func (x NODE_INFO) String() string {
	return proto.EnumName(NODE_INFO_name, int32(x))
}
func (NODE_INFO) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterEnum("pb.MsgID", MsgID_name, MsgID_value)
	proto.RegisterEnum("pb.NODE_TYPE", NODE_TYPE_name, NODE_TYPE_value)
	proto.RegisterEnum("pb.NODE_STATU", NODE_STATU_name, NODE_STATU_value)
	proto.RegisterEnum("pb.NODE_INFO", NODE_INFO_name, NODE_INFO_value)
}

func init() { proto.RegisterFile("define.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xdb, 0xae, 0xad, 0xdb, 0xdd, 0xc4, 0xcb, 0x05, 0x07, 0x8a, 0xbf, 0xa0, 0x0f, 0xbe,
	0xe8, 0x1f, 0x08, 0xed, 0x6d, 0x0d, 0xb6, 0xe9, 0x4c, 0xae, 0x93, 0x81, 0x10, 0x18, 0x4e, 0xf1,
	0xc5, 0x0d, 0xf5, 0x6f, 0xf8, 0x9f, 0x25, 0xad, 0xe0, 0x8b, 0xbe, 0x9d, 0x9c, 0x2f, 0x39, 0xe7,
	0x10, 0x58, 0x3c, 0xed, 0x9e, 0x5f, 0xdf, 0x76, 0x97, 0x87, 0xf7, 0xfd, 0xe7, 0x9e, 0x92, 0xc3,
	0xb6, 0xf8, 0x8a, 0x21, 0xeb, 0x3e, 0x5e, 0x74, 0x45, 0x67, 0x70, 0x5a, 0x3f, 0x74, 0xde, 0xf4,
	0x15, 0x7b, 0xcb, 0x8d, 0x76, 0xc2, 0xd6, 0x5b, 0xbe, 0xc3, 0xe8, 0x6f, 0xa4, 0xca, 0x5b, 0x8c,
	0x69, 0x09, 0x14, 0x90, 0x63, 0x19, 0xb1, 0x13, 0x25, 0xf7, 0x98, 0xd0, 0x39, 0x2c, 0x83, 0xdf,
	0xa8, 0x8e, 0x47, 0xd0, 0x6a, 0x27, 0x43, 0xdc, 0xe4, 0x1f, 0x16, 0xf2, 0xd2, 0xe2, 0x11, 0x66,
	0x83, 0x25, 0x9b, 0x15, 0xd3, 0x1c, 0x8e, 0xb4, 0x59, 0xab, 0x56, 0x57, 0x18, 0x85, 0x43, 0xa7,
	0xca, 0x1b, 0x6d, 0x18, 0x63, 0x02, 0xc8, 0x4b, 0x36, 0xc2, 0x16, 0x13, 0x9a, 0x42, 0x1a, 0xa2,
	0x70, 0x32, 0x2a, 0x61, 0x4c, 0x69, 0x06, 0x59, 0xdb, 0x37, 0xda, 0x60, 0x36, 0x5c, 0x6d, 0x35,
	0x1b, 0xc1, 0xbc, 0xb8, 0x06, 0xf8, 0x5d, 0x49, 0xc7, 0xa1, 0x2b, 0x8c, 0x52, 0xd5, 0x06, 0xa3,
	0xf0, 0x66, 0x94, 0x31, 0x2d, 0x60, 0xda, 0xd7, 0xb5, 0x6f, 0x43, 0x59, 0x52, 0x5c, 0xfc, 0x6c,
	0xd2, 0xa6, 0xee, 0xe9, 0x04, 0xe6, 0x8e, 0xed, 0x9a, 0xad, 0x5f, 0xf5, 0x56, 0x30, 0xda, 0xe6,
	0xc3, 0x67, 0x5e, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x32, 0x2b, 0x1d, 0xd6, 0x5c, 0x01, 0x00,
	0x00,
}
