package mgr

import (
	"errors"
	"github.com/name5566/leaf/log"
	"net"
	"net/http"
	"net/rpc"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/common/rpc_client"
	"strconv"
)

type RpcMgr struct {
	RpcClient  map[uint32]*rpc_client.RpcClient
	ServerAddr string
	ServerType pb.SERVER_TYPE
}

func (mgr *RpcMgr) OnInit() {

}

func (mgr *RpcMgr) OnDestroy() {

}

func (mgr *RpcMgr) Run() {

}

func (mgr *RpcMgr) NewRpcClient(serverAddr string, moduleId uint32, serverType pb.SERVER_TYPE) *rpc_client.RpcClient {
	//rpcClient, err := rpc.DialHTTP("tcp", serverAddr)
	rpcClient, err := rpc.DialHTTPPath("tcp", serverAddr, rpc.DefaultRPCPath+strconv.FormatUint(uint64(moduleId), 10))
	if err != nil {
		log.Error("NewRpcClient ERROR: can not Dial to server %v, %v", serverAddr, err)
	} else {
		rpcClient := &rpc_client.RpcClient{
			RpcClient:  rpcClient,
			ModuleId:   moduleId,
			ServerAddr: serverAddr,
			ServerType: serverType,
		}

		return rpcClient
	}

	return nil
}

func (mgr *RpcMgr) StartRpcServer(rcvr interface{}, serverAddr string, moduleId uint32) {
	mgr.ServerAddr = serverAddr
	server := rpc.NewServer()
	server.Register(rcvr)
	server.HandleHTTP(rpc.DefaultRPCPath+strconv.FormatUint(uint64(moduleId), 10), rpc.DefaultDebugPath+strconv.FormatUint(uint64(moduleId), 10))

	log.Debug("Start Server %v, ModuleId:%v, ServerType:%v", serverAddr, moduleId, mgr.ServerType)

	l, e := net.Listen("tcp", serverAddr)
	if e != nil {
		log.Fatal("listen error:", e)
		return
	}
	go http.Serve(l, nil)
}

func (mgr *RpcMgr) AddMoudleClient(moduleId uint32, rpcClient *rpc_client.RpcClient) error {
	_, ok := mgr.RpcClient[moduleId]
	if ok {
		return errors.New("RpcClient is exist:" + mgr.RpcClient[moduleId].String())
	}

	mgr.RpcClient[moduleId] = rpcClient
	return nil
}

func (mgr *RpcMgr) RemoveMoudleClient(moduleId uint32) error {
	_, ok := mgr.RpcClient[moduleId]
	if !ok {
		return errors.New("RpcClient is not exist:" + mgr.RpcClient[moduleId].String())
	}

	mgr.RpcClient[moduleId] = nil
	return nil
}

func (mgr *RpcMgr) RpcCallByModuleId(moduleId uint32, method string, req interface{}, resp interface{}) error {
	rpcClient, ok := mgr.RpcClient[moduleId]
	if !ok {
		log.Error("can not find Module id = %v", moduleId)
		return errors.New("can not find Module id = " + strconv.FormatUint(uint64(moduleId), 10))
	}

	if rpcClient.RpcClient == nil {
		log.Error("RpcClient is nil, ModuleId = %v", moduleId)
		return errors.New("RpcClient is nil, ModuleId = " + strconv.FormatUint(uint64(moduleId), 10))
	}

	return rpcClient.RpcClient.Call(method, req, resp)
}

func (mgr *RpcMgr) RpcCallByServerType(serverType pb.SERVER_TYPE, method string, req interface{}, resp interface{}) {
	for moduleId, rpcClient := range mgr.RpcClient {
		if rpcClient.ServerType == serverType {
			err := mgr.RpcCallByModuleId(moduleId, method, req, resp)
			if err != nil {
				log.Error("RpcCallByModuleId error: %v", err)
			}
		}
	}
}

func (mgr *RpcMgr) RpcCallByServerTypeExcept(serverType pb.SERVER_TYPE, exceptModuleId uint32, method string, req interface{}, resp interface{}) {
	for moduleId, rpcClient := range mgr.RpcClient {
		if rpcClient.ServerType == serverType && rpcClient.ModuleId != exceptModuleId {
			err := mgr.RpcCallByModuleId(moduleId, method, req, resp)
			if err != nil {
				log.Error("RpcCallByModuleId error: %v", err)
			}
		}
	}
}

func (mgr *RpcMgr) GetRpcClientByModuleId(moduleId uint32) *rpc_client.RpcClient {
	rpcClient, ok := mgr.RpcClient[moduleId]
	if !ok {
		log.Error("can not find RpcClient moduleId = %v", moduleId)
		return nil
	}

	return rpcClient
}

func (mgr *RpcMgr) GetRpcClientByServerType(serverType pb.SERVER_TYPE) []*rpc_client.RpcClient {
	RpcClientList := []*rpc_client.RpcClient{}
	for _, rpcClient := range mgr.RpcClient {
		if rpcClient.ServerType == serverType {
			RpcClientList = append(RpcClientList, rpcClient)
		}
	}

	return RpcClientList
}
