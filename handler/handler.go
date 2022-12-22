package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"abimanyu.dev/todo-go/model"
	service "abimanyu.dev/todo-go/service"
	"abimanyu.dev/todo-go/temp"
	"github.com/gin-gonic/gin"
)

func FindOne(context *gin.Context) {
	id := context.Param("id")
	intId, _ := strconv.Atoi(id)

	todo, err := service.GetById(intId)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "To-do not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func FindAll(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, temp.Todos)
}

func Create(context *gin.Context) {
	var newTodo model.Todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	temp.Todos = append(temp.Todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func UpdateTitle(context *gin.Context) {
	id := context.Param("id")

	intId, _ := strconv.Atoi(id)

	todo, err := service.GetById(intId)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "To-do not found"})
		return
	}

	if err := context.BindJSON(&todo); err != nil {
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func ToggleCompleteStatus(context *gin.Context) {
	id := context.Param("id")
	intId, _ := strconv.Atoi(id)

	todo, err := service.GetById(intId)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "To-do not found"})
		return
	}

	todo.IsCompleted = !todo.IsCompleted
	context.IndentedJSON(http.StatusOK, todo)
}

func Delete(context *gin.Context) {
	id := context.Param("id")
	intId, _ := strconv.Atoi(id)

	idx := service.FindIndex(intId)

	if idx == -1 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "To-do not found"})
		return
	}

	temp.Todos = append(temp.Todos[:idx], temp.Todos[idx+1:]...)
	
	context.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successfully deleted to-do with id %d", intId)})
}