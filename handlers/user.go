package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/balasl342/apm-server-elastic-go/middleware"
	"github.com/balasl342/apm-server-elastic-go/models"
	"github.com/balasl342/apm-server-elastic-go/repository"

	"github.com/gorilla/mux"
	"go.elastic.co/apm/v2"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Start a transaction
	tx := apm.TransactionFromContext(r.Context())
	if tx == nil {
		tx = middleware.Tracer.StartTransaction("CreateUser", "request")
		defer tx.End()
	}
	// Start a span for the "createUser-segment"
	span := tx.StartSpan("createUser-segment", "custom", nil)
	defer span.End()

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := models.User{Name: req.Name, Email: req.Email}
	if err := repository.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Start a transaction
	tx := apm.TransactionFromContext(r.Context())
	if tx == nil {
		tx = middleware.Tracer.StartTransaction("GetAllUsers", "request")
		defer tx.End()
	}

	// Start a span for the "GetAllUsers-segment"
	span := tx.StartSpan("GetAllUsers-segment", "custom", nil)
	defer span.End()

	users, err := repository.GetUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Start a transaction
	tx := apm.TransactionFromContext(r.Context())
	if tx == nil {
		tx = middleware.Tracer.StartTransaction("GetUser", "request")
		defer tx.End()
	}

	// Start a span for the "createUser-segment"
	span := tx.StartSpan("GetUser-segment", "custom", nil)
	defer span.End()

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, err := repository.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
