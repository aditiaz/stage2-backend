package routes

import (
	"task1/handlers"
	"task1/pkg/mysql"
	"task1/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	// Create "/login" route using handler Login and method POST here ...
	r.HandleFunc("/login", h.Login).Methods("POST")
}
