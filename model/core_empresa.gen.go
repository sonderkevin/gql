// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCoreEmpresa = "core_empresa"

// CoreEmpresa mapped from table <core_empresa>
type CoreEmpresa struct {
	ID              int     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Estado          bool      `gorm:"column:estado;not null" json:"estado"`
	Fechacreado     time.Time `gorm:"column:fechacreado;not null" json:"fechacreado"`
	Fechamodificado time.Time `gorm:"column:fechamodificado;not null" json:"fechamodificado"`
	Usuariocrea     string    `gorm:"column:usuariocrea" json:"usuariocrea"`
	Nombre          string    `gorm:"column:nombre;not null" json:"nombre"`
	Tarjeta         string    `gorm:"column:tarjeta;not null" json:"tarjeta"`
	APagar          float64   `gorm:"column:a_pagar;not null" json:"a_pagar"`
	Premium         bool      `gorm:"column:premium;not null" json:"premium"`
	Latitud         float64   `gorm:"column:latitud;not null" json:"latitud"`
	Longitud        float64   `gorm:"column:longitud;not null" json:"longitud"`
	Movil           bool      `gorm:"column:movil;not null" json:"movil"`
	Urlweb          string    `gorm:"column:urlweb;not null" json:"urlweb"`
	Referencia      string    `gorm:"column:referencia" json:"referencia"`
	BarrioID        int     `gorm:"column:barrio_id;not null" json:"barrio_id"`
	CategoriaID     int     `gorm:"column:categoria_id;not null" json:"categoria_id"`
	CompaniaID      int     `gorm:"column:compania_id;not null" json:"compania_id"`
}

// TableName CoreEmpresa's table name
func (*CoreEmpresa) TableName() string {
	return TableNameCoreEmpresa
}