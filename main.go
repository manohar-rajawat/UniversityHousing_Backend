package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<Center><h1>Welcome to my web server!</h1></Center>"))
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
