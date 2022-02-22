package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/manohar-rajawat/universityhousing/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port Not Found")
	}
	r := router.Router()
	fmt.Println("Starting server on the port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func run() string {
	return "Setup Travis CI for Golang project"
}
