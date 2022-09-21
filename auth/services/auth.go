package services

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	pb "github.com/ruancaetano/grpc-graphql-store/auth/pbauth"
	"github.com/ruancaetano/grpc-graphql-store/users/clients"
	pbusers "github.com/ruancaetano/grpc-graphql-store/users/pbusers"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	userService *clients.UserServiceClient
}

func (service *AuthService) Authenticate(ctx context.Context, req *pb.AuthenticationRequest) (*pb.AuthenticationResponse, error) {
	user, err := service.userService.GetUserByCredentials(ctx, &pbusers.GetUserByCredentialsRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}

	return &pb.AuthenticationResponse{
		Token: token,
	}, nil
}

func generateToken(user *pbusers.User) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["email"] = user.GetEmail()
	claims["id"] = user.GetId()
	claims["name"] = user.GetName()

	return token.SignedString(secret)
}

func NewAuthService(userServiceClient *clients.UserServiceClient) *AuthService {
	return &AuthService{
		userService: userServiceClient,
	}
}
