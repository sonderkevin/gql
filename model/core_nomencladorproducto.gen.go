package model

import (
	"time"
)

const TableNameCoreNomencladorproducto = "core_nomencladorproducto"

// CoreNomencladorproducto mapped from table <core_nomencladorproducto>
type CoreNomencladorproducto struct {
	ID              int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Estado          bool      `gorm:"column:estado;not null" json:"estado"`
	Fechacreado     time.Time `gorm:"column:fechacreado;not null" json:"fechacreado"`
	Fechamodificado time.Time `gorm:"column:fechamodificado;not null" json:"fechamodificado"`
	Tipo            string    `gorm:"column:tipo;not null" json:"tipo"`
	Descripcion     string    `gorm:"column:descripcion;not null" json:"descripcion"`
	SubcategoriaID  int       `gorm:"column:subcategoria_id" json:"subcategoria_id"`
}

// TableName CoreNomencladorproducto's table name
func (*CoreNomencladorproducto) TableName() string {
	return TableNameCoreNomencladorproducto
}
