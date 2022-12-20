package services

import (
	"github.com/sonderkevin/gql/app/repositories"
	"github.com/sonderkevin/gql/model"
)

type CategoriaService struct {
	Repository *repositories.CategoriaRepository
}

func (s *CategoriaService) GetAllCategoria() ([]*model.Categoria, error) {
	dbCategories, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return dbCategories, nil
}

func (s *CategoriaService) GetById(id int) (*model.Categoria, error) {
	dbCategoria, err := s.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dbCategoria, nil
}
