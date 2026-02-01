package service

import (
	"errors"
	"product-api/model"
	"product-api/repository"
)

type ProductServiceInterface interface {
	GetAll() ([]model.Product, error)
	Create(data *model.Product) error
	GetByID(id int) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id int) error
}

type productService struct {
	productRepo  repository.ProductRepositoryInterface
	categoryRepo repository.CategoryRepositoryInterface
}

func NewProductService(productRepo repository.ProductRepositoryInterface, categoryRepo repository.CategoryRepositoryInterface) ProductServiceInterface {
	return &productService{productRepo: productRepo, categoryRepo: categoryRepo}
}

func (s *productService) GetAll() ([]model.Product, error) {
	products, err := s.productRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []model.Product
	for _, product := range products {
		product.Category, err = s.categoryRepo.GetByID(product.CategoryID)
		if err != nil {
			return nil, err
		}
		result = append(result, product)
	}
	return result, nil
}

func (s *productService) Create(data *model.Product) error {
	_, err := s.categoryRepo.GetByID(data.CategoryID)
	if err != nil {
		return errors.New("category not found")
	}
	err = s.productRepo.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) GetByID(id int) (*model.Product, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	product.Category, err = s.categoryRepo.GetByID(product.CategoryID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) Update(product *model.Product) error {
	_, err := s.categoryRepo.GetByID(product.CategoryID)
	if err != nil {
		return errors.New("category not found")
	}
	err = s.productRepo.Update(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) Delete(id int) error {
	return s.productRepo.Delete(id)
}
