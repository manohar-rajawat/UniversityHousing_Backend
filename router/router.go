package router

import (
	"github.com/manohar-rajawat/universityhousing/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/university/{name}", middleware.GetUniversity).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/university/", middleware.GetAllUniversity).Methods("GET", "OPTIONS")
	return router
}
