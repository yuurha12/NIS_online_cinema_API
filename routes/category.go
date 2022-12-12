package routes

import (
	"github.com/gorilla/mux"
	"server/handlers"
	"server/pkg/middleware"
	"server/pkg/mysql"
	"server/repositories"
)

func CategoryRoutes(r *mux.Router) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	r.HandleFunc("/categories", h.FindCategories).Methods("GET")
	r.HandleFunc("/category/{id}", h.GetCategory).Methods("GET")
	r.HandleFunc("/category", middleware.Auth(h.CreateCategory)).Methods("POST")
	r.HandleFunc("/category/{id}", h.UpdateCategory).Methods("PATCH")
	r.HandleFunc("/category/{id}", h.DeleteCategory).Methods("DELETE")

}
