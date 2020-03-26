//DEVELOP
package main

import (
	"html/template"
	"net/http"
)

type notification struct {
	read   int32
	unread int32
}

type user struct {
	Name          string
	Surname       string
	Nickname      string
	notifications notification
	balance       float32
}

func runHomePage(w http.ResponseWriter, r *http.Request) {

	user := []user{
		{
			Name:     "Naoufal",
			Surname:  "Dahouli",
			Nickname: "neufal79",
			notifications: notification{
				read:   2,
				unread: 32,
			},
			balance: 4.8,
		},
		{
			Name:     "Marshall",
			Surname:  "Mathers",
			Nickname: "mm89",
			notifications: notification{
				read:   5,
				unread: 12,
			},
			balance: 5.7,
		},
	}

	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, user[1])

}

func main() {

	http.HandleFunc("/", runHomePage)
	http.ListenAndServe(":8000", nil)
}
