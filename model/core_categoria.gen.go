// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCategoria = "core_categoria"

// Categoria mapped from table <core_categoria>
type Categoria struct {
	ID               int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Estado           bool      `gorm:"column:estado;not null" json:"estado"`
	Fechacreado      time.Time `gorm:"column:fechacreado;not null" json:"fechacreado"`
	Fechamodificado  time.Time `gorm:"column:fechamodificado;not null" json:"fechamodificado"`
	Descripcion      string    `gorm:"column:descripcion;not null" json:"descripcion"`
	ImageDefault     string    `gorm:"column:image_default;not null" json:"image_default"`
	Descuento        int32     `gorm:"column:descuento;not null" json:"descuento"`
	CodigoPais       string    `gorm:"column:codigo_pais" json:"codigo_pais"`
	CodigoLetra      string    `gorm:"column:codigo_letra" json:"codigo_letra"`
	CategoriaPadre *Categoria     `gorm:"column:categoria_padre_id;foreignKey:ID" json:"categoria_padre_id"`
	TipoCategoria  TipoCategoria     `gorm:"column:tipo_categoria_id;foreignKey:ID;not null" json:"tipo_categoria_id"`
}

// TableName Categoria's table name
func (*Categoria) TableName() string {
	return TableNameCategoria
}