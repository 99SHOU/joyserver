// Code generated by protoc-gen-go. DO NOT EDIT.
// source: framework_msg.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NodeRegisterReq struct {
	NodeType  NODE_TYPE  `protobuf:"varint,1,opt,name=NodeType,enum=pb.NODE_TYPE" json:"NodeType,omitempty"`
	NodeStatu NODE_STATU `protobuf:"varint,2,opt,name=NodeStatu,enum=pb.NODE_STATU" json:"NodeStatu,omitempty"`
	NodeId    uint32     `protobuf:"varint,3,opt,name=NodeId" json:"NodeId,omitempty"`
	NodePort  uint32     `protobuf:"varint,4,opt,name=NodePort" json:"NodePort,omitempty"`
}

func (m *NodeRegisterReq) Reset()                    { *m = NodeRegisterReq{} }
func (m *NodeRegisterReq) String() string            { return proto.CompactTextString(m) }
func (*NodeRegisterReq) ProtoMessage()               {}
func (*NodeRegisterReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *NodeRegisterReq) GetNodeType() NODE_TYPE {
	if m != nil {
		return m.NodeType
	}
	return NODE_TYPE_INVALID
}

func (m *NodeRegisterReq) GetNodeStatu() NODE_STATU {
	if m != nil {
		return m.NodeStatu
	}
	return NODE_STATU_NOT_READY
}

func (m *NodeRegisterReq) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeRegisterReq) GetNodePort() uint32 {
	if m != nil {
		return m.NodePort
	}
	return 0
}

type NodeRegisterAck struct {
	NodeType  NODE_TYPE  `protobuf:"varint,1,opt,name=NodeType,enum=pb.NODE_TYPE" json:"NodeType,omitempty"`
	NodeStatu NODE_STATU `protobuf:"varint,2,opt,name=NodeStatu,enum=pb.NODE_STATU" json:"NodeStatu,omitempty"`
	NodeId    uint32     `protobuf:"varint,3,opt,name=NodeId" json:"NodeId,omitempty"`
}

func (m *NodeRegisterAck) Reset()                    { *m = NodeRegisterAck{} }
func (m *NodeRegisterAck) String() string            { return proto.CompactTextString(m) }
func (*NodeRegisterAck) ProtoMessage()               {}
func (*NodeRegisterAck) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *NodeRegisterAck) GetNodeType() NODE_TYPE {
	if m != nil {
		return m.NodeType
	}
	return NODE_TYPE_INVALID
}

func (m *NodeRegisterAck) GetNodeStatu() NODE_STATU {
	if m != nil {
		return m.NodeStatu
	}
	return NODE_STATU_NOT_READY
}

func (m *NodeRegisterAck) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

type SetNodeStatu struct {
	NodeStatu NODE_STATU `protobuf:"varint,1,opt,name=NodeStatu,enum=pb.NODE_STATU" json:"NodeStatu,omitempty"`
}

func (m *SetNodeStatu) Reset()                    { *m = SetNodeStatu{} }
func (m *SetNodeStatu) String() string            { return proto.CompactTextString(m) }
func (*SetNodeStatu) ProtoMessage()               {}
func (*SetNodeStatu) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *SetNodeStatu) GetNodeStatu() NODE_STATU {
	if m != nil {
		return m.NodeStatu
	}
	return NODE_STATU_NOT_READY
}

