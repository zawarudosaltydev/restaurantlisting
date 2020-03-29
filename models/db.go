package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	db.SetMaxOpenConns(1000)

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
