package net

import (
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"net"
	"reflect"
)

// AgentInterface
type Agent interface {
	WriteMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()

	GetNodeType() pb.NODE_TYPE
	SetNodeType(nodeType pb.NODE_TYPE)
	GetAgentInfo(key interface{}) interface{}
	SetAgentInfo(key interface{}, value interface{})
}

// ClientAgent
type ClientAgent struct {
	BaseAgent
	Client       *Client
	onCloseAgent func(*ClientAgent)
}

func (a *ClientAgent) OnClose() {
	a.onCloseAgent(a)
}

// ServerAgent
type ServerAgent struct {
	BaseAgent
	Server       *Server
	onCloseAgent func(*ServerAgent)
}

func (a *ServerAgent) OnClose() {
	a.onCloseAgent(a)
}

// BaseAgent
type BaseAgent struct {
	conn      network.Conn
	processor *Processor
	agentInfo *AgentInfo
}

func (a *BaseAgent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			if err.Error() != "EOF" {
				log.Debug("read message: %v", err)
			}
			break
		}

		if a.processor != nil {
			msg, err := a.processor.Unmarshal(data)
			if err != nil {
				log.Debug("unmarshal message error: %v  %v", err, data)
				break
			}

			a.processor.Dispatch(msg, a)
		}
	}
}

func (a *BaseAgent) WriteMsg(msg interface{}) {
	if a.processor != nil {
		data, err := a.processor.Marshal(msg)
		if err != nil {
			log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
			return
		}
		err = a.conn.WriteMsg(data...)
		if err != nil {
			log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}

func (a *BaseAgent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *BaseAgent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *BaseAgent) Close() {
	a.conn.Close()
}

func (a *BaseAgent) Destroy() {
	a.conn.Destroy()
}

func (a *BaseAgent) GetNodeType() pb.NODE_TYPE {
	return a.agentInfo.NodeType
}

func (a *BaseAgent) SetNodeType(nodeType pb.NODE_TYPE) {
	a.agentInfo.NodeType = nodeType
}

func (a *BaseAgent) GetAgentInfo(key interface{}) interface{} {
	return a.GetAgentInfo(key)
}

func (a *BaseAgent) SetAgentInfo(key interface{}, value interface{}) {
	a.agentInfo.SetUserData(key, value)
}