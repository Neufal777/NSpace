//DEVELOP
package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/NSpace/domain"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 8080
	dbuser   = "postgres"
	password = "root"
	dbname   = "postgres"
)

type notification struct {
	read   int32
	unread int32
}

type conn struct {
	host   string
	port   int
	dbuser string
	pass   string
	dbname string
}

type user struct {
	Id       int32
	Name     string
	Surname  string
	Nickname string
	balance  float32
}

func runHomePage(w http.ResponseWriter, r *http.Request) {

	user := []user{
		{
			Id:       1,
			Name:     "Naoufal",
			Surname:  "Dahouli",
			Nickname: "neufal79",
			balance:  342,
		},
		{
			Id:       2,
			Name:     "Marshall",
			Surname:  "Mathers",
			Nickname: "mm89",
			balance:  8,
		},
	}

	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, user[1])

}

func main() {

	domain.registerUser("TEST", "TEST", "TEST", 0)
	u := user{}

	log.Println("Connecting to SQL...")

	//initialize the database [REFACTOR]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbuser, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected!")

	rows, _ := db.Query("SELECT id, name, surname, nickname, balance FROM users")

	for rows.Next() {
		rows.Scan(&u.Id, &u.Name, &u.Surname, &u.Nickname, &u.balance)
		log.Printf("|%v|%v|%v|%v|%v  ", u.Id, u.Name, u.Surname, u.Nickname, u.balance)
	}

	//http.HandleFunc("/", runHomePage)
	//http.ListenAndServe(":8000", nil)
}
