package modules

// import (
// 	"errors"
// 	"fmt"
// 	"github.com/99SHOU/joyserver/common/define"
// 	"github.com/99SHOU/joyserver/common/pb"
// 	"github.com/name5566/leaf/log"
// 	"gopkg.in/mgo.v2/bson"
// 	"time"
// )

// type LoginManager struct {
// 	BaseModule
// 	dbmanager       *DBManager
// 	playeridManager *PlayerIdManager
// }

// func (lm *LoginManager) Init() {

// }

// func (lm *LoginManager) AfterInit() {
// 	/* suggest to get module in afterinit if get module
// 	   in init not promise the target module is init
// 	*/
// 	lm.dbmanager = lm.mm.Find("DBManager").(*DBManager)
// 	lm.playeridManager = lm.mm.Find("PlayerIdManager").(*PlayerIdManager)
// }

// func (lm *LoginManager) BeforeDestroy() {

// }

// func (lm *LoginManager) Destroy() {

// }

// func (lm *LoginManager) Run() {

// }

// func (lm *LoginManager) AccountVerify(account string) bool {
// 	count, err := lm.dbmanager.DB.C("Account").Find(bson.M{"account": account}).Count()
// 	if err != nil {
// 		log.Error(err.Error())
// 		return false
// 	}

// 	switch count {
// 	case 0:
// 		doc := &pb.Account{Account: account, RegisterTime: time.Now().Unix()}
// 		lm.dbmanager.DB.C("Account").Insert(doc)
// 		if err != nil {
// 			log.Error(err.Error())
// 			return false
// 		}
// 		return true
// 	case 1:
// 		return true
// 	default:
// 		log.Error("Too many document when account = %v in collection", account)
// 		return false
// 	}
// }

// func (lm *LoginManager) GetPlayerIdList(account string, gameid define.GameID) ([]define.PlayerID, error) {
// 	playeridList := []define.PlayerID{}

// 	result := pb.Account{}
// 	err := lm.dbmanager.DB.C("Account").Find(bson.M{"account": account}).One(result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, playeridInfo := range result.PlayeridList {
// 		if playeridInfo.GameId == int64(gameid) {
// 			playeridList = append(playeridList, define.PlayerID(playeridInfo.Playerid))
// 		}
// 	}

// 	return playeridList, nil
// }

// func (lm *LoginManager) CreatePlayer(account string, gameid define.GameID) (define.PlayerID, error) {
// 	selector := bson.M{"account": account}
// 	result := pb.Account{}
// 	err := lm.dbmanager.DB.C("Account").Find(selector).One(result)
// 	if err != nil {
// 		return 0, err
// 	}

// 	for _, playerid := range result.PlayeridList {
// 		if playerid.GameId == int64(gameid) && int64(playerid.Playerid) > 0 {
// 			return 0, errors.New(fmt.Sprintf("Exist a player gameid %v playerid %v", playerid.GameId, playerid.Playerid))
// 		}
// 	}

// 	playerid, err := lm.playeridManager.GeneratePlayerId()
// 	if err != nil {
// 		return 0, errors.New(fmt.Sprintf("Generate playerid error"))
// 	}

// 	result.PlayeridList = append(result.PlayeridList, &pb.PlayerId{GameId: int64(gameid), Playerid: int64(playerid)})

// 	_, err = lm.dbmanager.DB.C("Account").Upsert(selector, result)
// 	if err != nil {
// 		return 0, errors.New(fmt.Sprintf("Add playerid error"))
// 	}

// 	return playerid, nil
// }
