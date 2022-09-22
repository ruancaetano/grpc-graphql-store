package clients

import (
	"context"
	"log"

	"github.com/ruancaetano/grpc-graphql-store/auth/interceptors"
	pb "github.com/ruancaetano/grpc-graphql-store/products/pbproducts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	conn pb.ProductServiceClient
}

func (client *ProductServiceClient) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.Product, error) {
	return client.conn.CreateProduct(ctx, in)
}

func (client *ProductServiceClient) UpdateProduct(ctx context.Context, in *pb.UpdateProductRequest) (*pb.Product, error) {
	return client.conn.UpdateProduct(ctx, in)
}

func (client *ProductServiceClient) DeleteProduct(ctx context.Context, in *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	return client.conn.DeleteProduct(ctx, in)
}

func (client *ProductServiceClient) ListProducts(ctx context.Context, in *pb.PaginationParams) (*pb.ProductListResponse, error) {
	return client.conn.ListProducts(ctx, in)
}

func (client *ProductServiceClient) GetProductById(ctx context.Context, in *pb.GetProductByIdRequest) (*pb.Product, error) {
	return client.conn.GetProductById(ctx, in)
}

func (client *ProductServiceClient) ValidateProductAvailability(ctx context.Context, in *pb.ValidateProductAvailabilityRequest) (*pb.ValidateProductAvailabilityResponse, error) {
	return client.conn.ValidateProductAvailability(ctx, in)
}

func (client *ProductServiceClient) UpdateProductAvailablesValue(ctx context.Context, in *pb.UpdateProductAvailablesValueRequest) (*pb.Product, error) {
	return client.conn.UpdateProductAvailablesValue(ctx, in)
}

func NewProductServiceClient(url string) *ProductServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.UnaryAuthClientInterceptor()),
	}
	connection, err := grpc.Dial(url, opts...)
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server %v", err)
	}

	return &ProductServiceClient{
		conn: pb.NewProductServiceClient(connection),
	}
}
