package graph

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"errors"

	"github.com/sonderkevin/gql/graph/model"
	db_model "github.com/sonderkevin/gql/model"
)

func ConvertCategory(category *db_model.Category) *model.Categoria {
	var categoriaPadre *model.Categoria

	if category.CategoriaPadre == nil || category.CategoriaPadreID == 1 {
		categoriaPadre = nil
	} else {
		categoriaPadre = ConvertCategory(category.CategoriaPadre)
	}

	var tipoCategoria *model.TipoCategoria
	if category.TipoCategoria == nil {
		tipoCategoria = nil
	} else {
		tipoCategoria = ConvertTipoCategoria(category.TipoCategoria)
	}

	result := &model.Categoria{
		ID:              EncodeIntToBase64(category.ID, "category"),
		Estado:          category.Estado,
		FechaCreado:     category.Fechacreado.String(),
		FechaModificado: category.Fechamodificado.String(),
		Descripcion:     category.Descripcion,
		ImageDefault:    category.ImageDefault,
		Descuento:       category.Descuento,
		CodigoPais:      &category.CodigoPais,
		CodigoLetra:     &category.CodigoLetra,
		CategoriaPadre:  categoriaPadre,
		TipoCategoria:   tipoCategoria,
	}

	return result
}

func ConvertTipoCategoria(tipoCategoria *db_model.CategoryType) *model.TipoCategoria {
	result := &model.TipoCategoria{
		ID:              EncodeIntToBase64(tipoCategoria.ID, "categoryType"),
		Estado:          tipoCategoria.Estado,
		Fechacreado:     tipoCategoria.Fechacreado.String(),
		Fechamodificado: tipoCategoria.Fechacreado.String(),
		Nombre:          tipoCategoria.Nombre,
		Abrev:           tipoCategoria.Abrev,
		Sub:             &tipoCategoria.Sub,
	}
	return result
}

func ConvertProduct(product *db_model.Product) *model.Producto {
	return &model.Producto{
		ID:                  EncodeIntToBase64(product.ID, "product"),
		Estado:              product.Estado,
		FechaCreado:         product.Fechacreado.String(),
		FechaModificado:     product.Fechamodificado.String(),
		UsuarioCreador:      nil,
		Codigo:              &product.Codigo,
		DescripcionCompleta: &product.DescripcionCompleta,
		PrecioCompra:        product.Preciocompra,
		PrecioVenta:         product.Precioventa,
		Cantidad:            product.Cantidad,
		Lote:                product.Lote,
		Visible:             product.Visible,
		Descuento:           product.Descuento,
		Latitud:             product.Latitud,
		Longitud:            product.Longitud,
		Destacado:           &product.Destacado,
		Descripcion:         ConvertDescripcion(product.Descripcion),
		Empresa:             ConvertEmpresa(product.Empresa),
	}
}

func ConvertDescripcion(descripcion *db_model.CoreNomencladorproducto) *model.NomencladorProducto {
	return &model.NomencladorProducto{
		ID:              EncodeIntToBase64(descripcion.ID, "nomencladorProducto"),
		Estado:          descripcion.Estado,
		FechaCreado:     descripcion.Fechacreado.String(),
		FechaModificado: descripcion.Fechamodificado.String(),
		Tipo:            descripcion.Tipo,
		Descripcion:     descripcion.Descripcion,
		Propiedades:     nil,
		Productos:       nil,
	}
}

func ConvertEmpresa(empresa *db_model.CoreEmpresa) *model.Empresa {
	return &model.Empresa{
		ID:              EncodeIntToBase64(empresa.ID, "empresa"),
		Estado:          empresa.Estado,
		Fechacreado:     empresa.Fechacreado.String(),
		Fechamodificado: empresa.Fechamodificado.String(),
		UsuarioCreador:  nil,
		Compania:        nil,
		Nombre:          empresa.Nombre,
		Tarjeta:         empresa.Tarjeta,
		APagar:          empresa.APagar,
		Premium:         empresa.Estado,
		Categoria:       nil,
		Latitud:         empresa.Latitud,
		Longitud:        empresa.Longitud,
		Movil:           empresa.Movil,
		Urlweb:          empresa.Urlweb,
		Referencia:      &empresa.Referencia,
	}
}

func EncodeIntToBase64(n int, secret string) string {
	// Convert the int to a byte slice
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(n))

	// Generate a MAC using the secret string
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(b)
	sum := mac.Sum(nil)

	// Encode the MAC to a base64 string and return it
	return base64.StdEncoding.EncodeToString(sum)
}

func DecodeBase64ToInt(s string, secret string) (int, error) {
	// Decode the base64 string to a byte slice
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return 0, err
	}

	// Verify the MAC using the secret string
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(b)
	expectedMAC := mac.Sum(nil)
	if !hmac.Equal(b, expectedMAC) {
		return 0, errors.New("invalid MAC")
	}

	// Convert the byte slice to an int and return it
	return int(binary.BigEndian.Uint64(b)), nil
}
