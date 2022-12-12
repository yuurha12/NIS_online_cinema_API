package transactiondto

import "server/models"

type CreateTransaction struct {
	ID     int64              `json:"id"`
	UserID int                `json:"user_id" form:"user_id"`
	User   models.UserProfile `json:"user" form:"user"`
	Status string             `json:"status"`
}

type UpdateTransaction struct {
	UserID int    `json:"user_id" form:"user_id"`
	Status string `json:"status"`
	Total  int    `json:"total"`
}

type TransactionResponse struct {
	ID     int64                  `json:"id"`
	UserID int                    `json:"user_id" form:"user_id"`
	FilmID int                    `json:"film_id" form:"film_id"`
	Film   models.FilmTransaction `json:"film"`
	Status string                 `json:"status"`
}

type TransactionRequest struct {
	FilmID        int    `json:"film_id" form:"film_id" gorm:"type: int"`
	Status        string `json:"status" gorm:"type:text" form:"status"`
	AccountNumber int    `json:"account_number" form:"account_number" gorm:"type: int"`
	Price         int    `gorm:"type: int" json:"price" validate:"required"`
}
