package base

import (
	"github.com/name5566/leaf/log"
	"net"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/mgr"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/common/rpc_struct"
	"strconv"
	"strings"
)

type Module struct {
	ServerType     pb.SERVER_TYPE
	ModuleId       uint32
	RpcPort        uint32
	Port           uint32
	MachineModueId uint32
	CenterModuleId uint32

	RpcMgr        *mgr.RpcMgr
	RpcHandler    interface{}
	ServerHandler ServerHandler
}

func (m *Module) CreateRpcClientToMachine() {
	machineModule := m.RpcMgr.NewRpcClient(define.MACHINE_RPC_IP_ADDR+":"+strconv.Itoa(define.MACHINE_RPC_PORT), define.MACHINE_MODULE_ID, pb.SERVER_TYPE_MACHINE)
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
	centerModule := m.RpcMgr.NewRpcClient(define.CENTER_RPC_IP_ADDR+":"+strconv.Itoa(define.CENTER_RPC_PORT), define.CENTER_MODULE_ID, pb.SERVER_TYPE_CENTER)
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

func (m *Module) RegisterToCenter() bool {
	IP := ""

	//start get local IP
	conn, err := net.Dial("tcp", define.CENTER_RPC_IP_ADDR+":"+strconv.Itoa(define.CENTER_RPC_PORT))
	if err != nil {
		log.Error("%v", err.Error())
		return false
	}
	IP = strings.Split(conn.LocalAddr().String(), ":")[0]
	defer conn.Close()
	//end get local IP

	req := &rpc_struct.RegisterToCenterReq{
		ServerType: m.ServerType,
		ModuleId:   m.ModuleId,
		Ip:         IP,
		RpcPort:    m.RpcPort,
		Port:       m.Port,
	}
	resp := new(rpc_struct.RegisterToCenterResp)
	err = m.RpcMgr.RpcCallByModuleId(m.CenterModuleId, "RpcHandler.RegisterToCenter", req, resp)
	if err != nil {
		log.Error("RpcCallByModuleId error: %v", err)
	}
	if resp.ResponCode == define.REGISTER_RESPOND_CODE_SUCCESS {
		log.Debug("Register to Center SUCCESS")
		return true
	} else {
		log.Error("Register to Center FAIL sfsdf %v", resp.ResponCode)
	}

	return false
}

func (m *Module) StartRpcServer(rpcHandler interface{}) {
	m.RpcMgr.StartRpcServer(rpcHandler, "127.0.0.1"+":"+strconv.FormatUint(uint64(m.RpcPort), 10), m.ModuleId)
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

	if resp.RpcPort != 0 {
		m.RpcPort = resp.RpcPort
		m.Port = resp.Port
	}
}
