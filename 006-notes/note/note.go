package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string `json:"title"`
	Content string	`json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note {
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Note title: %v\nContent: \n\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "-")
	fileName = strings.ToLower(fileName) + ".json"

	jsonContent, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0644)
}