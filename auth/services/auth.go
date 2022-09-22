package services

import (
	"context"

	pb "github.com/ruancaetano/grpc-graphql-store/auth/pbauth"
	"github.com/ruancaetano/grpc-graphql-store/auth/utils"
	"github.com/ruancaetano/grpc-graphql-store/users/clients"
	pbusers "github.com/ruancaetano/grpc-graphql-store/users/pbusers"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	userService *clients.UserServiceClient
}

func (service *AuthService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	user, err := service.userService.GetUserByCredentials(ctx, &pbusers.GetUserByCredentialsRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateJwtUserToken(user)
	if err != nil {
		return nil, err
	}

	return &pb.SignInResponse{
		Token: token,
	}, nil
}

func NewAuthService(userServiceClient *clients.UserServiceClient) *AuthService {
	return &AuthService{
		userService: userServiceClient,
	}
}
