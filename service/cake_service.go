package service

import (
	"kuromi_cakes/models"
	"kuromi_cakes/repository"
)

type CakeService interface {
	GetAllCakes() ([]models.Cake, error)
	GetCakeById(id int) (models.Cake, error)
	CreateCake(cake models.Cake) (models.Cake, error)
}

type cakeService struct {
	repo repository.CakeRepository
}

func NewCakeService(repo repository.CakeRepository) CakeService {
	return &cakeService{repo: repo}
}

// Получение всех тортов
func (s *cakeService) GetAllCakes() ([]models.Cake, error) {
	return s.repo.GetAllCakes()
}

// Получение торта по ID
func (s *cakeService) GetCakeById(id int) (models.Cake, error) {
	return s.repo.GetCakeById(id)
}

// Создание нового торта
func (s *cakeService) CreateCake(cake models.Cake) (models.Cake, error) {
	return s.repo.CreateCake(cake)
}
