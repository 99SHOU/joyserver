package define

const (
	//should write to conf file start
	MACHINE_IP_ADDR = "127.0.0.1"
	CENTER_IP_ADDR  = "127.0.0.1"
	//should write to conf file end

	//machine define start
	MACHINE_MODULE_ID = 1
	MACHINE_PORT      = 20000
	//machine define end

	//center define start
	CENTER_MODULE_ID = 2
	CENTER_PORT      = 21000
	//center define end

	//port define start
	SERVER_PORT_START_LOGIN = 22000
	SERVER_PORT_START_GATE  = 23000
	SERVER_PORT_START_LOGIC = 24000
	//port define end

	//machine port

)

const ()

type SERVER_TYPE int

const (
	SERVER_TYPE_START SERVER_TYPE = 1 + iota

	SERVER_TYPE_MACHINE
	SERVER_TYPE_LOGIN
	SERVER_TYPE_GATE
	SERVER_TYPE_CENTER
	SERVER_TYPE_LOGIC

	SERVER_TYPE_END
)

var serverType = [...]string{
	"SERVER_TYPE_START",

	"SERVER_TYPE_MACHINE",
	"SERVER_TYPE_LOGIN",
	"SERVER_TYPE_GATE",
	"SERVER_TYPE_CENTER",
	"SERVER_TYPE_LOGIC",

	"SERVER_TYPE_END",
}

func (st SERVER_TYPE) String() string {
	if SERVER_TYPE_START <= st && st <= SERVER_TYPE_END {
		return serverType[st-1]
	}

	return "Invaild ServerType"
}

type REGISTER_RESPOND_CODE int

const (
	REGISTER_RESPOND_CODE_START REGISTER_RESPOND_CODE = 1 + iota

	REGISTER_RESPOND_CODE_SUCCESS
	REGISTER_RESPOND_CODE_FAIL

	REGISTER_RESPOND_CODE_END
)

type SERVER_STATU int

const (
	SERVER_STATU_INVALUE        SERVER_STATU = 1 + iota // invalue statu
	SERVER_STATU_REFUSE_SERVICE                         // not ready to work
	SERVER_STATU_START_SERVICE                          // ready to work
	SERVER_STATU_STOP                                   // ready to stop
)
