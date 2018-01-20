package modules

import (
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
)

type AgentManager struct {
	agents map[string]*net.Agent
}

func (am *AgentManager) AddAgent(key string, agent *net.Agent) {
	if agent, ok := am.agents[key]; ok {
		log.Error("Agent is exist, Key: %v NodeType: %v", key, agent.GetNodeType().String)
		return
	}

	am.agents[key] = agent
}

func (am *AgentManager) RemoveAgent(key string) {
	delete(am.agents, key)
}

func (am *AgentManager) GetAgentByKey(key string) *net.Agent {
	if agent, ok := am.agents[key]; ok {
		return agent
	}

	return nil
}

func (am *AgentManager) GetAgentByNodeType(nodeType pb.NODE_TYPE) []*net.Agent {
	agents := []net.Agent{}

	for _, v := range am.agents {
		if v.GetNodeType() == nodeType {
			append(agents, v)
		}
	}

	return agents
}

func (am *AgentManager) GetAgentByNodeInfo(infoKey interface{}, infoValue interface{}) []*net.Agent {
	agents := []net.Agent{}

	for _, v := range am.agents {
		if v.GetAgentInfo(infoKey) == infoValue {
			append(agents, v)
		}
	}

	return agents
}
