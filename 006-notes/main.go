package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

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
	
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if text == "" || err != nil{
		return "", errors.New("value is empty")
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text, nil
}