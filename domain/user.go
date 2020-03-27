package domain

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 8080
	dbuser   = "postgres"
	password = "root"
	dbname   = "postgres"
)

func registerUser(name string, surname string, username string, balance int) {

	//initialize the database [REFACTOR]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbuser, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	//insert in tthe database the user information
	insertStatement := `INSERT INTO users (name, surname, nickname, balance) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(insertStatement, name, surname, username, balance)

	if err != nil {
		panic(err)
	}

	log.Println("Registered succesfully")
}
