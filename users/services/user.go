package services

import (
	"context"

	"github.com/ruancaetano/grpc-graphql-store/users/pb"
	"github.com/ruancaetano/grpc-graphql-store/users/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
	pb.UnimplementedUserServiceServer
}

func (service *UserService) CreateUser(contexts context.Context, user *pb.CreateUserRequest) (*pb.User, error) {
	createdUser, err := service.repository.InsertUser(user.GetName(), user.GetEmail())

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

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		repository,
		pb.UnimplementedUserServiceServer{},
	}
}
