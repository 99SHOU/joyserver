package modules

import (
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2"
)

type DBManager struct {
	BaseModule
	Session *mgo.Session
	DB      *mgo.Database
}

func (dbm *DBManager) Init() {
	dbm.open("")
}

func (dbm *DBManager) AfterInit() {

}

func (dbm *DBManager) BeforeDestroy() {

}

func (dbm *DBManager) Destroy() {
	if dbm.Session != nil {
		dbm.Session.Close()
	}
}

func (dbm *DBManager) Run() {

}

func (dbm *DBManager) open(url string) {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbm.Session = session
	dbm.DB = session.DB("taxespoker")
}
