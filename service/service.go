package service

import (
	"errors"

	"abimanyu.dev/todo-go/model"
	"abimanyu.dev/todo-go/temp"
)

func GetById(id int) (*model.Todo, error) {
	for i, t := range temp.Todos {
		if t.Id == id {
			return &temp.Todos[i], nil
		}
	}

	return nil, errors.New("to-do not found")
}

func FindIndex(id int) int {
	for i, t := range temp.Todos {
		if t.Id == id {
			return i
		}
	}

	return -1
}