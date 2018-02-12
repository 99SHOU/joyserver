package modules

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
)

/*
TODO:maybe this module need a lock
*/

type AgentManager struct {
	BaseModule
	agents map[string]net.Agent
}

func (am *AgentManager) Init() {
	am.agents = make(map[string]net.Agent)
	log.Debug("init")
}

func (am *AgentManager) AfterInit() {
	log.Debug("afterinit")
}

func (am *AgentManager) BeforeDestroy() {
	log.Debug("beforedestroy")
}

func (am *AgentManager) Destroy() {
	log.Debug("destroy")
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

func (am *AgentManager) GetAgentByKey(key string) net.Agent {
	if agent, ok := am.agents[key]; ok {
		return agent
	}

	return nil
}

func (am *AgentManager) GetAgentByNodeID(nodeID define.NodeID) net.Agent {
	for _, agent := range am.agents {
		if agent.GetNodeID() == nodeID && agent.GetNodeStatu() == pb.NODE_STATU_READY {
			return agent
		}
	}

	return nil
}

func (am *AgentManager) GetAgentAll(f func(net.Agent) bool) []net.Agent {
	agents := []net.Agent{}

	for _, agent := range am.agents {
		if agent.GetNodeStatu() == pb.NODE_STATU_READY {
			if f == nil || f(agent) {
				agents = append(agents, agent)
			}
		}
	}

	return agents
}

func (am *AgentManager) GetAgentByNodeType(nodeTypes []pb.NODE_TYPE, f func(net.Agent) bool) []net.Agent {
	agents := []net.Agent{}

	for _, agent := range am.agents {
		for _, nt := range nodeTypes {
			if agent.GetNodeType() == nt && agent.GetNodeStatu() == pb.NODE_STATU_READY {
				if f == nil || f(agent) {
					agents = append(agents, agent)
				}
			}
		}
	}

	return agents
}

func (am *AgentManager) GetAgentByNodeInfo(infoKey interface{}, infoValue interface{}, f func(net.Agent) bool) []net.Agent {
	agents := []net.Agent{}

	for _, agent := range am.agents {
		if agent.GetAgentInfo(infoKey) == infoValue && agent.GetNodeStatu() == pb.NODE_STATU_READY {
			if f == nil || f(agent) {
				agents = append(agents, agent)
			}
		}
	}

	return agents
}
