package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SignUp struct {
	Email    string
	UserName string
	Password string
}

func main() {
	tmpl := template.Must(template.ParseFiles("form.html"))
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			tmpl.Execute(w, nil)
			return
		}

		form := SignUp{
			Email:    r.FormValue("email"),
			UserName: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		_, err := strconv.ParseBool(r.FormValue("Success"))

		IsSuccess := true
		if err == nil {
			fmt.Printf("Error")
			IsSuccess = false
		}

		// do something with form
		_ = form
		tmpl.Execute(w, struct{ Success bool }{IsSuccess})

	}).Methods("POST")

	http.ListenAndServe(":8080", r)
}
