package graph

import (
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
		ID:              string(categoria.ID),
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
		ID:              string(tipoCategoria.ID),
		Estado:          tipoCategoria.Estado,
		Fechacreado:     tipoCategoria.Fechacreado.String(),
		Fechamodificado: tipoCategoria.Fechacreado.String(),
		Nombre:          tipoCategoria.Nombre,
		Abrev:           tipoCategoria.Abrev,
		Sub:             &tipoCategoria.Sub,
	}
	return result
}
