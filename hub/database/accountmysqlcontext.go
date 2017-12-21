package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func NewAccountMySQLContext(db *sql.DB) AccountMySQLContext {
	context := AccountMySQLContext{db}
	return context
}

type AccountMySQLContext struct {
	db *sql.DB
}

func (this *AccountMySQLContext) register(username string, password string, long float64, lat float64) bool {
	stm, err := this.db.Prepare("INSERT INTO User(username, password, homeLongitude, homeLatitude) VALUES(?,?,?,?)")
	defer stm.Close()
	if (err != nil) {
		panic(err.Error())
	}
	res, err := stm.Exec(username, password, long, lat)
	if (err!= nil || res == nil) {
		return false
	}
	return true
}

func (this *AccountMySQLContext) login(username string, password string) bool {
	rows, err := this.db.Query("SELECT * FROM User WHERE username = ? AND password = ?",
	username, password)
	defer rows.Close()
	if (err != nil) {
		panic(err.Error())
	}
	if (rows.Next()) {
		return true
	}
	return false
}
