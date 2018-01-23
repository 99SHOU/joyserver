package internal

import (
	"database/sql"
)

type AccountVerifyManager struct {
	db *sql.DB
}

func (avm *AccountVerifyManager) Init() {

}

func (avm *AccountVerifyManager) Destroy() {
	avm.db.Close()
}

func (avm *AccountVerifyManager) Run() {

}

func (avm *AccountVerifyManager) VarifyAccount(account string) (bool, error) {
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
