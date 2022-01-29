package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Welcome to my web server!</h1>"))
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
