package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/ruancaetano/grpc-graphql-store/bff-graphql/graph"
	"github.com/ruancaetano/grpc-graphql-store/bff-graphql/graph/generated"

	corders "github.com/ruancaetano/grpc-graphql-store/orders/clients"
	cproducts "github.com/ruancaetano/grpc-graphql-store/products/clients"
	cusers "github.com/ruancaetano/grpc-graphql-store/users/clients"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	resolver := &graph.Resolver{
		UserServiceClient:    cusers.NewUserServiceClient(os.Getenv("USERS_SERVICE_URL")),
		ProductServiceClient: cproducts.NewProductServiceClient(os.Getenv("PRODUCTS_SERVICE_URL")),
		OrderServiceClient:   corders.NewOrderServiceClient(os.Getenv("ORDERS_SERVICE_URL")),
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
