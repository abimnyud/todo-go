package main

import (
	"abimanyu.dev/todo-go/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", handler.FindAll)
	router.GET("/todo/:id", handler.FindOne)
	router.PATCH("/todo/:id/edit-title", handler.UpdateTitle)
	router.PATCH("/todo/:id/toggle-status", handler.ToggleCompleteStatus)
	router.POST("/todo", handler.Create)
	router.DELETE("/todo/:id", handler.Delete)

	router.Run("localhost:8080")
}