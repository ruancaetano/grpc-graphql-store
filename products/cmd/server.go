package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/ruancaetano/grpc-graphql-store/products/repositories"
	"github.com/ruancaetano/grpc-graphql-store/products/services"
	"github.com/ruancaetano/grpc-graphql-store/shared/db"

	pb "github.com/ruancaetano/grpc-graphql-store/products/pbproducts"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	flag.Parse()

	err := godotenv.Load("./products/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Listen port %s", port)

	dbConnection := db.NewDbConnection()
	defer dbConnection.Close()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterProductServiceServer(grpcServer, makeProductService(dbConnection))

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve %v", err)
	}
}

func makeProductService(dbConnection *sql.DB) *services.ProductService {
	productRepository := repositories.NewProductRepository(dbConnection)
	return services.NewProductService(productRepository)
}
