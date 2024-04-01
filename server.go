package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/NaturalBug/patients-api/graph"
	"github.com/NaturalBug/patients-api/postgres"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"
const (
	HOST     = "localhost"
	USER     = "admin"
	PASSWORD = "12345"
	DATABASE = "patients"
	PORT     = "1234"
)

var db *postgres.Db

func main() {
	// Initialize connection object.
	db, err := postgres.Connect(postgres.BuildConnectionString("localhost", 1234, "admin", "12345", "patient"))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
