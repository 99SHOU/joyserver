package mgr

import (
	"errors"
	"github.com/name5566/leaf/log"
	"net"
	"server/common/define"
	"strconv"
	"sync"
)

type PortMgr struct {
	mutex           sync.Mutex
	ServerStartPort map[define.SERVER_TYPE]int
}

func (mgr *PortMgr) OnInit() {
	mgr.ServerStartPort = make(map[define.SERVER_TYPE]int)
	mgr.ServerStartPort[define.SERVER_TYPE_LOGIN] = define.SERVER_PORT_START_LOGIN
	mgr.ServerStartPort[define.SERVER_TYPE_GATE] = define.SERVER_PORT_START_GATE
	mgr.ServerStartPort[define.SERVER_TYPE_LOGIC] = define.SERVER_PORT_START_LOGIC
}

func (mgr *PortMgr) OnDestroy() {

}

func (mgr *PortMgr) Run() {

}

func (mgr *PortMgr) GetModulePort(serverType define.SERVER_TYPE) (int, error) {
	port := 0
	err := errors.New("get module port fail")

	if serverType == define.SERVER_TYPE_CENTER {
		return define.CENTER_PORT, nil
	}

	mgr.mutex.Lock()
	for i := 0; i < 1000; i++ {
		port = mgr.ServerStartPort[serverType] + i
		if mgr.CheckPort(port) {
			err = nil
			break
		}
	}
	mgr.mutex.Unlock()

	return port, err
}

func (mgr *PortMgr) CheckPort(port int) bool {
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err == nil {
		conn.Close()
		return true
	}

	log.Debug("%v  %v", err, port)

	return false
}
