package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/ruancaetano/grpc-graphql-store/bff-graphql/graph"
	"github.com/ruancaetano/grpc-graphql-store/bff-graphql/graph/generated"
	"github.com/ruancaetano/grpc-graphql-store/shared/clients"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	resolver := &graph.Resolver{
		UserServiceClient: clients.NewUserServiceClient(os.Getenv("USERS_SERVICE_URL")),
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
