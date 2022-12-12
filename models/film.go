package models

import "time"

type Film struct {
	ID          int              `json:"id"`
	Title       string           `json:"title" gorm:"type: varchar(255)"`
	CategoryID  int              `json:"-"`
	Category    CategoryResponse `json:"category"`
	Price       int              `json:"price" gorm:"type: int"`
	LinkFilm    string           `json:"linkfilm" gorm:"type: varchar(255)"`
	Description string           `json:"description" gorm:"type: varchar(255)"`
	Thumbnail   string           `json:"image" gorm:"type: varchar(255)"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type FilmTransaction struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	CategoryID  int              `json:"-"`
	Category    CategoryResponse `json:"category"`
	Price       int              `json:"price"`
	LinkFilm    string           `json:"linkfilm"`
	Description string           `json:"description"`
	Thumbnail   string           `json:"image"`
}

type FilmResponse struct {
	ID          int              `json:"id" `
	Title       string           `json:"title"`
	Price       int              `json:"price"`
	CategoryID  int              `json:"-"`
	Category    CategoryResponse `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Description string           `json:"description"`
	LinkFilm    string           `json:"link"`
	Thumbnail   string           `json:"image"`
}

type FilmCategory struct {
	ID          int              `json:"id" `
	Title       string           `json:"title"`
	CategoryID  int              `json:"-"`
	Category    CategoryResponse `json:"category"`
	Description string           `json:"description"`
	Thumbnail   string           `json:"image"`
}

func (FilmTransaction) TableName() string {
	return "Films"
}

func (FilmResponse) TableName() string {
	return "Films"
}

func (FilmCategory) TableName() string {
	return "Films"
}

func (Film) TableName() string {
	return "Films"
}
