package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ruancaetano/grpc-graphql-store/auth/interceptors"
	"github.com/ruancaetano/grpc-graphql-store/users/db"
	pb "github.com/ruancaetano/grpc-graphql-store/users/pbusers"
	"github.com/ruancaetano/grpc-graphql-store/users/repositories"
	"github.com/ruancaetano/grpc-graphql-store/users/services"
)

func main() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Listen port %s", port)

	dbConnection := db.NewDbConnection()
	defer dbConnection.Close()

	mappedMethodsAndRoles := map[string][]string{
		"/pbusers.UserService/CreateUser":           {"public"},
		"/pbusers.UserService/GetUserById":          {"user", "admin"},
		"/pbusers.UserService/GetUserByCredentials": {"public"},
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptors.UnaryAuthServerInterceptor(mappedMethodsAndRoles)),
	}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterUserServiceServer(grpcServer, makeUserService(dbConnection))

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve %v", err)
	}
}

func makeUserService(dbConnection *sql.DB) *services.UserService {
	userRepository := repositories.NewUserRepository(dbConnection)
	return services.NewUserService(userRepository)
}
