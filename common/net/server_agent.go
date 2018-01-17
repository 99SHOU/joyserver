package net

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"net"
	"reflect"
)

type ServerAgent struct {
	conn     network.Conn
	server   *Server
	userData interface{}
}

func (a *ServerAgent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			if err.Error() != "EOF" {
				log.Debug("read message: %v", err)
			}
			break
		}

		if a.server.Processor != nil {
			msg, err := a.server.Processor.Unmarshal(data)
			if err != nil {
				log.Debug("unmarshal message error: %v  %v", err, data)
				break
			}

			a.server.Processor.Dispatch(msg, a)
		}
	}
}

func (a *ServerAgent) OnClose() {
	a.server.OnCloseAgent(a)
}

func (a *ServerAgent) WriteMsg(msg interface{}) {
	if a.server.Processor != nil {
		data, err := a.server.Processor.Marshal(msg)
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

func (a *ServerAgent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *ServerAgent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *ServerAgent) Close() {
	a.conn.Close()
}

func (a *ServerAgent) Destroy() {
	a.conn.Destroy()
}

func (a *ServerAgent) UserData() interface{} {
	return a.userData
}

func (a *ServerAgent) SetUserData(data interface{}) {
	a.userData = data
}
