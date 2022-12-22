package model

import (
	"time"
)

const TableNameCategory = "core_categoria"

// Category mapped from table <core_categoria>
type Category struct {
	ID               int           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Estado           bool          `gorm:"column:estado;not null" json:"estado"`
	Fechacreado      time.Time     `gorm:"column:fechacreado;not null" json:"fechacreado"`
	Fechamodificado  time.Time     `gorm:"column:fechamodificado;not null" json:"fechamodificado"`
	Descripcion      string        `gorm:"column:descripcion;not null" json:"descripcion"`
	ImageDefault     string        `gorm:"column:image_default;not null" json:"image_default"`
	Descuento        int           `gorm:"column:descuento;not null" json:"descuento"`
	CodigoPais       string        `gorm:"column:codigo_pais" json:"codigo_pais"`
	CodigoLetra      string        `gorm:"column:codigo_letra" json:"codigo_letra"`
	CategoriaPadreID int           `gorm:"column:categoria_padre_id" json:"categoria_padre_id"`
	CategoriaPadre   *Category     `gorm:"many2one:CategoriaPadreID" json:"categoria_padre"`
	TipoCategoriaID  int           `gorm:"column:tipo_categoria_id;not null" json:"tipo_categoria_id"`
	TipoCategoria    *CategoryType `gorm:"many2one:TipoCategoriaID" json:"tipo_categoria"`
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
