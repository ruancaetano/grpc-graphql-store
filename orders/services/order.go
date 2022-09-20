package services

import (
	"context"
	"errors"

	pb "github.com/ruancaetano/grpc-graphql-store/orders/pborders"
	pbproducts "github.com/ruancaetano/grpc-graphql-store/products/pbproducts"
	pbusers "github.com/ruancaetano/grpc-graphql-store/users/pbusers"

	"github.com/ruancaetano/grpc-graphql-store/orders/repositories"
	"github.com/ruancaetano/grpc-graphql-store/shared/clients"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	repository     *repositories.OrderRepository
	productService *clients.ProductServiceClient
	userService    *clients.UserServiceClient
}

// CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error)
// 	ListUserOrders(ctx context.Context, in *ListUserOrdersRequest, opts ...grpc.CallOption) (*ListUserOrdersResponse, error)

func NewOrderService(repository *repositories.OrderRepository, productService *clients.ProductServiceClient, userService *clients.UserServiceClient) *OrderService {
	return &OrderService{
		repository:     repository,
		productService: productService,
		userService:    userService,
	}
}

func (service *OrderService) CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*pb.Order, error) {
	// Quantity validation
	if request.GetQuantity() <= 0 {
		return nil, errors.New("invalid quantity")
	}

	// User validation
	_, err := service.userService.GetUserById(ctx, &pbusers.GetUserRequest{
		Id: request.GetUser(),
	})
	if err != nil {
		return nil, err
	}

	// Product availability validation
	_, err = service.productService.ValidateProductAvailability(ctx, &pbproducts.ValidateProductAvailabilityRequest{
		Id:       request.GetProduct(),
		Quantity: request.GetQuantity(),
	})
	if err != nil {
		return nil, err
	}

	// Update product
	_, err = service.productService.UpdateProductAvailablesValue(ctx, &pbproducts.UpdateProductAvailablesValueRequest{
		Id:         request.GetProduct(),
		ValueToAdd: -int32(request.GetQuantity()),
	})
	if err != nil {
		return nil, err
	}

	// Create order
	return service.repository.CreateOrder(request)
}

func (service *OrderService) ListUserOrders(ctx context.Context, request *pb.ListUserOrdersRequest) (*pb.ListUserOrdersResponse, error) {
	return service.repository.ListUserOrders(request)
}
