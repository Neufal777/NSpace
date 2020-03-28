package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

//const of DB connection information
const (
	host     = "HOST"
	port     = 8080
	dbuser   = "DBUSER"
	password = "PASS"
	dbname   = "DBNAME"
)

type conn struct {
	host   string
	port   int
	dbuser string
	pass   string
	dbname string
}

//type identifying the user fields
type user struct {
	Id       int32
	Name     string
	Surname  string
	Nickname string
	balance  float32
}

func runHomePage(w http.ResponseWriter, r *http.Request) {

	user := user{

		Id:       1,
		Name:     "Naoufal",
		Surname:  "Dahouli",
		Nickname: "neufal79",
		balance:  342,
	}

	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, user)

}

func registerUser(name string, surname string, username string, balance int) {

	//initialize the database [REFACTOR]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbuser, password, dbname)

	//open the database
	db, err := sql.Open("postgres", psqlInfo)

	//insert in tthe database the user information
	insertStatement := `INSERT INTO users (name, surname, nickname, balance) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(insertStatement, name, surname, username, balance)

	if err != nil {
		panic(err)
	}

	//Success message
	log.Println("Registered succesfully")
}

func showUsers() {

	u := user{}

	//information message
	log.Println("Connecting to SQL...")

	//initialize the database [REFACTOR]
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbuser, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	//close the database at the end
	defer db.Close()

	//success mesage display [information]
	fmt.Println("Successfully connected!")

	//SQL to get values from the database
	rows, _ := db.Query("SELECT id, name, surname, nickname, balance FROM users")

	//Display all the records from the database
	for rows.Next() {
		rows.Scan(&u.Id, &u.Name, &u.Surname, &u.Nickname, &u.balance)
		log.Printf("|%v|%v|%v|%v|%v  ", u.Id, u.Name, u.Surname, u.Nickname, u.balance)
	}
}

func main() {

	//Introduce value in the database
	registerUser("NAME", "SURNAME", "USERNAME", 0)

	//show results
	showUsers()
}
