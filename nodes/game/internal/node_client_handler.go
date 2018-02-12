package internal

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/msg"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/modules"
	"github.com/name5566/leaf/log"
)

type NodeClientHandler struct {
	Node         *Node
	agentManager *modules.AgentManager
}

func (h *NodeClientHandler) Register(client *net.Client) {
	h.agentManager = h.Node.ModulesManager.Find("AgentManager").(*modules.AgentManager)

	client.OnNewAgent = h.NewAgent
	client.OnCloseAgent = h.CloseAgent

	msg.RegisterMsg(client.Processor)
	client.SetHandler(pb.MsgID_FWM_NODE_REGISTER_ACK, h.OnNodeRegisterAck)
}

func (h *NodeClientHandler) NewAgent(agent net.Agent) {
	h.agentManager.AddAgent(agent.RemoteAddr().String(), agent)

	h.NodeRegisterReq(agent)
}

func (h *NodeClientHandler) CloseAgent(agent net.Agent) {
	h.agentManager.RemoveAgent(agent.RemoteAddr().String())
}

func (h *NodeClientHandler) NodeRegisterReq(agent net.Agent) {
	agent.WriteMsg(&pb.NodeRegisterReq{NodeType: h.Node.NodeType, NodeStatu: h.Node.NodeStatu, NodeId: uint32(h.Node.NodeID), NodePort: h.Node.NodeCfg.NodePort})
}

func (h *NodeClientHandler) OnNodeRegisterAck(message interface{}, agent interface{}) {
	msg := message.(*pb.NodeRegisterAck)
	a := agent.(*net.BaseAgent)
	a.SetNodeType(msg.NodeType)
	a.SetNodeStatu(msg.NodeStatu)
	a.SetNodeID(define.NodeID(msg.NodeId))

	if a.GetNodeType() == pb.NODE_TYPE_CENTER {
		h.Node.NodeStatu = pb.NODE_STATU_READY
		allAgent := h.agentManager.GetAgentAll(nil)
		net.BroadcastMsg(allAgent, &pb.SetNodeStatu{NodeStatu: h.Node.NodeStatu})
	}

	log.Debug("Register to NodeType: %v NodeID: %v success", msg.NodeType.String(), msg.NodeId)
}
