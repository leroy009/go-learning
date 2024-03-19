package main

import (
	"errors"
	"fmt"

	"leroy.africa/leroy/notes/note"
)

func main() {
	title, content, err := getNoteData()

	if (err != nil) {
		fmt.Println(err)
		return 
	}

	newNote, err := note.New(title, content)

	if (err != nil) {
		fmt.Println(err)
		return 
	}

	newNote.Display()
	
}

func getNoteData() (title string, content string, err error) {
	title, err = getUserInput("Note title: ")
	if (err != nil) {
		fmt.Println(err)
		return "", "", err
	}
	content, err = getUserInput("Note content: ")

	if (err != nil) {
		fmt.Println(err)
		return "", "", err
	}

	return title, content, nil
}


func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scan(value)

	if value == "" {
		return "", errors.New("value is empty")
	}

	return value, nil
}