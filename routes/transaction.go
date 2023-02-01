package routes

import (
	"task1/handlers"
	"task1/pkg/middleware"
	"task1/pkg/mysql"
	"task1/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransactionRepository(transactionRepository)
	// h := handlers.HandlerPropertyHouse(houseRepository)

	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transactionCreate",middleware.Auth(middleware.UploadFile(h.CreateTransaction,"image_payment"))).Methods("POST")
	r.HandleFunc("/transactionUpdate/{id}",middleware.Auth(middleware.UploadFile( h.UpdateTransaction,"image_payment"))).Methods("PATCH")
	// r.HandleFunc("/houseUpdate/{id}",middleware.Auth(middleware.UploadFile(h.UpdateTransaction,"image_payment"))).Methods("PATCH")
	// r.HandleFunc("/house/{id}", h.UpdateHouse).Methods("PATCH")
	// r.HandleFunc("/deleteHouse/{id}", h.DeleteHouse).Methods("DELETE")
}
