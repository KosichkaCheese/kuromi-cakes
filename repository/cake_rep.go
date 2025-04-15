package repository

import (
	"kuromi_cakes/models"

	"gorm.io/gorm"
)

type CakeRepository interface {
	GetAllCakes() ([]models.Cake, error)
	GetCakeById(id int) (models.Cake, error)
	CreateCake(cake models.Cake) (models.Cake, error)
	DeleteCake(id int) error
	UpdateCake(cake models.Cake) (models.Cake, error)
}

type cakeRepository struct {
	DB *gorm.DB
}

func NewCakeRepository(db *gorm.DB) CakeRepository {
	return &cakeRepository{DB: db}
}

// Получение всех тортов
func (r *cakeRepository) GetAllCakes() ([]models.Cake, error) {
	var cakes []models.Cake
	err := r.DB.Find(&cakes).Error
	return cakes, err
}

// Получение торта по ID
func (r *cakeRepository) GetCakeById(id int) (models.Cake, error) {
	var cake models.Cake
	err := r.DB.Where("id = ?", id).First(&cake).Error
	return cake, err
}

// Создание нового торта
func (r *cakeRepository) CreateCake(cake models.Cake) (models.Cake, error) {
	err := r.DB.Create(&cake).Error
	return cake, err
}

func (r *cakeRepository) DeleteCake(id int) error {
	err := r.DB.Where("id = ?", id).Delete(&models.Cake{}).Error
	return err
}

func (r *cakeRepository) UpdateCake(cake models.Cake) (models.Cake, error) {
	err := r.DB.Updates(&cake).Error
	return cake, err
}
