package sission_module

func NewAgentInfo() *AgentInfo {
	agentInfo := new(AgentInfo)
	agentInfo.userData = make(map[interface{}]interface{})

	return agentInfo
}

type SessionInfo struct {
	NodeType  pb.NODE_TYPE
	NodeStatu pb.NODE_STATU
	NodeID    define.NodeID
	userData  map[interface{}]interface{}
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
