package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/ryuji-cre8ive/huyou-server/db"
	"github.com/ryuji-cre8ive/huyou-server/graph"
	"github.com/ryuji-cre8ive/huyou-server/graph/model"
)

const defaultPort = "8080"

func main() {
	loadEnv()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := db.ConnectGORM()

	db.AutoMigrate(&model.User{},&model.ShopItem{}, &model.Comment{})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080", "https://huyou-ryuji-cre8ive.vercel.app"},
		AllowCredentials: true,
	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("err has occurred: %+v", err)
	}
}
