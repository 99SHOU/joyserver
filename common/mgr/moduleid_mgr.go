package mgr

import (
	"sync"
)

type ModuleIdMgr struct {
	mutex   sync.Mutex
	StartId uint32
}

func (mgr *ModuleIdMgr) GetModuleId() uint32 {
	var moduleId uint32 = 0

	mgr.mutex.Lock()
	mgr.StartId = mgr.StartId + 1
	moduleId = mgr.StartId
	mgr.mutex.Unlock()

	return moduleId
}
