package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func gateWayPipe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", gateWayPipe)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
