// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCorePropiedadprod = "core_propiedadprod"

// CorePropiedadprod mapped from table <core_propiedadprod>
type CorePropiedadprod struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Estado          bool      `gorm:"column:estado;not null" json:"estado"`
	Fechacreado     time.Time `gorm:"column:fechacreado;not null" json:"fechacreado"`
	Fechamodificado time.Time `gorm:"column:fechamodificado;not null" json:"fechamodificado"`
	Usuariocrea     string    `gorm:"column:usuariocrea" json:"usuariocrea"`
	Tipo            string    `gorm:"column:tipo;not null" json:"tipo"`
	Descripcion     string    `gorm:"column:descripcion;not null" json:"descripcion"`
}

// TableName CorePropiedadprod's table name
func (*CorePropiedadprod) TableName() string {
	return TableNameCorePropiedadprod
}
