package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
	"time"

	"github.com/nrfta/go-paging"
	"github.com/sonderkevin/gql/graph/generated"
	"github.com/sonderkevin/gql/graph/model"
	db_model "github.com/sonderkevin/gql/model"
)

// HasPreviousPage is the resolver for the hasPreviousPage field.
func (r *pageInfoResolver) HasPreviousPage(ctx context.Context, obj *paging.PageInfo) (bool, error) {
	panic(fmt.Errorf("not implemented: HasPreviousPage - hasPreviousPage"))
}

// HasNextPage is the resolver for the hasNextPage field.
func (r *pageInfoResolver) HasNextPage(ctx context.Context, obj *paging.PageInfo) (bool, error) {
	panic(fmt.Errorf("not implemented: HasNextPage - hasNextPage"))
}

// TotalCount is the resolver for the totalCount field.
func (r *pageInfoResolver) TotalCount(ctx context.Context, obj *paging.PageInfo) (*int, error) {
	panic(fmt.Errorf("not implemented: TotalCount - totalCount"))
}

// StartCursor is the resolver for the startCursor field.
func (r *pageInfoResolver) StartCursor(ctx context.Context, obj *paging.PageInfo) (*string, error) {
	panic(fmt.Errorf("not implemented: StartCursor - startCursor"))
}

// EndCursor is the resolver for the endCursor field.
func (r *pageInfoResolver) EndCursor(ctx context.Context, obj *paging.PageInfo) (*string, error) {
	panic(fmt.Errorf("not implemented: EndCursor - endCursor"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, page *paging.PageArgs) (*model.UserFriendsConnection, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

func ConvertCategoria(categoria db_model.Categoria, categoriaPadreMap map[int32]*model.CategoriaNode, r *queryResolver) *model.CategoriaNode {
	var categoriaPadre *model.CategoriaNode

	if categoriaPadreMap[categoria.CategoriaPadreID] != nil {
		categoriaPadre = categoriaPadreMap[categoria.CategoriaPadreID]
	}

	if categoria.CategoriaPadreID == 1 {
		categoriaPadre = nil
	} else {
		padre, _ := r.CategoriaService.GetById(int(categoria.CategoriaPadreID))
		categoriaPadre = ConvertCategoria(*padre, categoriaPadreMap, r)
	}

	categoriaPadreMap[categoria.CategoriaPadreID] = categoriaPadre

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
		TipoCategoria: &model.TipoCategoriaNode{
			ID:              string(1),
			Estado:          true,
			Fechacreado:     time.Now().String(),
			Fechamodificado: time.Now().String(),
			Nombre:          "1",
			Abrev:           "a",
			Sub:             nil,
		},
	}

	return result
}

// func ConvertTipoCategoria(tipoCategoria db_model.TipoCategoria) *model.TipoCategoriaNode {
// 	result := &model.TipoCategoriaNode{
// 		ID:              string(tipoCategoria.ID),
// 		Estado:          tipoCategoria.Estado,
// 		Fechacreado:     tipoCategoria.Fechacreado.String(),
// 		Fechamodificado: tipoCategoria.Fechacreado.String(),
// 		Nombre:          tipoCategoria.Nombre,
// 		Abrev:           tipoCategoria.Abrev,
// 		Sub:             &tipoCategoria.Sub,
// 	}
// 	return result
// }

// AllCategorias is the resolver for the allCategorias field.
func (r *queryResolver) AllCategorias(ctx context.Context, before *string, after *string, first *int, last *int, descripcion *string, descripcionIcontains *string, tipoCategoriaNombre *string, tipoCategoriaNombreIcontains *string, tipoCategoriaAbrev *string, tipoCategoriaSub *string, tipoCategoriaID *string, categoriaPadreID *string, codigoPais *string, codigoLetra *string, id *string) (*model.CategoriaNodeConnection, error) {
	dbCategories, err := r.CategoriaService.GetAllCategoria()
	if err != nil {
		return nil, err
	}

	totalCount := len(dbCategories)

	paginator := paging.NewOffsetPaginator(&paging.PageArgs{First: first, After: after}, int64(totalCount))

	result := &model.CategoriaNodeConnection{
		PageInfo: &paginator.PageInfo,
	}
	categoriaPadreMap := map[int32]*model.CategoriaNode{}

	for i, categoria := range dbCategories {
		result.Edges = append(result.Edges, &model.CategoriaNodeEdge{
			Cursor: *paging.EncodeOffsetCursor(paginator.Offset + i + 1),
			Node:   ConvertCategoria(*categoria, categoriaPadreMap, r),
		})
	}

	return result, nil
}

// PageInfo returns generated.PageInfoResolver implementation.
func (r *Resolver) PageInfo() generated.PageInfoResolver { return &pageInfoResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type pageInfoResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
