package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type MySQLConnection struct {
	db *sql.DB
	err error
}

func (this *MySQLConnection) OpenConnection() {
	this.db, this.err = sql.Open("mysql", "root:rootDB123!@tcp(smarthomehubdatabase.cfq1p9brsdwm.eu-west-2.rds.amazonaws.com:3306)/smarthomehub")
	if (this.err != nil) {
		panic(this.err.Error())
	}
//	defer this.db.Close()
	this.err = this.db.Ping()
	if (this.err != nil) {
		panic(this.err.Error())
	}
}

func (this *MySQLConnection) GetConnection() *sql.DB {
	return this.db
}
