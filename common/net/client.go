package net

import (
	"github.com/99SHOU/joyserver/common/conf"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"time"
)

func NewClient(addr string, clientHandler ClientHandler, processor *Processor) Client {
	client := Client{
		Addr:            addr,
		ConnNum:         1,
		ConnectInterval: 3 * time.Second,
		PendingWriteNum: conf.PendingWriteNum,
		AutoReconnect:   false,
		LenMsgLen:       conf.LenMsgLen,
		MaxMsgLen:       conf.MaxMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       processor,
		AgentInfo:       NewAgentInfo(),
	}

	client.Processor.SetByteOrder(client.LittleEndian)
	clientHandler.Register(&client)

	return client
}

type Client struct {
	Addr            string
	ConnNum         int
	ConnectInterval time.Duration
	PendingWriteNum int
	AutoReconnect   bool
	LenMsgLen       int
	MinMsgLen       uint32
	MaxMsgLen       uint32
	LittleEndian    bool
	Processor       *Processor

	TcpClient *network.TCPClient
	Agent     *ClientAgent
	AgentInfo *AgentInfo

	OnNewAgent   func(*ClientAgent)
	OnCloseAgent func(*ClientAgent)
}

func (client *Client) SetHandler(id pb.MsgID, handler MsgHandler) {
	client.Processor.SetHandler(uint16(id), handler)
}

func (client *Client) Start() {
	if client.TcpClient != nil {
		log.Error("client is start!")
		return
	}

	var tcpClient *network.TCPClient
	if client.Addr != "" {
		tcpClient = new(network.TCPClient)
		tcpClient.Addr = client.Addr
		tcpClient.ConnNum = client.ConnNum
		tcpClient.ConnectInterval = client.ConnectInterval
		tcpClient.PendingWriteNum = client.PendingWriteNum
		tcpClient.AutoReconnect = client.AutoReconnect
		tcpClient.LenMsgLen = client.LenMsgLen
		tcpClient.MaxMsgLen = client.MaxMsgLen
		tcpClient.LittleEndian = client.LittleEndian
		tcpClient.NewAgent = func(conn *network.TCPConn) network.Agent {
			a := &ClientAgent{Client: client, onCloseAgent: client.OnCloseAgent, BaseAgent: BaseAgent{conn: conn, processor: client.Processor, agentInfo: client.AgentInfo}}
			client.Agent = a

			if client.OnNewAgent == nil {
				log.Error("Must set OnNewAgent")
			}

			if a == nil {
				log.Error("Agent is nil")
			}

			client.OnNewAgent(a)
			return a
		}
	}

	client.TcpClient = tcpClient

	if tcpClient != nil {
		tcpClient.Start()
	}
}

func (client *Client) Close() {
	if client.TcpClient != nil {
		client.TcpClient.Close()
	}

	client.TcpClient = nil
	client.Agent = nil
	client.AgentInfo = nil
}
