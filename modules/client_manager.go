package modules

import (
	"github.com/99SHOU/joyserver/common/net"
)

type ClientManager struct {
	clients map[string]*net.Client
}

func (cm *ClientManager) Init() {
	cm.clients = make(map[string]*net.Client)
}

func (cm *ClientManager) Destroy() {

}

func (cm *ClientManager) Run() {

}

func (cm *ClientManager) CreateClient() {

}
