package mysql

import (
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/name5566/leaf/log"
)

func Open(dns string) *sql.DB {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		//panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		log.Fatal(err.Error())
	}

	stmtOut, err := db.Prepare("SELECT account FROM account WHERE account = ?")
	if err != nil {
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query("hello")
	if err == nil {
		log.Debug("%v", rows.Next())
	}

	return db
}
