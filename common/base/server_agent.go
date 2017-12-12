package base

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"net"
	"reflect"
)

type Agent struct {
	conn     network.Conn
	server   *Server
	userData interface{}
}

func (a *Agent) Run() {
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

func (a *Agent) OnClose() {
	a.server.OnCloseAgent(a)
}

func (a *Agent) WriteMsg(msg interface{}) {
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

func (a *Agent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *Agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *Agent) Close() {
	a.conn.Close()
}

func (a *Agent) Destroy() {
	a.conn.Destroy()
}

func (a *Agent) UserData() interface{} {
	return a.userData
}

func (a *Agent) SetUserData(data interface{}) {
	a.userData = data
}
