package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	FindFilm() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	CreateFilm(Film models.Film) (models.Film, error)
	UpdateFilm(Film models.Film) (models.Film, error)
	DeleteFilm(Film models.Film) (models.Film, error)
}


func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFilm() ([]models.Film, error) {
	var Films []models.Film
	err := r.db.Preload("Category").Find(&Films).Error

	return Films, err
}

func (r *repository) FindCategoryById(CategoryID int) (models.Category, error) {
	var categories models.Category
	err := r.db.Find(&categories, CategoryID).Error

	return categories, err
}

func (r *repository) GetFilm(ID int) (models.Film, error) {
	var Film models.Film
	err := r.db.Preload("Category").First(&Film, ID).Error

	return Film, err
}

func (r *repository) CreateFilm(Film models.Film) (models.Film, error) {
	err := r.db.Create(&Film).Error

	return Film, err
}

func (r *repository) UpdateFilm(Film models.Film) (models.Film, error) {
	r.db.Model(&Film).Association("Category").Replace(Film.Category)

	err := r.db.Save(&Film).Error

	return Film, err
}

func (r *repository) DeleteFilm(Film models.Film) (models.Film, error) {
	err := r.db.Delete(&Film).Error

	return Film, err
}
