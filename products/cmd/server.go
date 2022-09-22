package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/ruancaetano/grpc-graphql-store/auth/interceptors"
	"github.com/ruancaetano/grpc-graphql-store/products/db"
	"github.com/ruancaetano/grpc-graphql-store/products/repositories"
	"github.com/ruancaetano/grpc-graphql-store/products/services"

	pb "github.com/ruancaetano/grpc-graphql-store/products/pbproducts"

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

	mappedMethodsAndRoles := map[string][]string{
		"/pbproducts.ProductService/ListProducts":                 {"user", "admin"},
		"/pbproducts.ProductService/GetProductById":               {"user", "admin"},
		"/pbproducts.ProductService/ValidateProductAvailability":  {"user", "admin"},
		"/pbproducts.ProductService/UpdateProductAvailablesValue": {"user", "admin"},
		"/pbproducts.ProductService/CreateProduct":                {"admin"},
		"/pbproducts.ProductService/UpdateProduct":                {"admin"},
		"/pbproducts.ProductService/DeleteProduct":                {"admin"},
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptors.UnaryAuthServerInterceptor(mappedMethodsAndRoles)),
	}
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
