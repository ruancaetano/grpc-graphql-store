package clients

import (
	"context"
	"log"

	pb "github.com/ruancaetano/grpc-graphql-store/users/pbusers"
	"google.golang.org/grpc"
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
	connection, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server %v", err)
	}

	return &UserServiceClient{
		conn: pb.NewUserServiceClient(connection),
	}
}
