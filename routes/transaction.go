package routes

import (
	"server/handlers"
	"server/pkg/midleware"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoute(r *mux.Router) {
	transactionRepo := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepo)

	r.HandleFunc("/transactions", h.FindTransaction).Methods("GET")
	r.HandleFunc("/transaction/user/{id}", midleware.Auth(h.FindTransactionUserId)).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransactionID).Methods("GET")
	r.HandleFunc("/transaction/create", midleware.Auth(h.CreateTransaction)).Methods("POST")
	// midtrans
	r.HandleFunc("/notif", h.Notification).Methods("POST")
}
