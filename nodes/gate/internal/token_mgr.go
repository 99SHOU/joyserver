package internal

import (
	"github.com/99SHOU/joyserver/modules"
	"github.com/name5566/leaf/log"
)

type TokenMgr struct {
	modules.BaseModule
	tokenList map[string]string
}

func (tm *TokenMgr) Init() {
	tm.tokenList = make(map[string]string)
}

func (tm *TokenMgr) AfterInit() {

}

func (tm *TokenMgr) BeforeDestroy() {

}

func (tm *TokenMgr) Destroy() {

}

func (tm *TokenMgr) AddTokenInfo(account string, token string) bool {
	tm.tokenList[account] = token
	return true
}

func (tm *TokenMgr) RemoveTokenInfo(account string) bool {
	delete(tm.tokenList, account)
	return true
}

func (tm *TokenMgr) GetToken(account string) string {
	token, ok := tm.tokenList[account]
	if ok {
		return token
	} else {
		log.Error("Get Login Token Error")
	}

	return ""
}
