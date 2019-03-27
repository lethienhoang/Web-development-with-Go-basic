package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

var users []User

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get parameter
	params := mux.Vars(r)
	id := params["id"]

	for _, item := range users {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&User{})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Welcome ...")
}

func main() {

	// Create a handler which receives all incomming HTTP connections from browsers
	// An http.ResponseWriter which is where you write your text/html response to
	// An http.Request which contains all information about this HTTP request including things like the URL or header fields

	r := mux.NewRouter()

	users = append(users, User{
		ID:        "1",
		Age:       12,
		Firstname: "testing",
		Lastname:  "Abc",
	})

	users = append(users, User{
		ID:        "2",
		Age:       22,
		Firstname: "joe",
		Lastname:  "c",
	})
	// default routes
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", getUsers).Methods("GET")
	http.ListenAndServe(":8080", r)

}
