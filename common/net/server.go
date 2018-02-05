package net

import (
	"github.com/99SHOU/joyserver/common/conf"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"strconv"
)

func NewServer(port uint32, serverHandler ServerHandler, processor *Processor) Server {
	tcpAddr := "127.0.0.1" + ":" + strconv.FormatUint(uint64(port), 10)

	server := Server{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		Processor:       processor,

		TCPAddr:      tcpAddr,
		LenMsgLen:    conf.LenMsgLen,
		LittleEndian: conf.LittleEndian,
	}

	server.Processor.SetByteOrder(server.LittleEndian)
	serverHandler.Register(&server)

	return server
}

type Server struct {
	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	Processor       *Processor

	// tcp
	TCPAddr      string
	LenMsgLen    int
	LittleEndian bool
	TcpServer    *network.TCPServer

	OnNewAgent   func(Agent)
	OnCloseAgent func(Agent)
}

func (server *Server) SetHandler(id pb.MsgID, handler MsgHandler) {
	server.Processor.SetHandler(uint16(id), handler)
}

func (server *Server) SetOtherHandler(handler MsgHandler) {
	server.Processor.SetOtherHandler(handler)
}

func (server *Server) Start() {
	if server.TcpServer != nil {
		log.Error("server is start!")
		return
	}

	var tcpServer *network.TCPServer
	if server.TCPAddr != "" {
		tcpServer = new(network.TCPServer)
		tcpServer.Addr = server.TCPAddr
		tcpServer.MaxConnNum = server.MaxConnNum
		tcpServer.PendingWriteNum = server.PendingWriteNum
		tcpServer.LenMsgLen = server.LenMsgLen
		tcpServer.MaxMsgLen = server.MaxMsgLen
		tcpServer.LittleEndian = server.LittleEndian
		tcpServer.NewAgent = func(conn *network.TCPConn) network.Agent {
			agent := &BaseAgent{onCloseAgent: server.OnCloseAgent, conn: conn, processor: server.Processor, agentInfo: NewAgentInfo()}
			server.OnNewAgent(agent)
			return agent
		}
	}

	server.TcpServer = tcpServer

	if tcpServer != nil {
		tcpServer.Start()
	}
}

func (server *Server) Close() {
	if server.TcpServer != nil {
		server.TcpServer.Close()
	}

	server.TcpServer = nil
}
