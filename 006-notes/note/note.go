package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title string
	content string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note {
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Note title: %v\nContent: \n\n%v", note.title, note.content)
}

func (note Note) Save() {
	fileName := strings.ReplaceAll(note.title, " ", "-")
	fileName = strings.ToLower(fileName)
	os.WriteFile(fileName, []byte(note.content), 0644)
}