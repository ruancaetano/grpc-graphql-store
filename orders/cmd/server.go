package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/ruancaetano/grpc-graphql-store/orders/pborders"
	"github.com/ruancaetano/grpc-graphql-store/orders/repositories"
	"github.com/ruancaetano/grpc-graphql-store/orders/services"
	"github.com/ruancaetano/grpc-graphql-store/shared/clients"
	"github.com/ruancaetano/grpc-graphql-store/shared/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterOrderServiceServer(grpcServer, makeOrderService(dbConnection))

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve %v", err)
	}
}

func makeOrderService(dbConnection *sql.DB) *services.OrderService {
	orderRepository := repositories.NewOrderRepository(dbConnection)

	userService := clients.NewUserServiceClient(os.Getenv("USER_SERVICE_URL"))
	productService := clients.NewProductServiceClient(os.Getenv("PRODUCT_SERVICE_URL"))

	return services.NewOrderService(orderRepository, productService, userService)
}