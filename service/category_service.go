package service

import (
	"kasir-api/models"
	"kasir-api/repository"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.FetchAll()
}

func (s *CategoryService) GetByID(id int) (models.Category, error) {
	return s.repo.FetchByID(id)
}

func (s *CategoryService) Create(input *models.Category) error {
	// Di sini bisa ditambahkan validasi bisnis, misal nama tidak boleh kosong
	return s.repo.Store(input)
}

func (s *CategoryService) Update(id int, input models.Category) (models.Category, error) {
	// 1. Cek data lama
	existingCategory, err := s.repo.FetchByID(id)
	if err != nil {
		return models.Category{}, err
	}

	// 2. Update field
	existingCategory.Name = input.Name
	existingCategory.Description = input.Description

	// 3. Simpan perubahan
	err = s.repo.Update(&existingCategory)
	if err != nil {
		return models.Category{}, err
	}

	return existingCategory, nil
}

func (s *CategoryService) Delete(id int) error {
	_, err := s.repo.FetchByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
