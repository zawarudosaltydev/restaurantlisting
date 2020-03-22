package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db, err := sql.Open("mysql", "user:password@/dbname")

	if err != nil {
		panic(err.Error())
	}
}
