package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func InitDB(host, port, user, password, dbName string) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}
