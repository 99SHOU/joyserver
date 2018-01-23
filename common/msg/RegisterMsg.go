package msg

import (
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
)

var (
	Processor = net.NewProcessor()
)

func init() {
	RegisterMsg(Processor)
}

// all of msg must register in this func
func RegisterMsg(processor *net.Processor) {
	processor.Register(uint16(pb.MsgID_FWM_NODE_REGISTER_REQ), &pb.NodeRegisterReq{})
	processor.Register(uint16(pb.MsgID_FWM_NODE_REGISTER_ACK), &pb.NodeRegisterAck{})
}
