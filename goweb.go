package main

import (
	"fmt"
	"log"
	"net/http"
)

/* ===== HANDLERS ===== */

// Hello world
func Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello world")
}

// Page not found
func PageNF(res http.ResponseWriter, req *http.Request) {
	http.NotFound(res, req)
}

// Errors
func Error(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "Has error!!", http.StatusConflict)
}

// Argument - URL
func Greet(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Println(req.URL.RawQuery)
	fmt.Println(req.URL.Query())

	name := req.URL.Query().Get("name")
	age := req.URL.Query().Get("age")

	fmt.Fprintf(res, "Name: %s - Age: %s", name, age)
}

/* ===== MAIN ===== */
func main() {
	/* ===== ROUTER - MUX ===== */
	mux := http.NewServeMux()

	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/page", PageNF)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/greet", Greet)

	/* ===== SERVER ===== */
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("App listening at http://localhost:3000")
	log.Fatal(server.ListenAndServe())
}
