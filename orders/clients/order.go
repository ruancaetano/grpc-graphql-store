package clients

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ruancaetano/grpc-graphql-store/auth/interceptors"
	pb "github.com/ruancaetano/grpc-graphql-store/orders/pborders"
)

type OrderServiceClient struct {
	conn pb.OrderServiceClient
}

func (client *OrderServiceClient) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.Order, error) {
	return client.conn.CreateOrder(ctx, in)
}

func (client *OrderServiceClient) ListUserOrders(ctx context.Context, in *pb.ListUserOrdersRequest) (*pb.ListUserOrdersResponse, error) {
	return client.conn.ListUserOrders(ctx, in)
}

func NewOrderServiceClient(url string) *OrderServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.UnaryAuthClientInterceptor()),
	}
	connection, err := grpc.Dial(url, opts...)
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server %v", err)
	}

	return &OrderServiceClient{
		conn: pb.NewOrderServiceClient(connection),
	}
}
