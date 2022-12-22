package services

import (
	"github.com/sonderkevin/gql/app/repositories"
	"github.com/sonderkevin/gql/model"
)

type ProductService struct {
	Repository *repositories.ProductRepository
}

func (s *ProductService) GetAllProduct() ([]*model.Product, error) {
	dbProducts, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return dbProducts, nil
}

func (s *ProductService) GetById(id int) (*model.Product, error) {
	dbProduct, err := s.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dbProduct, nil
}
