package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/keviinliuu/leetlist/auth"
	"github.com/keviinliuu/leetlist/database"
	"github.com/keviinliuu/leetlist/graph"
	"github.com/keviinliuu/leetlist/graph/resolvers"
	"github.com/keviinliuu/leetlist/util"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbConfig := &database.Config {
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}
	
	db, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err) 
	}

	database.AutoMigrate(db)

	browser := util.InitBrowser()

	resolver := &resolvers.Resolver{
		DB: db,
		Browser: browser,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", auth.Middleware(srv))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})

	log.Printf("Starting server on port %s\n", port)
	err = http.ListenAndServe(":"+port, c.Handler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}
