package clients

import (
	"context"
	"log"

	"github.com/ruancaetano/grpc-graphql-store/auth/interceptors"
	pb "github.com/ruancaetano/grpc-graphql-store/auth/pbauth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	conn pb.AuthServiceClient
}

func (client *AuthServiceClient) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	return client.conn.SignIn(ctx, in)
}

func NewAuthServiceClient(url string) *AuthServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.UnaryAuthClientInterceptor()),
	}
	connection, err := grpc.Dial(url, opts...)
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server %v", err)
	}

	return &AuthServiceClient{
		conn: pb.NewAuthServiceClient(connection),
	}
}
