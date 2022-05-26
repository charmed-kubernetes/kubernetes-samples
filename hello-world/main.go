package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", SayHello)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s", r.RemoteAddr)
	fmt.Fprint(w, "Hello, world!\n")
}
