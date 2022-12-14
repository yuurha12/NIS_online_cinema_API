package routes

import (
	"server/handlers"
	"server/pkg/midleware"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/gorilla/mux"
)

func FilmsRoute(r *mux.Router) {
	categoryRepo := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(categoryRepo)

	r.HandleFunc("/films", h.GetFilm).Methods("GET")
	r.HandleFunc("/film/{id}", h.GetFilmId).Methods("GET")
	r.HandleFunc("/film/create", midleware.Auth(midleware.UploadFile(h.CreateFilm))).Methods("POST")
	r.HandleFunc("/film/update/{id}", midleware.Auth(midleware.UploadFile(h.UpdateFilm))).Methods("PATCH")
	r.HandleFunc("/film/delete/{id}", h.DeleteFilm).Methods("DELETE")

}
