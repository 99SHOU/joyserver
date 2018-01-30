package internal

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/msg"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
	"strconv"
	"strings"
)

type ServerHandler struct {
	Node *Node
}

func (sh *ServerHandler) Register(server *net.Server) {
	server.OnNewAgent = sh.NewAgent
	server.OnCloseAgent = sh.CloseAgent

	msg.RegisterMsg(server.Processor)
	server.SetHandler(pb.MsgID_FWM_NODE_REGISTER_REQ, sh.OnNodeRegisterReq)
	server.SetHandler(pb.MsgID_FWM_SET_NODE_STATU, sh.OnSetNodeStatu)
	server.SetHandler(pb.MsgID_FWM_GAME_NODE_LIST_REQ, sh.OnGameNodeListReq)
}

func (sh *ServerHandler) NewAgent(agent net.Agent) {
	sh.Node.AgentManager.AddAgent(agent.RemoteAddr().String(), agent)
}

func (sh *ServerHandler) CloseAgent(agent net.Agent) {
	sh.Node.AgentManager.RemoveAgent(agent.RemoteAddr().String())
}

func (sh *ServerHandler) OnNodeRegisterReq(message interface{}, agent interface{}) {
	msg := message.(*pb.NodeRegisterReq)
	a := agent.(*net.BaseAgent)

	a.SetNodeType(msg.NodeType)
	a.SetNodeStatu(pb.NODE_STATU_NOT_READY)
	a.SetNodeID(define.NodeID(msg.NodeId))
	a.SetAgentInfo(pb.NODE_INFO_SERVER_PORT, msg.ServerPort)

	log.Debug("New node register to center NodeType: %v NodeID: %v", msg.NodeType.String(), msg.NodeId)

	a.WriteMsg(&pb.NodeRegisterAck{NodeType: sh.Node.NodeType, NodeId: uint32(sh.Node.NodeID)})
}

func (sh *ServerHandler) OnSetNodeStatu(message interface{}, agent interface{}) {
	msg := message.(*pb.SetNodeStatu)
	a := agent.(*net.BaseAgent)

	a.SetNodeStatu(msg.NodeStatu)

	log.Debug("Change node statu : NodeId: %v NodeType: %v NodeStatu: %v", a.GetNodeID(), a.GetNodeType(), a.GetNodeStatu())

	if a.GetNodeType() == pb.NODE_TYPE_GAME {
		gateAgent := sh.Node.AgentManager.GetAgentByNodeType([]pb.NODE_TYPE{pb.NODE_TYPE_GATE})
		net.BroadcastMsg(gateAgent, sh.BuildGameNodeList())
	}
}

func (sh *ServerHandler) OnGameNodeListReq(message interface{}, agent interface{}) {
	a := agent.(*net.BaseAgent)

	a.WriteMsg(sh.BuildGameNodeList())
}

func (sh *ServerHandler) BuildGameNodeList() *pb.GameNodeListAck {
	gameAgents := sh.Node.AgentManager.GetAgentByNodeType([]pb.NODE_TYPE{pb.NODE_TYPE_GAME})
	gameNodeList := pb.GameNodeListAck{}
	gameNodeList.NodeInfos = []*pb.NodeInfo{}

	for _, agent := range gameAgents {
		if agent.GetNodeStatu() == pb.NODE_STATU_READY {

			serverPort := agent.GetAgentInfo(pb.NODE_INFO_SERVER_PORT).(uint32)
			serverAddr := strings.Split(agent.RemoteAddr().String(), ":")[0] + ":" + strconv.FormatUint(uint64(serverPort), 10)

			gameNodeList.NodeInfos = append(gameNodeList.NodeInfos, &pb.NodeInfo{NodeType: agent.GetNodeType(), NodeStatu: agent.GetNodeStatu(), NodeId: uint32(agent.GetNodeID()), Addr: serverAddr})
		}
	}

	return &gameNodeList
}
