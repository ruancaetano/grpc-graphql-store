package clients

import (
	"context"
	"log"

	"github.com/ruancaetano/grpc-graphql-store/auth/interceptors"
	pb "github.com/ruancaetano/grpc-graphql-store/users/pbusers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	conn pb.UserServiceClient
}

func (client *UserServiceClient) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	return client.conn.CreateUser(ctx, in)
}

func (client *UserServiceClient) GetUserById(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	return client.conn.GetUserById(ctx, in)
}

func (client *UserServiceClient) GetUserByCredentials(ctx context.Context, in *pb.GetUserByCredentialsRequest) (*pb.User, error) {
	return client.conn.GetUserByCredentials(ctx, in)
}

func NewUserServiceClient(url string) *UserServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.UnaryAuthClientInterceptor()),
	}
	connection, err := grpc.Dial(url, opts...)
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server %v", err)
	}

	return &UserServiceClient{
		conn: pb.NewUserServiceClient(connection),
	}
}
