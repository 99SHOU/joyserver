package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/base"
)

type AgentMgr struct {
	agentList map[string]*base.Agent
}

func NewAgentMgr() *AgentMgr {
	agentMgr := &AgentMgr{}
	agentMgr.agentList = make(map[string]*base.Agent)

	return agentMgr
}

func (am *AgentMgr) AddAgent(account string, agent *base.Agent) bool {
	_, ok := am.agentList[account]
	if ok {
		log.Error("Account %v is exist a agent", account)
		return false
	}

	am.agentList[account] = agent

	return true
}

func (am *AgentMgr) RemoveAgent(account string) bool {
	_, ok := am.agentList[account]
	if !ok {
		log.Error("Account %v can not find agent", account)
		return false
	}

	delete(am.agentList, account)

	return true
}

func (am *AgentMgr) GetAgentCount() int {
	return len(am.agentList)
}
