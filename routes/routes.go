package routes

import (
	"github.com/balasl342/apm-server-elastic-go/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	// Register user routes
	r.HandleFunc("/create_user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/get_user/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/get_all_users", handlers.GetAllUsers).Methods("GET")
}
