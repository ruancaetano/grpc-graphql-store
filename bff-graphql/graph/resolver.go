package graph

import "github.com/ruancaetano/grpc-graphql-store/shared/clients"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserServiceClient *clients.UserServiceClient
}
