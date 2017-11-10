package main

import (
	"database/sql"
	"log"
	"os/user"

	_ "github.com/mattn/go-sqlite3"
)

func getMapsKey() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get current user: %v", err)
	}
	db, err := sql.Open("sqlite3", usr.HomeDir+"/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT value FROM settings WHERE name = 'mapskey'")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var value string
	err = stmt.QueryRow().Scan(&value)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
