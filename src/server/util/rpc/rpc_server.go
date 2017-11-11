package rpc

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"net"
	"reflect"
)

var ()

type RpcMap map[string]reflect.Value

type RpcServer struct {
	rpcMap    RpcMap
	tcpServer *network.TCPServer

	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	Processor       network.Processor

	TCPAddr      string
	LenMsgLen    int
	LittleEndian bool
}

func (rpcServer *RpcServer) OnInit(rpcObj interface{}) {
	rpcServer.rpcMap = make(RpcMap, 0)

	if rpcServer.TCPAddr != "" {
		rpcServer.tcpServer = new(network.TCPServer)
		rpcServer.tcpServer.Addr = rpcServer.TCPAddr
		rpcServer.tcpServer.MaxConnNum = rpcServer.MaxConnNum
		rpcServer.tcpServer.PendingWriteNum = rpcServer.PendingWriteNum
		rpcServer.tcpServer.LenMsgLen = rpcServer.LenMsgLen
		rpcServer.tcpServer.MaxMsgLen = rpcServer.MaxMsgLen
		rpcServer.tcpServer.LittleEndian = rpcServer.LittleEndian
		rpcServer.tcpServer.NewAgent = func(conn *network.TCPConn) network.Agent {
			a := &agent{conn: conn, rpcServer: rpcServer}
			return a
		}
	}

	if rpcServer.tcpServer != nil {
		rpcServer.tcpServer.Start()
	}
}

func (rpcServer *RpcServer) AddRpcObj(rpcObj interface{}) {
	rt := reflect.TypeOf(rpcObj)
	rv := reflect.ValueOf(rpcObj)
	for i := 0; i < rt.NumMethod(); i++ {
		rpcServer.rpcMap[rt.Method(i).Name] = rv.Method(i)
	}
}

func (rpcServer *RpcServer) Call(method string, args []interface{}) {
	parms := []reflect.Value{}
	for i := 0; i < len(args); i++ {
		parms = append(parms, reflect.ValueOf(args[i]))
	}

	rpcServer.rpcMap[method].Call(parms)
}

func (rpcServer *RpcServer) OnDestroy() {
	rpcServer.rpcMap = nil

	if rpcServer.tcpServer != nil {
		rpcServer.tcpServer.Close()
	}
}

func (rpcServer *RpcServer) Run() {

}

type agent struct {
	conn      network.Conn
	rpcServer *RpcServer
	userData  interface{}
}

func (a *agent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}

		if a.rpcServer.Processor != nil {
			msg, err := a.rpcServer.Processor.Unmarshal(data)
			if err != nil {
				log.Debug("unmarshal message error: %v", err)
				break
			}
			err = a.rpcServer.Processor.Route(msg, a)
			if err != nil {
				log.Debug("route message error: %v", err)
				break
			}
		}
	}
}

func (a *agent) OnClose() {

}

func (a *agent) WriteMsg(msg interface{}) {
	if a.rpcServer.Processor != nil {
		data, err := a.rpcServer.Processor.Marshal(msg)
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

func (a *agent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *agent) Close() {
	a.conn.Close()
}

func (a *agent) Destroy() {
	a.conn.Destroy()
}

func (a *agent) UserData() interface{} {
	return a.userData
}

func (a *agent) SetUserData(data interface{}) {
	a.userData = data
}
