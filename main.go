package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}

func main() {

	http.HandleFunc("/", HelloServer)
	fmt.Println("Server started at port 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
