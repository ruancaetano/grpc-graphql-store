package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ruancaetano/grpc-graphql-store/bff-graphql/graph/generated"
	"github.com/ruancaetano/grpc-graphql-store/bff-graphql/graph/model"
	"github.com/ruancaetano/grpc-graphql-store/products/pbproducts"
	"github.com/ruancaetano/grpc-graphql-store/users/pbusers"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUserInput) (*model.User, error) {
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

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProductInput) (*model.Product, error) {
	product, err := r.ProductServiceClient.CreateProduct(ctx, &pbproducts.CreateProductRequest{
		Title:       input.Title,
		Description: input.Description,
		Thumb:       input.Thumb,
		Availables:  uint32(input.Availables),
		Price:       float32(input.Price),
	})
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          product.GetId(),
		Title:       product.GetTitle(),
		Description: product.GetDescription(),
		Thumb:       product.GetThumb(),
		Availables:  int(product.GetAvailables()),
		Price:       float64(product.GetPrice()),
		CreatedAt:   &product.CreatedAt,
		UpdatedAt:   &product.UpdatedAt,
	}, nil
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, input model.UpdateProductInput) (*model.Product, error) {
	product, err := r.ProductServiceClient.UpdateProduct(ctx, &pbproducts.UpdateProductRequest{
		Id:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		Thumb:       input.Thumb,
		Price:       float32(input.Price),
	})
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          input.ID,
		Title:       product.GetTitle(),
		Description: product.GetDescription(),
		Thumb:       product.GetThumb(),
		Availables:  int(product.GetAvailables()),
		Price:       float64(product.GetPrice()),
		CreatedAt:   &product.CreatedAt,
		UpdatedAt:   &product.UpdatedAt,
	}, nil
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, input model.DeleteProductInput) (*model.GenericResponse, error) {
	_, err := r.ProductServiceClient.DeleteProduct(ctx, &pbproducts.DeleteProductRequest{
		Id: input.ID,
	})
	if err != nil {
		return nil, err
	}

	return &model.GenericResponse{
		Success: true,
	}, nil
}

// UpdateProductAvailablesValue is the resolver for the updateProductAvailablesValue field.
func (r *mutationResolver) UpdateProductAvailablesValue(ctx context.Context, input model.UpdateProductAvailablesInput) (*model.Product, error) {
	product, err := r.ProductServiceClient.UpdateProductAvailablesValue(ctx, &pbproducts.UpdateProductAvailablesValueRequest{
		Id:         input.ID,
		ValueToAdd: int32(input.ValueToAdd),
	})
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          product.GetId(),
		Title:       product.GetTitle(),
		Description: product.GetDescription(),
		Thumb:       product.GetThumb(),
		Availables:  int(product.GetAvailables()),
		Price:       float64(product.GetPrice()),
		CreatedAt:   &product.CreatedAt,
		UpdatedAt:   &product.UpdatedAt,
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

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	product, err := r.Resolver.ProductServiceClient.GetProductById(ctx, &pbproducts.GetProductByIdRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          product.GetId(),
		Title:       product.GetTitle(),
		Description: product.GetDescription(),
		Thumb:       product.GetThumb(),
		Availables:  int(product.GetAvailables()),
		Price:       float64(product.GetPrice()),
		CreatedAt:   &product.CreatedAt,
		UpdatedAt:   &product.UpdatedAt,
	}, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, page int, limit int) ([]*model.Product, error) {
	response, err := r.Resolver.ProductServiceClient.ListProducts(ctx, &pbproducts.PaginationParams{
		Page:  uint32(page),
		Limit: uint32(limit),
	})
	if err != nil {
		return nil, err
	}

	products := []*model.Product{}

	for _, product := range response.Items {
		products = append(products, &model.Product{
			ID:          product.GetId(),
			Title:       product.GetTitle(),
			Description: product.GetDescription(),
			Thumb:       product.GetThumb(),
			Availables:  int(product.GetAvailables()),
			Price:       float64(product.GetPrice()),
			CreatedAt:   &product.CreatedAt,
			UpdatedAt:   &product.UpdatedAt,
		})
	}

	return products, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
