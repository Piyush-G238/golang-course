package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"content"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {

	fileName := "todo.json"

	jsonData, error := json.Marshal(todo)

	if error != nil {
		return error
	}

	os.WriteFile(fileName, jsonData, 0644)
	return nil
}

func NewTodo(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("content is mandatory to create todo")
	}
	return Todo{Text: content}, nil
}
