package routes

import (
	"task1/handlers"
	"task1/pkg/middleware"
	"task1/pkg/mysql"
	"task1/repositories"

	"github.com/gorilla/mux"
)

func HouseRoutes(r *mux.Router) {
	houseRepository := repositories.RepositoryHouse(mysql.DB)
	h := handlers.HandlerPropertyHouse(houseRepository)

	r.HandleFunc("/houses", h.FindHouses).Methods("GET")
	r.HandleFunc("/house/{id}", h.GetHouse).Methods("GET")
	r.HandleFunc("/houseCreate",middleware.Auth(middleware.UploadFile(h.CreateHouse, "image_property"))).Methods("POST")
	r.HandleFunc("/houseUpdate/{id}",middleware.Auth(middleware.UploadFile(h.UpdateHouse, "image_property"))).Methods("PATCH")
	r.HandleFunc("/deleteHouse/{id}", h.DeleteHouse).Methods("DELETE")
}
