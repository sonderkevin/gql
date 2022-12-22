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

func ConvertCategory(category *db_model.Category) *model.CategoriaNode {
	var categoriaPadre *model.CategoriaNode

	if category.CategoriaPadre == nil || category.CategoriaPadreID == 1 {
		categoriaPadre = nil
	} else {
		categoriaPadre = ConvertCategory(category.CategoriaPadre)
	}

	var tipoCategoria *model.TipoCategoriaNode
	if category.TipoCategoria == nil {
		tipoCategoria = nil
	} else {
		tipoCategoria = ConvertTipoCategoria(category.TipoCategoria)
	}

	result := &model.CategoriaNode{
		ID:              EncodeIntToBase64(category.ID, "category"),
		Estado:          category.Estado,
		Fechacreado:     category.Fechacreado.String(),
		Fechamodificado: category.Fechamodificado.String(),
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

func ConvertTipoCategoria(tipoCategoria *db_model.CategoryType) *model.TipoCategoriaNode {
	result := &model.TipoCategoriaNode{
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

func ConvertProduct(product *db_model.Product) *model.ProductoNode {
	return &model.ProductoNode{
		ID:                  EncodeIntToBase64(product.ID, "product"),
		Estado:              product.Estado,
		Fechacreado:         product.Fechacreado.String(),
		Fechamodificado:     product.Fechamodificado.String(),
		Usuariocrea:         &product.Usuariocrea,
		Codigo:              &product.Codigo,
		DescripcionCompleta: &product.DescripcionCompleta,
		Preciocompra:        product.Preciocompra,
		Precioventa:         product.Precioventa,
		Cantidad:            product.Cantidad,
		Lote:                product.Lote,
		Visible:             product.Visible,
		Descuento:           product.Descuento,
		Latitud:             product.Latitud,
		Longitud:            product.Longitud,
		Destacado:           &product.Destacado,
		Descripcion:         ConvertDescripcion(product.Descripcion),
		Empresa:             ConvertEmpresa(product.Empresa),
		Count:               nil,
	}
}

func ConvertDescripcion(descripcion *db_model.CoreNomencladorproducto) *model.NomProductoNode {
	return &model.NomProductoNode{
		ID:              EncodeIntToBase64(descripcion.ID, "nomencladorProducto"),
		Estado:          descripcion.Estado,
		Fechacreado:     descripcion.Fechacreado.String(),
		Fechamodificado: descripcion.Fechamodificado.String(),
		Tipo:            descripcion.Tipo,
		Descripcion:     descripcion.Descripcion,
		Propiedades:     nil,
		ProductoSet:     nil,
	}
}

func ConvertEmpresa(empresa *db_model.CoreEmpresa) *model.EmpresaNode {
	return &model.EmpresaNode{
		ID:              EncodeIntToBase64(empresa.ID, "empresa"),
		Estado:          empresa.Estado,
		Fechacreado:     empresa.Fechacreado.String(),
		Fechamodificado: empresa.Fechamodificado.String(),
		Usuariocrea:     &empresa.Usuariocrea,
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
