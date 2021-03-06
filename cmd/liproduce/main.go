package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SatvikR/liproduce/pkg/database"
	"github.com/SatvikR/liproduce/pkg/gen"
	"github.com/SatvikR/liproduce/pkg/gen/generated"
	"github.com/SatvikR/liproduce/pkg/utils"
	_ "github.com/joho/godotenv/autoload"
)

const defaultPort = "8080"

func main() {
	database.CreateConnection()
	utils.InitHash()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &gen.Resolver{}}))

	if os.Getenv("ENVIRONMENT") != "production" {
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}

	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
