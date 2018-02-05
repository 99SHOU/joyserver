package msg

import (
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
)

// all of msg must register in this func
func RegisterMsg(processor *net.Processor) {
	processor.Register(uint16(pb.MsgID_FWM_NODE_REGISTER_REQ), &pb.NodeRegisterReq{})
	processor.Register(uint16(pb.MsgID_FWM_NODE_REGISTER_ACK), &pb.NodeRegisterAck{})
	processor.Register(uint16(pb.MsgID_FWM_SET_NODE_STATU), &pb.SetNodeStatu{})
	processor.Register(uint16(pb.MsgID_FWM_GAME_NODE_LIST_REQ), &pb.GameNodeListReq{})
	processor.Register(uint16(pb.MsgID_FWM_GAME_NODE_LIST_ACK), &pb.GameNodeListAck{})
	processor.Register(uint16(pb.MsgID_FWM_GAME_MSG_TRANSFER), &pb.GameMsgTransfer{})
}
