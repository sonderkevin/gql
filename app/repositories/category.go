package repositories

import (
	"github.com/sonderkevin/gql/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (r CategoryRepository) Save(category *model.Category) (*model.Category, error) {
	if err := r.DB.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r CategoryRepository) GetAll() ([]*model.Category, error) {
	var categorias []*model.Category

	if err := r.DB.Preload("CategoriaPadre.CategoriaPadre.CategoriaPadre.CategoriaPadre.CategoriaPadre").Preload("TipoCategoria").Find(&categorias).Error; err != nil {
		return nil, err
	}

	// if err := GetAllRecursive(r.DB, &categorias); err != nil {
	// 	return nil, err
	// }

	return categorias, nil
}

// func GetAllRecursive(db *gorm.DB, categorias *[]*model.Category) error {
// 	if err := db.Preload("CategoriaPadre").Preload("TipoCategoria").Find(&categorias).Error; err != nil {
// 		return err
// 	}

// 	for _, category := range *categorias {
// 		if category.CategoriaPadre != nil && category.CategoriaPadreID != 1 {
// 			if err := GetAllRecursive(db, &[]*model.Category{category.CategoriaPadre}); err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

func (r CategoryRepository) GetById(id int) (*model.Category, error) {
	category := &model.Category{}

	if err := r.DB.First(category, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r CategoryRepository) Update(category *model.Category) (*model.Category, error) {
	if err := r.DB.Save(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r CategoryRepository) Delete(category *model.Category) (*model.Category, error) {
	if err := r.DB.Delete(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
