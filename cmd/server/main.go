package main

import (
	"context"
	"fmt"
	"go-web/pkg/api"
	"go-web/pkg/app"
	"go-web/pkg/repository"
	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseUri = "mongodb://localhost:27017"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseUri))
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()

	router.Use(cors.Default())

	storage := repository.NewStorage(client)

	albumService := api.NewAlbumService(storage)

	server := app.NewServer(router, albumService)

	return server.Run()
}
