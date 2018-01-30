package modules

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
)

type AgentManager struct {
	agents map[string]net.Agent
}

func (am *AgentManager) Init() {
	am.agents = make(map[string]net.Agent)
}

func (am *AgentManager) Destroy() {

}

func (am *AgentManager) Run() {

}

func (am *AgentManager) AddAgent(key string, agent net.Agent) {
	if agent, ok := am.agents[key]; ok {
		log.Error("Agent is exist, Key: %v NodeType: %v", key, agent.GetNodeType().String())
		return
	}

	am.agents[key] = agent
}

func (am *AgentManager) RemoveAgent(key string) {
	delete(am.agents, key)
}

func (am *AgentManager) GetAgentAll() []net.Agent {
	agents := []net.Agent{}

	for _, agent := range am.agents {
		if agent.GetNodeType() != pb.NODE_TYPE_INVALID {
			agents = append(agents, agent)
		}
	}
	return agents
}

func (am *AgentManager) GetAgentByKey(key string) net.Agent {
	if agent, ok := am.agents[key]; ok && agent.GetNodeStatu() == pb.NODE_STATU_READY {
		return agent
	}

	return nil
}

func (am *AgentManager) GetAgentByNodeType(nodeTypes []pb.NODE_TYPE) []net.Agent {
	agents := []net.Agent{}

	for _, agent := range am.agents {
		for _, nt := range nodeTypes {
			if agent.GetNodeType() == nt && agent.GetNodeStatu() == pb.NODE_STATU_READY {
				agents = append(agents, agent)
			}
		}
	}

	return agents
}

func (am *AgentManager) GetAgentByNodeID(nodeID define.NodeID) []net.Agent {
	agents := []net.Agent{}

	for _, agent := range am.agents {
		if agent.GetNodeID() == nodeID && agent.GetNodeStatu() == pb.NODE_STATU_READY {
			agents = append(agents, agent)
		}
	}

	return agents
}

func (am *AgentManager) GetAgentByNodeInfo(infoKey interface{}, infoValue interface{}) []net.Agent {
	agents := []net.Agent{}

	for _, agent := range am.agents {
		if agent.GetAgentInfo(infoKey) == infoValue && agent.GetNodeStatu() == pb.NODE_STATU_READY {
			agents = append(agents, agent)
		}
	}

	return agents
}
