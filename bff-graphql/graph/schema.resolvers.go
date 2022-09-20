package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ruancaetano/grpc-graphql-store/bff-graphql/graph/generated"
	"github.com/ruancaetano/grpc-graphql-store/bff-graphql/graph/model"
	"github.com/ruancaetano/grpc-graphql-store/users/pbusers"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := r.UserServiceClient.CreateUser(ctx, &pbusers.CreateUserRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        user.GetId(),
		Name:      user.GetName(),
		Email:     user.GetEmail(),
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := r.Resolver.UserServiceClient.GetUserById(ctx, &pbusers.GetUserRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        user.GetId(),
		Name:      user.GetName(),
		Email:     user.GetEmail(),
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
