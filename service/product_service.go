package service

import (
	"kasir-api/models"
	"kasir-api/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(name string) ([]models.Product, error) {
	return s.repo.FetchAll(name)
}

func (s *ProductService) GetByID(id int) (models.Product, error) {
	return s.repo.FetchByID(id)
}

func (s *ProductService) Create(input *models.Product) error {
	return s.repo.Store(input)
}

func (s *ProductService) Update(id int, input models.Product) (models.Product, error) {
	existingProduct, err := s.repo.FetchByID(id)
	if err != nil {
		return models.Product{}, err
	}

	// Update field
	existingProduct.Name = input.Name
	existingProduct.Price = input.Price
	existingProduct.Stock = input.Stock

	// Cek jika category ID berubah
	if input.CategoryID != 0 {
		existingProduct.CategoryID = input.CategoryID
	}

	err = s.repo.Update(&existingProduct)
	if err != nil {
		return models.Product{}, err
	}

	return existingProduct, nil
}

func (s *ProductService) Delete(id int) error {
	_, err := s.repo.FetchByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
