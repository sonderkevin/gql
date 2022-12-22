package repositories

import (
	"github.com/sonderkevin/gql/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (r ProductRepository) Save(product *model.Product) (*model.Product, error) {
	if err := r.DB.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductRepository) GetAll() ([]*model.Product, error) {
	var products []*model.Product

	if err := r.DB.Preload("Descripcion").Preload("Empresa").Preload("Marca").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r ProductRepository) GetById(id int) (*model.Product, error) {
	product := &model.Product{}

	if err := r.DB.First(product, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductRepository) Update(product *model.Product) (*model.Product, error) {
	if err := r.DB.Save(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductRepository) Delete(product *model.Product) (*model.Product, error) {
	if err := r.DB.Delete(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
