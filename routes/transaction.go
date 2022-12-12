package routes

import (
	"server/handlers"
	"server/pkg/middleware"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoute(r *mux.Router) {
	transactionRepo := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepo)

	r.HandleFunc("/transactions", h.FindTransaction).Methods("GET")
	r.HandleFunc("/transaction/user/{id}", middleware.Auth(h.FindTransactionUserId)).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransactionID).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	// midtrans
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}