package model

import "main/entity"

type Todo struct {
	Ref   string `json:"ref"`
	User  string `json:"user"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func NewTodoFromEntity(todo *entity.Todo, userName string) *Todo {
	return &Todo{
		Ref:   todo.Ref,
		User:  userName,
		Title: todo.Title,
		Done:  todo.Done,
	}
}
