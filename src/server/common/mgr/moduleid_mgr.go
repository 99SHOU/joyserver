package mgr

import (
	"sync"
)

type ModuleIdMgr struct {
	mutex   sync.Mutex
	StartId int
}

func (mgr *ModuleIdMgr) GetModuleId() int {
	moduleId := 0

	mgr.mutex.Lock()
	mgr.StartId = mgr.StartId + 1
	moduleId = mgr.StartId
	mgr.mutex.Unlock()

	return moduleId
}
