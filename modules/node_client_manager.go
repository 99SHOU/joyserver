package modules

import (
	"github.com/99SHOU/joyserver/common/net"
	"github.com/name5566/leaf/log"
)

type NodeClientManager struct {
	clients map[string]*net.Client
}

func (cm *NodeClientManager) Init() {
	cm.clients = make(map[string]*net.Client)
}

func (cm *NodeClientManager) Destroy() {
	for _, client := range cm.clients {
		client.Close()
	}
}

func (cm *NodeClientManager) Run() {

}

func (cm *NodeClientManager) NewAndStart(addr string, clientHandler net.ClientHandler, processor *net.Processor) {
	if _, ok := cm.clients[addr]; ok {
		log.Error("Client is exist: Addr: %v", addr)
		return
	}

	client := net.NewClient(addr, clientHandler, processor)
	client.Start()
	cm.clients[addr] = client
}

func (cm *NodeClientManager) GetClient(addr string) *net.Client {
	if _, ok := cm.clients[addr]; !ok {
		return nil
	}

	return cm.clients[addr]
}

func (cm *NodeClientManager) ExistClient(addr string) bool {
	if _, ok := cm.clients[addr]; ok {
		return true
	} else {
		return false
	}
}
