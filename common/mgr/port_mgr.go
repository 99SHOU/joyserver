package mgr

import (
	"errors"
	"github.com/name5566/leaf/log"
	"net"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/pb"
	"strconv"
	"sync"
)

type PortMgr struct {
	mutex              sync.Mutex
	RpcServerStartPort map[pb.SERVER_TYPE]uint32
	ServerStartPort    map[pb.SERVER_TYPE]uint32
}

func (mgr *PortMgr) OnInit() {
	mgr.RpcServerStartPort = make(map[pb.SERVER_TYPE]uint32)
	mgr.RpcServerStartPort[pb.SERVER_TYPE_LOGIN] = define.SERVER_RPC_PORT_START_LOGIN
	mgr.RpcServerStartPort[pb.SERVER_TYPE_GATE] = define.SERVER_RPC_PORT_START_GATE
	mgr.RpcServerStartPort[pb.SERVER_TYPE_LOGIC] = define.SERVER_RPC_PORT_START_LOGIC

	mgr.ServerStartPort = make(map[pb.SERVER_TYPE]uint32)
	mgr.ServerStartPort[pb.SERVER_TYPE_LOGIN] = define.SERVER_PORT_START_LOGIN
	mgr.ServerStartPort[pb.SERVER_TYPE_GATE] = define.SERVER_PORT_START_GATE
}

func (mgr *PortMgr) OnDestroy() {

}

func (mgr *PortMgr) Run() {

}

func (mgr *PortMgr) GetModulePort(serverType pb.SERVER_TYPE) (uint32, uint32, error) {
	var rpcPort uint32 = 0
	var port uint32 = 0
	err := errors.New("get module port fail")

	if serverType == pb.SERVER_TYPE_CENTER {
		return define.CENTER_RPC_PORT, port, nil
	}

	mgr.mutex.Lock()
	for i := 0; i < 1000; i++ {
		rpcPort = mgr.RpcServerStartPort[serverType] + uint32(i)
		if mgr.CheckPort(rpcPort) {
			err = nil
			break
		}
	}

	if serverType == pb.SERVER_TYPE_LOGIN || serverType == pb.SERVER_TYPE_GATE {
		for i := 0; i < 1000; i++ {
			port = mgr.ServerStartPort[serverType] + uint32(i)
			if mgr.CheckPort(rpcPort) {
				err = nil
				break
			}
		}
	}
	mgr.mutex.Unlock()

	return rpcPort, port, err
}

func (mgr *PortMgr) CheckPort(port uint32) bool {
	conn, err := net.Listen("tcp", ":"+strconv.FormatUint(uint64(port), 10))
	if err == nil {
		conn.Close()
		return true
	}

	log.Debug("%v  %v", err, port)

	return false
}
