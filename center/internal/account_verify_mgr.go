package internal

import (
	"database/sql"
)

type AccountVerifyMgr struct {
	db *sql.DB
}

func (avm *AccountVerifyMgr) Init() {

}

func (avm *AccountVerifyMgr) Destroy() {
	avm.db.Close()
}

func (avm *AccountVerifyMgr) VarifyAccount(account string) (bool, error) {
	stmtOut, err := avm.db.Prepare("SELECT account FROM account WHERE account = ?")
	if err != nil {
		return false, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(account)
	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}
