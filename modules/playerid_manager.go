package modules

// import (
// 	"github.com/99SHOU/joyserver/common/define"
// 	"github.com/99SHOU/joyserver/common/pb"
// 	"github.com/name5566/leaf/log"
// 	"gopkg.in/mgo.v2/bson"
// 	"sync"
// )

// type PlayerIdManager struct {
// 	BaseModule
// 	dbmanager *DBManager
// 	mutex     sync.Mutex
// }

// func (pm *PlayerIdManager) Init() {

// }

// func (pm *PlayerIdManager) AfterInit() {
// 	/* suggest to get module in afterinit if get module
// 	   in init not promise the target module is init
// 	*/
// 	pm.dbmanager = pm.mm.Find("DBManager").(*DBManager)

// 	pm.CheckDatabase()
// }

// func (pm *PlayerIdManager) BeforeDestroy() {

// }

// func (pm *PlayerIdManager) Destroy() {

// }

// func (pm *PlayerIdManager) Run() {

// }

// func (pm *PlayerIdManager) CheckDatabase() {
// 	count, err := pm.dbmanager.DB.C("PlayerIdPool").Find(nil).Count()
// 	if err != nil || count != 1 {
// 		log.Fatal("PlayidPool Error maybe you need to create collection PlayerIdPool manual: %v", err.Error())
// 		return
// 	}

// 	playeridPool := &pb.PlayerIdPool{}
// 	err = pm.dbmanager.DB.C("PlayerIdPool").Find(nil).One(playeridPool)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 		return
// 	}

// 	if playeridPool.CurMaxId <= 0 {
// 		log.Fatal("CurMaxId must greater then zero")
// 	}
// }

// func (pm *PlayerIdManager) GeneratePlayerId() (define.PlayerID, error) {
// 	pm.mutex.Lock()
// 	defer pm.mutex.Unlock()
// 	playeridPool := &pb.PlayerIdPool{}
// 	err := pm.dbmanager.DB.C("PlayerIdPool").Find(nil).One(playeridPool)
// 	if err != nil {
// 		return 0, err
// 	}

// 	playerid := define.PlayerID(playeridPool.CurMaxId)

// 	selector := bson.M{"curmaxid": playeridPool.CurMaxId}
// 	data := bson.M{"$set": bson.M{"curmaxid": playeridPool.CurMaxId + 1}}
// 	err = pm.dbmanager.DB.C("PlayerIdPool").Update(selector, data)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return playerid, nil
// }
