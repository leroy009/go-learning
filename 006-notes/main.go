package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"leroy.africa/leroy/notes/note"
	"leroy.africa/leroy/notes/todo"
)

func main() {
	title, content, err := getNoteData()
	
	if (err != nil) {
		fmt.Println(err)
		return 
	}
	text, err := getTodoData()

	if err != nil {
		fmt.Println(err)
		return 
	}	

	newNote, err := note.New(title, content)
	
	if (err != nil) {
		fmt.Println(err)
		return 
	}
	
	newTodo, err := todo.New(text)

	if err != nil {
		fmt.Println(err)
		return 
	}

	newNote.Display()

	newTodo.Display()

	err = newNote.Save()

	if (err != nil) {
		fmt.Printf("Error saving nore: %v\n",err)
		return 
	}

	err = newTodo.Save()

	if (err != nil) {
		fmt.Printf("Error saving todo: %v\n",err)
		return 
	}

	fmt.Println("Note saved successfully")

	fmt.Println("Todo saved successfully")
	
}

func getTodoData() (text string, err error) {
	text, err = getUserInput("Todo content: ")

	if (err != nil) {
		fmt.Println(err)
		return "", err
	}

	return text, nil
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