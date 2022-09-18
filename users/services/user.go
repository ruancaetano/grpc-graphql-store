package services

import (
	"context"
	"net/mail"

	"golang.org/x/crypto/bcrypt"

	"github.com/ruancaetano/grpc-graphql-store/users/pb"
	"github.com/ruancaetano/grpc-graphql-store/users/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
	pb.UnimplementedUserServiceServer
}

func (service *UserService) CreateUser(contexts context.Context, user *pb.CreateUserRequest) (*pb.User, error) {
	_, err := mail.ParseAddress(user.GetEmail())

	if err != nil {
		return nil, err
	}

	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), 8)

	if err != nil {
		return nil, err
	}

	createdUser, err := service.repository.InsertUser(user.GetName(), user.GetEmail(), string(hashedPassowrd))

	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (service *UserService) GetUserById(contexts context.Context, request *pb.GetUserRequest) (*pb.User, error) {
	createdUser, err := service.repository.GetUserById(request.GetId())

	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (service *UserService) GetUserByCredentials(contexts context.Context, request *pb.GetUserByCredentialsRequest) (*pb.User, error) {
	foundUser, err := service.repository.GetUserByEmail(request.GetEmail())

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.GetPassword()), []byte(request.GetPassword()))

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:    foundUser.GetId(),
		Name:  foundUser.GetName(),
		Email: foundUser.GetEmail(),
	}, nil
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		repository,
		pb.UnimplementedUserServiceServer{},
	}
}
