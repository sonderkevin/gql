// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTipoCategoria = "core_tipocategoria"

// TipoCategoria mapped from table <core_tipocategoria>
type TipoCategoria struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Estado          bool      `gorm:"column:estado;not null" json:"estado"`
	Fechacreado     time.Time `gorm:"column:fechacreado;not null" json:"fechacreado"`
	Fechamodificado time.Time `gorm:"column:fechamodificado;not null" json:"fechamodificado"`
	Nombre          string    `gorm:"column:nombre;not null" json:"nombre"`
	Abrev           string    `gorm:"column:abrev;not null" json:"abrev"`
	Sub             string    `gorm:"column:sub" json:"sub"`
}

// TableName TipoCategoria's table name
func (*TipoCategoria) TableName() string {
	return TableNameTipoCategoria
}