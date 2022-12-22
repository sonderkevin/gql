package graph

import (
	"github.com/sonderkevin/gql/app/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryService *services.CategoryService
	ProductService  *services.ProductService
}
