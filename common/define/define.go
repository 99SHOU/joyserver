package define

const (
	MYSQL_DNS              = "root:123456@/arcs"
	GATE_RANDOM_STRING_LEN = 16
	GATE_TOKEN_EXPIRY_TIME = 60
)

const (
	//should write to conf file start
	MACHINE_RPC_IP_ADDR = "127.0.0.1"
	CENTER_RPC_IP_ADDR  = "127.0.0.1"
	//should write to conf file end

	//machine define start
	MACHINE_MODULE_ID = 1
	MACHINE_RPC_PORT  = 20000
	//machine define end

	//center define start
	CENTER_MODULE_ID = 2
	CENTER_RPC_PORT  = 21000
	//center define end

	//rpc port define start
	SERVER_RPC_PORT_START_LOGIN = 22000
	SERVER_PORT_START_LOGIN     = 23000
	SERVER_RPC_PORT_START_GATE  = 24000
	SERVER_PORT_START_GATE      = 25000
	SERVER_RPC_PORT_START_LOGIC = 26000
	//rpc port define end
)

// type SERVER_TYPE int

// const (
// 	SERVER_TYPE_START SERVER_TYPE = 1 + iota

// 	SERVER_TYPE_MACHINE
// 	SERVER_TYPE_LOGIN
// 	SERVER_TYPE_GATE
// 	SERVER_TYPE_CENTER
// 	SERVER_TYPE_LOGIC

// 	SERVER_TYPE_END
// )

// var serverType = [...]string{
// 	"SERVER_TYPE_START",

// 	"SERVER_TYPE_MACHINE",
// 	"SERVER_TYPE_LOGIN",
// 	"SERVER_TYPE_GATE",
// 	"SERVER_TYPE_CENTER",
// 	"SERVER_TYPE_LOGIC",

// 	"SERVER_TYPE_END",
// }

// func (st SERVER_TYPE) String() string {
// 	if SERVER_TYPE_START <= st && st <= SERVER_TYPE_END {
// 		return serverType[st-1]
// 	}

// 	return "Invaild ServerType"
// }

type COMMON_RESPOND_CODE int

const (
	COMMON_RESPOND_CODE_SUCCESS COMMON_RESPOND_CODE = 1 + iota
	COMMON_RESPOND_CODE_FAIL
)

type REGISTER_RESPOND_CODE int

const (
	REGISTER_RESPOND_CODE_SUCCESS REGISTER_RESPOND_CODE = 1 + iota
	REGISTER_RESPOND_CODE_FAIL
)

type SERVER_STATU int

// const (
// 	SERVER_STATU_INVALUE        SERVER_STATU = 1 + iota // invalue statu
// 	SERVER_STATU_REFUSE_SERVICE                         // not ready to work
// 	SERVER_STATU_START_SERVICE                          // ready to work
// 	SERVER_STATU_STOP                                   // ready to stop
// )

type VERIFY_ACCOUNT_RESPOND_CODE int

const (
	VERIFY_ACCOUNT_RESPOND_CODE_SUCCESS VERIFY_ACCOUNT_RESPOND_CODE = 1 + iota
	VERIFY_ACCOUNT_RESPOND_CODE_FAIL
)

type GET_GATE_ADDR_RESPOND_CODE int

const (
	GET_GATE_ADDR_RESPOND_CODE_SUCCESS GET_GATE_ADDR_RESPOND_CODE = 1 + iota
	GET_GATE_ADDR_RESPOND_CODE_FAIL
)

type GATE_TOKEN_RESPOND_CODE int

const (
	GATE_TOKEN_RESPOND_CODE_SUCCESS GATE_TOKEN_RESPOND_CODE = 1 + iota
	GATE_TOKEN_RESPOND_CODE_FAIL
)
