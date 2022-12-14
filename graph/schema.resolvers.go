package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/nrfta/go-paging"
	"github.com/sonderkevin/gql/graph/generated"
	"github.com/sonderkevin/gql/graph/model"
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
func (r *queryResolver) Users(ctx context.Context, page *paging.PageArgs) (*model.UserNodeConnection, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Categoria is the resolver for the categoria field.
func (r *queryResolver) Categoria(ctx context.Context, id string) (*model.Categoria, error) {
	panic(fmt.Errorf("not implemented: Categoria - categoria"))
}

// Categorias is the resolver for the categorias field.
func (r *queryResolver) Categorias(ctx context.Context, page *paging.PageArgs, input *model.CategoriaInput) (*model.CategoriaNodeConnection, error) {
	dbCategories, err := r.CategoryService.GetAllCategory()
	if err != nil {
		return nil, err
	}

	totalCount := len(dbCategories)

	paginator := paging.NewOffsetPaginator(page, int64(totalCount))

	result := &model.CategoriaNodeConnection{
		PageInfo: &paginator.PageInfo,
	}

	for i, category := range dbCategories {
		result.Edges = append(result.Edges, &model.CategoriaNodeEdge{
			Cursor: *paging.EncodeOffsetCursor(paginator.Offset + i + 1),
			Node:   ConvertCategory(category),
		})
	}

	return result, nil
}

// Producto is the resolver for the producto field.
func (r *queryResolver) Producto(ctx context.Context, id string) (*model.Producto, error) {
	decodedId, err := DecodeBase64ToInt(id, "product")
	if err != nil {
		return nil, err
	}

	dbProduct, err := r.ProductService.GetById(decodedId)
	if err != nil {
		return nil, err
	}

	return ConvertProduct(dbProduct), nil

}

// Productos is the resolver for the productos field.
func (r *queryResolver) Productos(ctx context.Context, page *paging.PageArgs, input *model.ProductoInput) (*model.ProductoNodeConnection, error) {
	dbProducts, err := r.ProductService.GetAllProduct()
	if err != nil {
		return nil, err
	}

	totalCount := len(dbProducts)

	paginator := paging.NewOffsetPaginator(page, int64(totalCount))

	result := &model.ProductoNodeConnection{
		PageInfo: &paginator.PageInfo,
	}

	for i, product := range dbProducts {
		result.Edges = append(result.Edges, &model.ProductoNodeEdge{
			Cursor: *paging.EncodeOffsetCursor(paginator.Offset + i + 1),
			Node:   ConvertProduct(product),
		})
	}

	return result, nil
}

// Last is the resolver for the last field.
func (r *pageArgsResolver) Last(ctx context.Context, obj *paging.PageArgs, data *int) error {
	panic(fmt.Errorf("not implemented: Last - last"))
}

// Before is the resolver for the before field.
func (r *pageArgsResolver) Before(ctx context.Context, obj *paging.PageArgs, data *string) error {
	panic(fmt.Errorf("not implemented: Before - before"))
}

// PageInfo returns generated.PageInfoResolver implementation.
func (r *Resolver) PageInfo() generated.PageInfoResolver { return &pageInfoResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// PageArgs returns generated.PageArgsResolver implementation.
func (r *Resolver) PageArgs() generated.PageArgsResolver { return &pageArgsResolver{r} }

type pageInfoResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type pageArgsResolver struct{ *Resolver }
