package models

import "time"

type Transaction struct {
	ID        int             `json:"id"`
	UserID    int             `json:"user_id"`
	User      UserProfile     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status    string          `json:"status"`
	Price     int             `json:"price"`
	FilmID    int             `json:"film_id"`
	Film      FilmTransaction `json:"film"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

type TransactionResponse struct {
	ID     int64       `json:"id"`
	UserID int         `json:"user_id"`
	User   UserProfile `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status string      `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
