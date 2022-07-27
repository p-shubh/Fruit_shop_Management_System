package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "ppt"
	dbname   = "fruit_vendor"
)

var DB *sql.DB

func Connection_with_db() {
	db_dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", db_dsn) //postgres is the driver and db_dsn is the connector

	if err != nil {
		log.Fatal("Failed to open the DB connection", err)
	} else {
		fmt.Println("connected")
	}

}
