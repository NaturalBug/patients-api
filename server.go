package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/NaturalBug/patients-api/graph"
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Initialize connection string.
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s", HOST, USER, PASSWORD, DATABASE, PORT)

	// Initialize connection object.
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")

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
