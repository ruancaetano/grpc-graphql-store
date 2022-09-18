package clients

import (
	"context"
	"log"

	"github.com/ruancaetano/grpc-graphql-store/users/pb"
	"google.golang.org/grpc"
)

type UserServiceClient struct {
	Conn *grpc.ClientConn
}

func (client *UserServiceClient) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	return client.CreateUser(ctx, in)
}

func (client *UserServiceClient) GetUserById(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	return client.GetUserById(ctx, in)
}

func (client *UserServiceClient) GetUserByCredentials(ctx context.Context, in *pb.GetUserByCredentialsRequest) (*pb.User, error) {
	return client.GetUserByCredentials(ctx, in)
}

func NewUserServiceClient(url string) *UserServiceClient {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server %v", err)
	}

	return &UserServiceClient{
		Conn: connection,
	}
}
