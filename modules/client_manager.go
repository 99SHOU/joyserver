package modules

import (
	"github.com/99SHOU/joyserver/common/net"
)

type ClientManager struct {
	clients []net.Client
}

func (cm *ClientManager) Init() {
	cm.clients = []net.Client{}
}

func (cm *ClientManager) Destroy() {
	for _, client := range cm.clients {
		client.Close()
	}
}

func (cm *ClientManager) Run() {

}

func (cm *ClientManager) NewAndStart(addr string, clientHandler net.ClientHandler, processor *net.Processor) {
	client := net.NewClient(addr, clientHandler, processor)
	client.Start()
	cm.clients = append(cm.clients, client)
}
