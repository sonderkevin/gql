package model

import (
	"time"
)

const TableNameCategoryType = "core_tipocategoria"

// CategoryType mapped from table <core_tipocategoria>
type CategoryType struct {
	ID              int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Estado          bool      `gorm:"column:estado;not null" json:"estado"`
	Fechacreado     time.Time `gorm:"column:fechacreado;not null" json:"fechacreado"`
	Fechamodificado time.Time `gorm:"column:fechamodificado;not null" json:"fechamodificado"`
	Nombre          string    `gorm:"column:nombre;not null" json:"nombre"`
	Abrev           string    `gorm:"column:abrev;not null" json:"abrev"`
	Sub             string    `gorm:"column:sub" json:"sub"`
}

// TableName CategoryType's table name
func (*CategoryType) TableName() string {
	return TableNameCategoryType
}
