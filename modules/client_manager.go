package modules

import (
	"github.com/99SHOU/joyserver/common/net"
	"github.com/name5566/leaf/log"
)

type ClientManager struct {
	clients map[string]*net.Client
}

func (cm *ClientManager) Init() {
	cm.clients = make(map[string]*net.Client)
}

func (cm *ClientManager) Destroy() {
	for _, client := range cm.clients {
		client.Close()
	}
}

func (cm *ClientManager) Run() {

}

func (cm *ClientManager) NewAndStart(addr string, clientHandler net.ClientHandler, processor *net.Processor) {
	if _, ok := cm.clients[addr]; ok {
		log.Error("Client is exist: Addr: %v", addr)
		return
	}

	client := net.NewClient(addr, clientHandler, processor)
	client.Start()
	cm.clients[addr] = client
}

func (cm *ClientManager) GetClient(addr string) *net.Client {
	if _, ok := cm.clients[addr]; !ok {
		return nil
	}

	return cm.clients[addr]
}

func (cm *ClientManager) ExistClient(addr string) bool {
	if _, ok := cm.clients[addr]; ok {
		return true
	} else {
		return false
	}
}
