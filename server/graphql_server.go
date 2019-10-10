package main

import (
	"github.com/99designs/gqlgen/handler"
	"log"
	"net/http"
	"os"
	"self_gqlgen/gqlgen"
	"self_gqlgen/resolvers"
)

const defaultPort = "8888"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &resolvers.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
