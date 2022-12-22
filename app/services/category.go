package services

import (
	"github.com/sonderkevin/gql/app/repositories"
	"github.com/sonderkevin/gql/model"
)

type CategoryService struct {
	Repository *repositories.CategoryRepository
}

func (s *CategoryService) GetAllCategory() ([]*model.Category, error) {
	dbCategories, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return dbCategories, nil
}

func (s *CategoryService) GetById(id int) (*model.Category, error) {
	dbCategory, err := s.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dbCategory, nil
}
