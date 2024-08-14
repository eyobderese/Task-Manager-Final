package main

import (
	"context"
	"fmt"
	"log"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/Delivery/router"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// main is the entry point of the program.
// It sets up the router and starts the server on localhost:8080.

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	GeneralRouter := gin.Default()
	router.NewTaskRouter(*client.Database("test"), GeneralRouter.Group("tasks"))
	router.NewUserRouter(*client.Database("test"), GeneralRouter.Group(""))
	GeneralRouter.Run("localhost:8080")

}
