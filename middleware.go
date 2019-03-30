package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Loggin() Middleware {

	// Create new middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandleFunc
		return func(w http.ResponseWriter, r *http.Request) {
			// Do middleware
			start := time.Now()
			// A defer statement defers the execution of a function until the surrounding function returns.
			defer func() { log.Printf(r.URL.Path, time.Since(start)) }()

			// Call next middleware
			f(w, r)
		}
	}
}

func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(w, r)
		}
	}
}

// Applied middleware
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Chain(Hello, Method("POST"), Loggin()))
	http.ListenAndServe(":8080", r)
}