type NodeInfo struct {
	NodeType  NODE_TYPE  `protobuf:"varint,1,opt,name=NodeType,enum=pb.NODE_TYPE" json:"NodeType,omitempty"`
	NodeStatu NODE_STATU `protobuf:"varint,2,opt,name=NodeStatu,enum=pb.NODE_STATU" json:"NodeStatu,omitempty"`
	NodeId    uint32     `protobuf:"varint,3,opt,name=NodeId" json:"NodeId,omitempty"`
	Addr      string     `protobuf:"bytes,4,opt,name=Addr" json:"Addr,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *NodeInfo) GetNodeType() NODE_TYPE {
	if m != nil {
		return m.NodeType
	}
	return NODE_TYPE_INVALID
}

func (m *NodeInfo) GetNodeStatu() NODE_STATU {
	if m != nil {
		return m.NodeStatu
	}
	return NODE_STATU_NOT_READY
}

func (m *NodeInfo) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type GameNodeListReq struct {
}

func (m *GameNodeListReq) Reset()                    { *m = GameNodeListReq{} }
func (m *GameNodeListReq) String() string            { return proto.CompactTextString(m) }
func (*GameNodeListReq) ProtoMessage()               {}
func (*GameNodeListReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type GameNodeListAck struct {
	NodeInfos []*NodeInfo `protobuf:"bytes,1,rep,name=NodeInfos" json:"NodeInfos,omitempty"`
}

func (m *GameNodeListAck) Reset()                    { *m = GameNodeListAck{} }
func (m *GameNodeListAck) String() string            { return proto.CompactTextString(m) }
func (*GameNodeListAck) ProtoMessage()               {}
func (*GameNodeListAck) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *GameNodeListAck) GetNodeInfos() []*NodeInfo {
	if m != nil {
		return m.NodeInfos
	}
	return nil
}

type GameMsgTransfer struct {
	CharacterId uint64 `protobuf:"varint,1,opt,name=CharacterId" json:"CharacterId,omitempty"`
	MsgId       []byte `protobuf:"bytes,2,opt,name=MsgId,proto3" json:"MsgId,omitempty"`
	MsgBody     []byte `protobuf:"bytes,3,opt,name=MsgBody,proto3" json:"MsgBody,omitempty"`
}

func (m *GameMsgTransfer) Reset()                    { *m = GameMsgTransfer{} }
func (m *GameMsgTransfer) String() string            { return proto.CompactTextString(m) }
func (*GameMsgTransfer) ProtoMessage()               {}
func (*GameMsgTransfer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *GameMsgTransfer) GetCharacterId() uint64 {
	if m != nil {
		return m.CharacterId
	}
	return 0
}

func (m *GameMsgTransfer) GetMsgId() []byte {
	if m != nil {
		return m.MsgId
	}
	return nil
}

func (m *GameMsgTransfer) GetMsgBody() []byte {
	if m != nil {
		return m.MsgBody
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeRegisterReq)(nil), "pb.NodeRegisterReq")
	proto.RegisterType((*NodeRegisterAck)(nil), "pb.NodeRegisterAck")
	proto.RegisterType((*SetNodeStatu)(nil), "pb.SetNodeStatu")
	proto.RegisterType((*NodeInfo)(nil), "pb.NodeInfo")
	proto.RegisterType((*GameNodeListReq)(nil), "pb.GameNodeListReq")
	proto.RegisterType((*GameNodeListAck)(nil), "pb.GameNodeListAck")
	proto.RegisterType((*GameMsgTransfer)(nil), "pb.GameMsgTransfer")
}

func init() { proto.RegisterFile("framework_msg.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xc1, 0x4e, 0x02, 0x31,
	0x10, 0x4d, 0x01, 0x51, 0xca, 0x02, 0xb1, 0x1a, 0xb3, 0xe1, 0xb4, 0xe9, 0x09, 0x8d, 0xe1, 0x80,
	0x57, 0x3d, 0xa0, 0x12, 0xb3, 0x89, 0x20, 0x29, 0xeb, 0xc1, 0x13, 0x59, 0xe8, 0xec, 0x4a, 0x08,
	0x74, 0x69, 0x6b, 0x0c, 0x57, 0x7f, 0xc0, 0x3f, 0xf0, 0x5b, 0x4d, 0x5b, 0x64, 0xc5, 0x8b, 0x37,
	0x6e, 0xf3, 0xde, 0xbc, 0xce, 0xbc, 0x79, 0x29, 0x3e, 0x49, 0x64, 0xbc, 0x80, 0x77, 0x21, 0xe7,
	0xe3, 0x85, 0x4a, 0xdb, 0x99, 0x14, 0x5a, 0x90, 0x42, 0x36, 0x69, 0x7a, 0x1c, 0x92, 0xd9, 0x12,
	0x1c, 0x43, 0xbf, 0x10, 0x6e, 0x0c, 0x04, 0x07, 0x06, 0xe9, 0x4c, 0x69, 0x90, 0x0c, 0x56, 0xe4,
	0x1c, 0x1f, 0x19, 0x2a, 0x5a, 0x67, 0xe0, 0xa3, 0x00, 0xb5, 0xea, 0x9d, 0x5a, 0x3b, 0x9b, 0xb4,
	0x07, 0x4f, 0xf7, 0xbd, 0x71, 0xf4, 0x32, 0xec, 0xb1, 0x6d, 0x9b, 0x5c, 0xe2, 0x8a, 0xa9, 0x47,
	0x3a, 0xd6, 0x6f, 0x7e, 0xc1, 0x6a, 0xeb, 0x5b, 0xed, 0x28, 0xea, 0x46, 0xcf, 0x2c, 0x17, 0x90,
	0x33, 0x5c, 0x36, 0x20, 0xe4, 0x7e, 0x31, 0x40, 0xad, 0x1a, 0xdb, 0x20, 0xd2, 0x74, 0x0b, 0x87,
	0x42, 0x6a, 0xbf, 0x64, 0x3b, 0x5b, 0x4c, 0x3f, 0xfe, 0x18, 0xec, 0x4e, 0xe7, 0x7b, 0x37, 0x48,
	0xaf, 0xb1, 0x37, 0x02, 0x9d, 0xeb, 0x76, 0xa6, 0xa2, 0x7f, 0xa6, 0xd2, 0x4f, 0xe4, 0xfc, 0x86,
	0xcb, 0x44, 0xec, 0x3f, 0x5c, 0x82, 0x4b, 0x5d, 0xce, 0xa5, 0x0d, 0xb6, 0xc2, 0x6c, 0x4d, 0x8f,
	0x71, 0xe3, 0x21, 0x5e, 0x80, 0x51, 0x3c, 0xce, 0x94, 0x66, 0xb0, 0xa2, 0x37, 0xbb, 0x94, 0x89,
	0xf9, 0xc2, 0xed, 0x37, 0xb6, 0x95, 0x8f, 0x82, 0x62, 0xab, 0xda, 0xf1, 0xec, 0xfe, 0x0d, 0xc9,
	0xf2, 0x36, 0x9d, 0xba, 0xe7, 0x7d, 0x95, 0x46, 0x32, 0x5e, 0xaa, 0x04, 0x24, 0x09, 0x70, 0xf5,
	0xee, 0x35, 0x96, 0xf1, 0x54, 0x83, 0x0c, 0xb9, 0x3d, 0xb6, 0xc4, 0x7e, 0x53, 0xe4, 0x14, 0x1f,
	0xf4, 0x55, 0x1a, 0x72, 0x7b, 0x9c, 0xc7, 0x1c, 0x20, 0x3e, 0x3e, 0xec, 0xab, 0xf4, 0x56, 0xf0,
	0xb5, 0xbd, 0xc4, 0x63, 0x3f, 0x70, 0x52, 0xb6, 0x7f, 0xf6, 0xea, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x4b, 0x48, 0x1e, 0x20, 0xdc, 0x02, 0x00, 0x00,
}
