package main

import (
	"context"
	"log"
	"todo-app/controller"
	"todo-app/dao"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	router := gin.Default()

	// Connect to DB
	dbClient, ctx := dao.ConnectDB()

	// Routes
	router.GET("/todos", controller.GetTodos)
	router.GET("/todo/:id", controller.GetTodo)

	router.DELETE("/todo/:id", controller.DeleteTodo)

	router.POST("/todos", controller.AddTodo)
	router.PUT("/todo/:id", controller.UpdateTodo)

	if err := router.Run(); err != nil {
		log.Fatal("Unable to start server!!!", err)
	}

	// Close DB connection when application is exiting
	defer disconnectDB(dbClient, ctx)
}

func disconnectDB(client *mongo.Client, ctx context.Context) {
	println("Disconnected")
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
