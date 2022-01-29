package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment variable file")
	}
	return os.Getenv(key)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("<h1>Jai Mahakal! Server</h1>"))
}

func main() {
	mux := http.NewServeMux()
	port := getEnvVariable("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	mux.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
