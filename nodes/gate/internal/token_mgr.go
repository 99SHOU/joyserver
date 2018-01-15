package internal

import (
	"github.com/name5566/leaf/log"
)

type TokenMgr struct {
	tokenList map[string]string
}

func NewTokenMgr() *TokenMgr {
	tokenMgr := &TokenMgr{}
	tokenMgr.tokenList = make(map[string]string)

	return tokenMgr
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
