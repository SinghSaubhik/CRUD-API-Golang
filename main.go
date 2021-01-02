package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"todo-app/controller"
	"todo-app/dao"
)

func main() {
	router := gin.Default()

	// Connect to DB
	dbClient, ctx := dao.ConnectDB()

	// Routes
	router.GET("/todos", controller.GetTodos)
	router.POST("/todos", controller.AddTodo)
	router.GET("/todo/:id", controller.GetTodo)

	if err := router.Run(); err != nil {
		log.Fatal("Unable to start server!!!", err)
	}

	defer disconnectDB(dbClient, ctx)
}

func disconnectDB(client *mongo.Client, ctx context.Context) {
	println("Disconnected")
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
