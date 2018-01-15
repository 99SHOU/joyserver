package base

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/mgr"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/common/rpc_struct"
	"github.com/name5566/leaf/log"
	"net"
	"strconv"
	"strings"
)

type Node struct {
	NodeType       pb.SERVER_TYPE
	NodeId         uint32
	RpcPort        uint32
	Port           uint32
	MachineModueId uint32
	CenterModuleId uint32

	RpcMgr        *mgr.RpcMgr
	RpcHandler    interface{}
	ServerHandler ServerHandler
}

func (n *Node) CreateRpcClientToMachine() {
	machineModule := n.RpcMgr.NewRpcClient(define.MACHINE_RPC_IP_ADDR+":"+strconv.Itoa(define.MACHINE_RPC_PORT), define.MACHINE_MODULE_ID, pb.SERVER_TYPE_MACHINE)
	if machineModule != nil {
		err := n.RpcMgr.AddMoudleClient(machineModule.ModuleId, machineModule)
		if err == nil {
			n.MachineModueId = machineModule.ModuleId
		} else {
			log.Error("%v", err)
		}
	} else {
		log.Error("Connect to machine error")
	}
}

func (n *Node) CreateRpcClientToCenter() {
	centerModule := n.RpcMgr.NewRpcClient(define.CENTER_RPC_IP_ADDR+":"+strconv.Itoa(define.CENTER_RPC_PORT), define.CENTER_MODULE_ID, pb.SERVER_TYPE_CENTER)
	if centerModule != nil {
		err := n.RpcMgr.AddMoudleClient(centerModule.ModuleId, centerModule)
		if err == nil {
			n.CenterModuleId = centerModule.ModuleId
		} else {
			log.Error("%v", err)
		}
	} else {
		log.Error("Connect to machine error")
	}
}

func (n *Node) RegisterToCenter() bool {
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
		ServerType: n.NodeType,
		ModuleId:   n.NodeId,
		Ip:         IP,
		RpcPort:    n.RpcPort,
		Port:       n.Port,
	}
	resp := new(rpc_struct.RegisterToCenterResp)
	err = n.RpcMgr.RpcCallByModuleId(n.CenterModuleId, "RpcHandler.RegisterToCenter", req, resp)
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

func (n *Node) StartRpcServer(rpcHandler interface{}) {
	n.RpcMgr.StartRpcServer(rpcHandler, "127.0.0.1"+":"+strconv.FormatUint(uint64(n.RpcPort), 10), n.NodeId)
}

func (n *Node) ModuleIdReq() {
	req := &rpc_struct.ModuleIdReq{
		ServerType: n.NodeType,
	}
	resp := new(rpc_struct.ModuleIdResp)
	n.RpcMgr.RpcCallByModuleId(n.CenterModuleId, "RpcHandler.ModuleIdReq", req, resp)

	if resp.ModuleId != 0 {
		n.NodeId = resp.ModuleId
	}
}

func (n *Node) ModulePortReq() {
	req := &rpc_struct.ModulePortReq{
		ServerType: n.NodeType,
	}
	resp := new(rpc_struct.ModulPortResp)
	n.RpcMgr.RpcCallByModuleId(n.MachineModueId, "RpcHandler.ModulePortReq", req, resp)

	if resp.RpcPort != 0 {
		n.RpcPort = resp.RpcPort
		n.Port = resp.Port
	}
}
