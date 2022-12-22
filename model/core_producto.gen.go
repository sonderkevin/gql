package model

import (
	"time"
)

const TableNameProduct = "core_producto"

// Product mapped from table <core_producto>
type Product struct {
	ID                  int                      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Estado              bool                     `gorm:"column:estado;not null" json:"estado"`
	Fechacreado         time.Time                `gorm:"column:fechacreado;not null" json:"fechacreado"`
	Fechamodificado     time.Time                `gorm:"column:fechamodificado;not null" json:"fechamodificado"`
	Usuariocrea         string                   `gorm:"column:usuariocrea" json:"usuariocrea"`
	Codigo              string                   `gorm:"column:codigo" json:"codigo"`
	DescripcionCompleta string                   `gorm:"column:descripcion_completa" json:"descripcion_completa"`
	Preciocompra        float64                  `gorm:"column:preciocompra;not null" json:"preciocompra"`
	Precioventa         float64                  `gorm:"column:precioventa;not null" json:"precioventa"`
	Cantidad            int                      `gorm:"column:cantidad;not null" json:"cantidad"`
	Lote                bool                     `gorm:"column:lote;not null" json:"lote"`
	Visible             bool                     `gorm:"column:visible;not null" json:"visible"`
	Descuento           int                      `gorm:"column:descuento;not null" json:"descuento"`
	DescripcionID       int                      `gorm:"column:descripcion_id;not null" json:"descripcion_id"`
	Descripcion         *CoreNomencladorproducto `gorm:"many2one:DecripcionID" json:"descripcion"`
	EmpresaID           int                      `gorm:"column:empresa_id;not null" json:"empresa_id"`
	Empresa             *CoreEmpresa             `gorm:"many2one:EmpresaID" json:"empresa"`
	MarcaID             int                      `gorm:"column:marca_id" json:"marca_id"`
	Marca               *CoreMarca               `gorm:"many2one:MarcaID" json:"marca"`
	Latitud             float64                  `gorm:"column:latitud;not null" json:"latitud"`
	Longitud            float64                  `gorm:"column:longitud;not null" json:"longitud"`
	Destacado           int                      `gorm:"column:destacado" json:"destacado"`
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}
