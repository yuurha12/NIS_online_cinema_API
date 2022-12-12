package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	FilmRoutes(r)
	UserRoutes(r)
	TransactionRoute(r)
	AuthRoutes(r)
	CategoryRoutes(r)
}
