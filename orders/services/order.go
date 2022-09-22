package services

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/ruancaetano/grpc-graphql-store/orders/pborders"
	pbproducts "github.com/ruancaetano/grpc-graphql-store/products/pbproducts"
	pbusers "github.com/ruancaetano/grpc-graphql-store/users/pbusers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	cproducts "github.com/ruancaetano/grpc-graphql-store/products/clients"
	cusers "github.com/ruancaetano/grpc-graphql-store/users/clients"

	"github.com/ruancaetano/grpc-graphql-store/orders/repositories"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	repository     *repositories.OrderRepository
	productService *cproducts.ProductServiceClient
	userService    *cusers.UserServiceClient
}

func NewOrderService(repository *repositories.OrderRepository, productService *cproducts.ProductServiceClient, userService *cusers.UserServiceClient) *OrderService {
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

	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(userId)
	// User validation
	user, err := service.userService.GetUserById(ctx, &pbusers.GetUserRequest{
		Id: userId,
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
	return service.repository.CreateOrder(user.GetId(), request)
}

func (service *OrderService) ListUserOrders(ctx context.Context, request *pb.Empty) (*pb.ListUserOrdersResponse, error) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return service.repository.ListUserOrders(userId)
}

func getUserIdFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromOutgoingContext(ctx)

	if !ok {
		return "", status.Error(codes.InvalidArgument, "Cannot get user id from token")
	}

	userIdMetadata := md.Get("userId")

	if len(userIdMetadata) == 0 {
		return "", status.Error(codes.InvalidArgument, "Cannot get user id from token")
	}

	return userIdMetadata[0], nil
}
