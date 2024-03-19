package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string	`json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo {
		Text: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	jsonText, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonText, 0644)
}