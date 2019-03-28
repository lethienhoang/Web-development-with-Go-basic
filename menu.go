package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type MenuItem struct {
	Title    string
	URL      string
	IsActive bool
}

type MenuPage struct {
	PageTitle string
	MenuItems []MenuItem
}

func main() {

	r := mux.NewRouter()

	tmpl := template.Must(template.ParseFiles("menu.html"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := MenuPage{
			PageTitle: "List Foods",
			MenuItems: []MenuItem{
				{Title: "Pork rice", URL: "/foods?query='porkrice'", IsActive: true},
				{Title: "Beef rice", URL: "/foods?query='beefrice'", IsActive: true},
				{Title: "Chicken rice", URL: "/foods?query='chichkenrice'", IsActive: true},
				{Title: "Fish rice", URL: "/foods?query='chichkenrice'", IsActive: false},
			},
		}
		tmpl.Execute(w, data)
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}
