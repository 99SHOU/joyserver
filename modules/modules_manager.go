package modules

import (
	"github.com/name5566/leaf/log"
	"reflect"
)

type Module interface {
	SetModulesManager(mm *ModulesManager)
	Init()
	AfterInit()
	BeforeDestroy()
	Destroy()
	Run()
}

type BaseModule struct {
	mm *ModulesManager
}

func (bm *BaseModule) SetModulesManager(mm *ModulesManager) {
	bm.mm = mm
}

func (bm *BaseModule) Init() {

}

func (bm *BaseModule) AfterInit() {

}

func (bm *BaseModule) BeforeDestroy() {

}

func (bm *BaseModule) Destroy() {

}

func (bm *BaseModule) Run() {

}

func NewModulesManager() ModulesManager {
	mm := ModulesManager{modules: make(map[string]Module)}
	return mm
}

type ModulesManager struct {
	modules map[string]Module
}

func (mm *ModulesManager) Register(module Module) {
	key := reflect.TypeOf(module).Elem().Name()
	if _, ok := mm.modules[key]; ok {
		log.Error("Module %v is already register", key)
	}

	module.SetModulesManager(mm)
	mm.modules[key] = module
}

func (mm *ModulesManager) Find(moduleName string) Module {
	module, ok := mm.modules[moduleName]
	if !ok {
		log.Error("Can not find module %v", moduleName)
		return nil
	}

	return module
}

func (mm *ModulesManager) Init() {
	for _, module := range mm.modules {
		module.Init()
	}
}

func (mm *ModulesManager) AfterInit() {
	for _, module := range mm.modules {
		module.AfterInit()
	}
}

func (mm *ModulesManager) BeforeDestroy() {
	for _, module := range mm.modules {
		module.BeforeDestroy()
	}
}

func (mm *ModulesManager) Destroy() {
	for _, module := range mm.modules {
		module.Destroy()
	}
}

func (mm *ModulesManager) Run() {
	for _, module := range mm.modules {
		module.Run()
	}
}
