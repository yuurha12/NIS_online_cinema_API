package categoriesdto

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment"`
	Name   string `json:"name" form:"name" validate:"required"`
	FilmID int    `json:"-"`
}