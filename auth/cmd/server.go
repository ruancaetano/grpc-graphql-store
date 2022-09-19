package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/ruancaetano/grpc-graphql-store/auth/pbauth"
	"github.com/ruancaetano/grpc-graphql-store/auth/services"
	"github.com/ruancaetano/grpc-graphql-store/shared/clients"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	flag.Parse()

	err := godotenv.Load("./auth/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Listen port %s", port)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	userServiceClient := clients.NewUserServiceClient(os.Getenv("USER_SERVICE_URL"))

	pb.RegisterAuthServiceServer(grpcServer, services.NewAuthService(userServiceClient))

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve %v", err)
	}
}
