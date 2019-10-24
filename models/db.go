package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// Doc: https://godoc.org/github.com/lib/pq
const connStr = "postgres://atp:atp@202.204.62.32:5432/atp?sslmode=disable"

var Conn *sql.DB

func init() {
	var err error
	Conn, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
