package base

import (
	"github.com/name5566/leaf/log"
	"net"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/mgr"
	"github.com/99SHOU/joyserver/common/rpc_struct"
	"strconv"
	"strings"
)

type Module struct {
	ServerStatu    define.SERVER_STATU
	ServerType     define.SERVER_TYPE
	ModuleId       int
	Port           int
	MachineModueId int
	CenterModuleId int

	RpcMgr     *mgr.RpcMgr
	RpcHandler interface{}
}

func (m *Module) CreateRpcClientToMachine() {
	machineModule := m.RpcMgr.NewModuleClient(define.MACHINE_IP_ADDR+":"+strconv.Itoa(define.MACHINE_PORT), define.MACHINE_MODULE_ID, define.SERVER_TYPE_MACHINE)
	if machineModule != nil {
		err := m.RpcMgr.AddMoudleClient(machineModule.ModuleId, machineModule)
		if err == nil {
			m.MachineModueId = machineModule.ModuleId
		} else {
			log.Error("%v", err)
		}
	} else {
		log.Error("Connect to machine error")
	}
}

func (m *Module) CreateRpcClientToCenter() {
	centerModule := m.RpcMgr.NewModuleClient(define.CENTER_IP_ADDR+":"+strconv.Itoa(define.CENTER_PORT), define.CENTER_MODULE_ID, define.SERVER_TYPE_CENTER)
	if centerModule != nil {
		err := m.RpcMgr.AddMoudleClient(centerModule.ModuleId, centerModule)
		if err == nil {
			m.CenterModuleId = centerModule.ModuleId
		} else {
			log.Error("%v", err)
		}
	} else {
		log.Error("Connect to machine error")
	}
}

func (m *Module) RegisterToCenter() {
	localAddr := ""

	conn, err := net.Dial("tcp", define.CENTER_IP_ADDR+":"+strconv.Itoa(define.CENTER_PORT))
	if err != nil {
		log.Error("%v", err.Error())
		return
	}
	localAddr = strings.Split(conn.LocalAddr().String(), ":")[0] + ":" + strconv.Itoa(m.Port)
	defer conn.Close()

	req := &rpc_struct.RegisterToCenterReq{
		ServerType: m.ServerType,
		ServerAddr: localAddr,
		ModuleId:   m.ModuleId,
	}
	resp := new(rpc_struct.RegisterToCenterResp)
	err = m.RpcMgr.RpcCallByModuleId(m.CenterModuleId, "RpcHandler.RegisterToCenter", req, resp)
	if err != nil {
		log.Error("RpcCallByModuleId error: %v", err)
	}
	if resp.ResponCode == define.REGISTER_RESPOND_CODE_SUCCESS {
		log.Debug("Register to Center SUCCESS")
	} else {
		log.Error("Register to Center FAIL")
	}
}

func (m *Module) StartRpcServer(rpcHandler interface{}) {
	m.RpcMgr.StartRpcServer(rpcHandler, "127.0.0.1"+":"+strconv.Itoa(m.Port), m.ModuleId)
}

func (m *Module) ModuleIdReq() {
	req := &rpc_struct.ModuleIdReq{
		ServerType: m.ServerType,
	}
	resp := new(rpc_struct.ModuleIdResp)
	m.RpcMgr.RpcCallByModuleId(m.CenterModuleId, "RpcHandler.ModuleIdReq", req, resp)

	if resp.ModuleId != 0 {
		m.ModuleId = resp.ModuleId
	}
}

func (m *Module) ModulePortReq() {
	req := &rpc_struct.ModulePortReq{
		ServerType: m.ServerType,
	}
	resp := new(rpc_struct.ModulPortResp)
	m.RpcMgr.RpcCallByModuleId(m.MachineModueId, "RpcHandler.ModulePortReq", req, resp)

	if resp.Port != 0 {
		m.Port = resp.Port
	}
}
