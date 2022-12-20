package repositories

import (
	"github.com/sonderkevin/gql/model"
	"gorm.io/gorm"
)

type CategoriaRepository struct {
	DB *gorm.DB
}

func (r CategoriaRepository) Save(categoria *model.Categoria) (*model.Categoria, error) {
	if err := r.DB.Create(categoria).Error; err != nil {
		return nil, err
	}

	return categoria, nil
}

func (r CategoriaRepository) GetAll() ([]*model.Categoria, error) {
	var categorias []*model.Categoria

	if err := r.DB.Find(&categorias).Error; err != nil {
		return nil, err
	}

	return categorias, nil
}

func (r CategoriaRepository) GetById(id string) (*model.Categoria, error) {
	categoria := &model.Categoria{}

	if err := r.DB.First(categoria, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return categoria, nil
}

func (r CategoriaRepository) Update(categoria *model.Categoria) (*model.Categoria, error) {
	if err := r.DB.Save(categoria).Error; err != nil {
		return nil, err
	}

	return categoria, nil
}

func (r CategoriaRepository) Delete(categoria *model.Categoria) (*model.Categoria, error) {
	if err := r.DB.Delete(categoria).Error; err != nil {
		return nil, err
	}

	return categoria, nil
}
