package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	HouseRoutes(r)
	TransactionRoutes(r)
	// ProfileRoutes(r)
	// ProductRoutes(r)
	AuthRoutes(r)
}
