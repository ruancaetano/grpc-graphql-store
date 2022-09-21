package clients

import (
	"context"
	"log"

	pb "github.com/ruancaetano/grpc-graphql-store/orders/pborders"
	"google.golang.org/grpc"
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
	connection, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server %v", err)
	}

	return &OrderServiceClient{
		conn: pb.NewOrderServiceClient(connection),
	}
}
