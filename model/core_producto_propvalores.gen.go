// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCoreProductoPropvalore = "core_producto_propvalores"

// CoreProductoPropvalore mapped from table <core_producto_propvalores>
type CoreProductoPropvalore struct {
	ID                   int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProductoID           int32 `gorm:"column:producto_id;not null" json:"producto_id"`
	PropiedadprodvalorID int32 `gorm:"column:propiedadprodvalor_id;not null" json:"propiedadprodvalor_id"`
}

// TableName CoreProductoPropvalore's table name
func (*CoreProductoPropvalore) TableName() string {
	return TableNameCoreProductoPropvalore
}
