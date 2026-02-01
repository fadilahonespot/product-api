package service

import (
	"product-api/model"
	"product-api/repository"
)

type CategoryServiceInterface interface {
	GetAll() ([]model.Category, error)
	Create(category *model.Category) error
	GetByID(id int) (*model.Category, error)
	Update(category *model.Category) error
	Delete(id int) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepositoryInterface
}

func NewCategoryService(categoryRepo repository.CategoryRepositoryInterface) CategoryServiceInterface {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) GetAll() ([]model.Category, error) {
	return s.categoryRepo.GetAll()
}

func (s *categoryService) Create(category *model.Category) error {
	return s.categoryRepo.Create(category)
}

func (s *categoryService) GetByID(id int) (*model.Category, error) {
	return s.categoryRepo.GetByID(id)
}

func (s *categoryService) Update(category *model.Category) error {
	return s.categoryRepo.Update(category)
}

func (s *categoryService) Delete(id int) error {
	return s.categoryRepo.Delete(id)
}