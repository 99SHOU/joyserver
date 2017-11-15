package mgr

import (
	"errors"
	"github.com/name5566/leaf/log"
	"net"
	"net/http"
	"net/rpc"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/module_client"
	"strconv"
)

type RpcMgr struct {
	ModuleClient map[int]*module_client.ModuleClient
	ServerAddr   string
	ServerType   define.SERVER_TYPE
}

func (mgr *RpcMgr) OnInit() {

}

func (mgr *RpcMgr) OnDestroy() {

}

func (mgr *RpcMgr) Run() {

}

func (mgr *RpcMgr) NewModuleClient(serverAddr string, moduleId int, serverType define.SERVER_TYPE) *module_client.ModuleClient {
	//rpcClient, err := rpc.DialHTTP("tcp", serverAddr)
	rpcClient, err := rpc.DialHTTPPath("tcp", serverAddr, rpc.DefaultRPCPath+strconv.Itoa(moduleId))
	if err != nil {
		log.Error("NewModuleClient ERROR: can not Dial to server %v, %v", serverAddr, err)
	} else {
		moduleClient := &module_client.ModuleClient{
			RpcClient:  rpcClient,
			ModuleId:   moduleId,
			ServerAddr: serverAddr,
			ServerType: serverType,
		}

		return moduleClient
	}

	return nil
}

func (mgr *RpcMgr) StartRpcServer(rcvr interface{}, serverAddr string, moduleId int) {
	mgr.ServerAddr = serverAddr
	server := rpc.NewServer()
	server.Register(rcvr)
	server.HandleHTTP(rpc.DefaultRPCPath+strconv.Itoa(moduleId), rpc.DefaultDebugPath+strconv.Itoa(moduleId))

	log.Debug("Start Server %v, ModuleId:%v, ServerType:%v", serverAddr, moduleId, mgr.ServerType)

	l, e := net.Listen("tcp", serverAddr)
	if e != nil {
		log.Fatal("listen error:", e)
		return
	}
	go http.Serve(l, nil)
}

func (mgr *RpcMgr) AddMoudleClient(moduleId int, moduleClient *module_client.ModuleClient) error {
	_, ok := mgr.ModuleClient[moduleId]
	if ok {
		return errors.New("ModuleClient is exist:" + mgr.ModuleClient[moduleId].String())
	}

	mgr.ModuleClient[moduleId] = moduleClient
	return nil
}

func (mgr *RpcMgr) RemoveMoudleClient(moduleId int) error {
	_, ok := mgr.ModuleClient[moduleId]
	if !ok {
		return errors.New("ModuleClient is not exist:" + mgr.ModuleClient[moduleId].String())
	}

	mgr.ModuleClient[moduleId] = nil
	return nil
}

func (mgr *RpcMgr) RpcCallByModuleId(moduleId int, method string, req interface{}, resp interface{}) error {
	moduleClient, ok := mgr.ModuleClient[moduleId]
	if !ok {
		log.Error("can not find Module id = %v", moduleId)
		return errors.New("can not find Module id = " + strconv.Itoa(moduleId))
	}

	if moduleClient.RpcClient == nil {
		log.Error("RpcClient is nil, ModuleId = %v", moduleId)
		return errors.New("RpcClient is nil, ModuleId = " + strconv.Itoa(moduleId))
	}

	return moduleClient.RpcClient.Call(method, req, resp)
}

func (mgr *RpcMgr) RpcCallByServerType(serverType define.SERVER_TYPE, method string, req interface{}, resp interface{}) {
	for moduleId, moduleClient := range mgr.ModuleClient {
		if moduleClient.ServerType == serverType {
			err := mgr.RpcCallByModuleId(moduleId, method, req, resp)
			if err != nil {
				log.Error("RpcCallByModuleId error: %v", err)
			}
		}
	}
}

func (mgr *RpcMgr) RpcCallByServerTypeExcept(serverType define.SERVER_TYPE, exceptModuleId int, method string, req interface{}, resp interface{}) {
	for moduleId, moduleClient := range mgr.ModuleClient {
		if moduleClient.ServerType == serverType && moduleClient.ModuleId != exceptModuleId {
			err := mgr.RpcCallByModuleId(moduleId, method, req, resp)
			if err != nil {
				log.Error("RpcCallByModuleId error: %v", err)
			}
		}
	}
}

func (mgr *RpcMgr) GetModuleClientByModuleId(moduleId int) *module_client.ModuleClient {
	moduleClient, ok := mgr.ModuleClient[moduleId]
	if !ok {
		log.Error("can not find ModuleClient moduleId = %v", moduleId)
		return nil
	}

	return moduleClient
}

// func (mgr *RpcMgr) GetModuleClientByServerType(serverType define.SERVER_TYPE) map[int]*module_client.ModuleClient {
// 	moduleClientList := make(map[int]*module_client.ModuleClient)
// 	for moduleId, moduleClient := range mgr.ModuleClient {
// 		if moduleClient.ServerType == serverType {
// 			moduleClientList[moduleId] = moduleClient
// 		}
// 	}

// 	return moduleClientList
// }

func (mgr *RpcMgr) GetModuleClientByServerType(serverType define.SERVER_TYPE) []*module_client.ModuleClient {
	moduleClientList := []*module_client.ModuleClient{}
	for _, moduleClient := range mgr.ModuleClient {
		if moduleClient.ServerType == serverType {
			moduleClientList = append(moduleClientList, moduleClient)
		}
	}

	return moduleClientList
}
