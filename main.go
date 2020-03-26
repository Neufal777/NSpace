//DEVELOP
package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "HOME")
}

func run_profile(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "PROFILE")

}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/profile", run_profile)
	http.ListenAndServe(":8000", nil)
}
