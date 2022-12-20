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

func ConvertCategoria(categoria db_model.Categoria) *model.CategoriaNode {
	var categoriaPadre *model.CategoriaNode

	if categoria.CategoriaPadre == nil || categoria.CategoriaPadreID == 1 {
		categoriaPadre = nil
	} else {
		categoriaPadre = ConvertCategoria(*categoria.CategoriaPadre)
	}

	var tipoCategoria *model.TipoCategoriaNode
	if categoria.TipoCategoria == nil {
		tipoCategoria = nil
	} else {
		tipoCategoria = ConvertTipoCategoria(*categoria.TipoCategoria)
	}

	result := &model.CategoriaNode{
		ID:              encodeIntToBase64(categoria.ID, "kevinsecret"),
		Estado:          categoria.Estado,
		Fechacreado:     categoria.Fechacreado.String(),
		Fechamodificado: categoria.Fechamodificado.String(),
		Descripcion:     categoria.Descripcion,
		ImageDefault:    categoria.ImageDefault,
		Descuento:       int(categoria.Descuento),
		CodigoPais:      &categoria.CodigoPais,
		CodigoLetra:     &categoria.CodigoLetra,
		CategoriaPadre:  categoriaPadre,
		TipoCategoria:   tipoCategoria,
	}

	return result
}

func ConvertTipoCategoria(tipoCategoria db_model.TipoCategoria) *model.TipoCategoriaNode {
	result := &model.TipoCategoriaNode{
		ID:              encodeIntToBase64(tipoCategoria.ID, "kevinsecret"),
		Estado:          tipoCategoria.Estado,
		Fechacreado:     tipoCategoria.Fechacreado.String(),
		Fechamodificado: tipoCategoria.Fechacreado.String(),
		Nombre:          tipoCategoria.Nombre,
		Abrev:           tipoCategoria.Abrev,
		Sub:             &tipoCategoria.Sub,
	}
	return result
}

func encodeIntToBase64(n int32, secret string) string {
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

func decodeBase64ToInt(s string, secret string) (int32, error) {
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
	return int32(binary.BigEndian.Uint64(b)), nil
}
