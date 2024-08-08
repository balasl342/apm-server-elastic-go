package main

import (
	"log"
	"net/http"

	"github.com/balasl342/apm-server-elastic-go/config"
	"github.com/balasl342/apm-server-elastic-go/database"
	"github.com/balasl342/apm-server-elastic-go/middleware"
	"github.com/balasl342/apm-server-elastic-go/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Load config
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// Initialize Elastic app
	tracer := middleware.Newelasticapp()
	defer tracer.Close()
	// Initialize database
	database.InitDatabase()

	// Create a new router
	r := mux.NewRouter()
	// Setup routes
	routes.SetupRoutes(r)

	wrappedRouter := middleware.WrapHandleFunc(r)

	// Start the server
	http.ListenAndServe(":8000", wrappedRouter)
}
