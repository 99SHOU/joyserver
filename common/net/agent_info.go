package net

import (
	"github.com/99SHOU/joyserver/common/pb"
)

func NewAgentInfo() *AgentInfo {
	agentInfo := new(AgentInfo)
	agentInfo.userData = make(map[interface{}]interface{})

	return agentInfo
}

type AgentInfo struct {
	NodeType pb.NODE_TYPE
	userData map[interface{}]interface{}
}

func (info *AgentInfo) GetUserData(key interface{}) interface{} {
	if value, ok := info.userData[key]; ok {
		return value
	}

	return nil
}

func (info *AgentInfo) SetUserData(key interface{}, value interface{}) {
	info.userData[key] = value
}
