package graph

import (
	cauth "github.com/ruancaetano/grpc-graphql-store/auth/clients"
	corders "github.com/ruancaetano/grpc-graphql-store/orders/clients"
	cproducts "github.com/ruancaetano/grpc-graphql-store/products/clients"
	cusers "github.com/ruancaetano/grpc-graphql-store/users/clients"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	AuthServiceClient    *cauth.AuthServiceClient
	UserServiceClient    *cusers.UserServiceClient
	ProductServiceClient *cproducts.ProductServiceClient
	OrderServiceClient   *corders.OrderServiceClient
}
