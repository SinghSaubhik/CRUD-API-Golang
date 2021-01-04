package controller

import (
	"todo-app/dao"
	"todo-app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateTodoInput struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

func GetTodos(context *gin.Context) {
	todos := dao.FindAll()
	context.JSON(200, gin.H{
		"todos": todos,
		"error": nil,
	})
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := dao.FindOne(id)
	if err != nil {
		context.JSON(404, gin.H{
			"todo":  nil,
			"error": err,
		})
	} else {
		context.JSON(200, gin.H{
			"todo":  todo,
			"error": nil,
		})
	}

}

func DeleteTodo(context *gin.Context) {
	id := context.Param("id")
	deletedDoc, err := dao.FindOneAndDelete(id)
	if err != nil {
		context.JSON(404, gin.H{
			"error": err.Error(),
			"todo":  nil,
		})
		return
	}

	context.JSON(200, gin.H{
		"error": nil,
		"todo":  deletedDoc,
	})
}

func AddTodo(context *gin.Context) {
	id := uuid.New()
	var input CreateTodoInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo := models.Todo{ID: id.String(), Title: input.Title, Completed: input.Completed}
	_ = dao.InsertOne(todo)
	context.JSON(201, gin.H{
		"message": "Successfully created",
	})
	return
}

func UpdateTodo(context *gin.Context) {
	id := context.Param("id")
	var input CreateTodoInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	todo := models.Todo{ID: id, Title: input.Title, Completed: input.Completed}

	updatedDoc, err := dao.FindOneAndUpdate(id, todo)
	if err != nil {
		context.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"error": nil,
		"todo":  updatedDoc,
	})

}
