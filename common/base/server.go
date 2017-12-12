package base

import (
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/network"
	//"github.com/name5566/leaf/network/protobuf"
	"github.com/99SHOU/joyserver/common/conf"
	"github.com/99SHOU/joyserver/common/pb"
)

type ServerHandler interface {
	Init(server *Server)
}

func NewServer(tcpAddr string) *Server {
	skeleton := NewSkeleton()
	server := &Server{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		Processor:       NewProcessor(),

		TCPAddr:      tcpAddr,
		LenMsgLen:    conf.LenMsgLen,
		LittleEndian: conf.LittleEndian,

		skeleton:     skeleton,
		agentChanRPC: skeleton.ChanRPCServer,

		functions: make(map[interface{}]func(interface{}, interface{})),
	}

	server.Processor.SetByteOrder(server.LittleEndian)

	return server
}

type Server struct {
	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	Processor       *Processor
	agentChanRPC    *chanrpc.Server

	// tcp
	TCPAddr      string
	LenMsgLen    int
	LittleEndian bool
	TcpServer    *network.TCPServer

	skeleton *module.Skeleton

	OnNewAgent   func(*Agent)
	OnCloseAgent func(*Agent)
	functions    map[interface{}]func(interface{}, interface{})
}

func (server *Server) Register(id pb.EGameMsgID, msg interface{}) {
	server.Processor.Register(uint16(id), msg)
}

func (server *Server) SetHandler(id pb.EGameMsgID, handler MsgHandler) {
	server.Processor.SetHandler(uint16(id), handler)
}

func (server *Server) RegisterAndSetHandler(id pb.EGameMsgID, msg interface{}, handler MsgHandler) {
	server.Register(id, msg)
	server.SetHandler(id, handler)
}

func (server *Server) Start() {
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
			a := &Agent{conn: conn, server: server}
			server.OnNewAgent(a)
			return a
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
}
