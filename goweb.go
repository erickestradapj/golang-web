package main

import (
	"fmt"
	"log"
	"net/http"
)

/* ===== HANDLERS ===== */
func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello world")
}

func main() {
	/* ===== ROUTER ===== */
	http.HandleFunc("/", hello)

	/* ===== CREATE SERVER ===== */
	fmt.Print("Example app listening at http://localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))

}
