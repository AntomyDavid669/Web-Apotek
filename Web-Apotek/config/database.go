package config

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

  var DB *sql.DB

func ConnectDB(){
	db, err := sql.Open("postgres", "user=postgres dbname=Apotek sslmode=disable password=admin")
	if err != nil {
	panic(err)
	}

	log.Println("Connected Database")
	DB = db
}