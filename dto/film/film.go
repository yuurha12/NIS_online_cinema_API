package filmsdto

import "server/models"

type CreateFilmRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	CategoryID  int    `json:"category_id"`
	Price       int    `json:"price" form:"price" gorm:"type: int" validate:"required"`
	LinkFilm    string `json:"linkfilm" validate:"required"`
	Description string `json:"description" validate:"required"`
	Thumbnail   string `json:"thumbnail" form:"id" validate:"required"`
}

type UpdateFilm struct {
	Title       string `json:"title" form:"title"`
	CategoryID  int    `json:"category_id"`
	Price       int    `json:"price" form:"price"`
	LinkFilm    string `json:"linkfilm"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail" form:"id"`
}

type FilmResponse struct {
	ID          int                     `json:"id"`
	Title       string                  `json:"title" form:"title" validate:"required"`
	CategoryID  int                     `json:"category_id" form:"category_id"`
	Category    models.CategoryResponse `json:"category" form:"category"`
	Price       int                     `json:"price" form:"price" gorm:"type: int" validate:"required"`
	LinkFilm    string                  `json:"linkfilm" validate:"required"`
	Description string                  `json:"description" validate:"required"`
	Thumbnail   string                  `json:"thumbnail" form:"id" validate:"required"`
}
