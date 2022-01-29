package main

import (
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Jai Mahakal! Server</h1>"))
}

func main() {
	mux := http.NewServeMux()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	mux.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
