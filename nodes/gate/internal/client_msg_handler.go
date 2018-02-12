package internal

// import (
// 	"github.com/99SHOU/joyserver/common/define"
// 	"github.com/99SHOU/joyserver/common/msg"
// 	"github.com/99SHOU/joyserver/common/net"
// 	"github.com/99SHOU/joyserver/common/pb"
// 	"github.com/name5566/leaf/log"
// 	// "strconv"
// 	// "strings"
// 	// "time"
// )

// type NodeServerHandler struct {
// 	Node      *Node
// 	processor *net.Processor
// }

// func (h *NodeServerHandler) Register(server *net.Server) {
// 	server.OnNewAgent = h.NewAgent
// 	server.OnCloseAgent = h.CloseAgent

// 	h.processor = server.Processor
// 	msg.RegisterMsg(server.Processor)

// 	server.SetOtherHandler(h.OnOtherMsgHandler)
// }

// func (h *NodeServerHandler) NewAgent(agent net.Agent) {
// }

// func (h *NodeServerHandler) CloseAgent(agent net.Agent) {
// }

// //all of msg gate do not handle will send to game
// func (h *NodeServerHandler) OnOtherMsgHandler(message interface{}, agent interface{}) {
// 	a := agent.(*net.BaseAgent)

// 	value := a.GetAgentInfo(pb.AGENT_INFO_KEY_CHARACTER_ID)
// 	if value == nil {
// 		log.Error("GetAgentInfo ERROR: value is nil, Key: %v", pb.AGENT_INFO_KEY_CHARACTER_ID.String())
// 		return
// 	}

// 	characterId := value.(define.CharacterID)
// 	if a.GetNodeType() == pb.NODE_TYPE_GAME {
// 		tempMsg, err := h.processor.Marshal(message)
// 		if err != nil {
// 			log.Error("Mashel Msg Error")
// 		}

// 		a.WriteMsg(&pb.GameMsgTransfer{CharacterId: uint64(characterId), MsgId: tempMsg[0], MsgBody: tempMsg[1]})
// 	}
// }
