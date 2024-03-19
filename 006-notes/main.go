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

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	// displayer
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {

	printAnything(1)
	printAnything("hello")
	printAnything(true)
	printAnything(3.14)
	printAnything(nil)
	printAnything([]int{1, 2, 3, 4, 5})


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

	err = outputData(newNote)
	if err != nil {
		fmt.Println(err)
		return 
	}

	err = outputData(newTodo)
	if err != nil {
		fmt.Println(err)
		return 
	}
}

func printAnything(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Interger: ", value)
	case string:
		fmt.Println("String: ", value)
	case bool:
		fmt.Println("Boolean: ", value)
	case float64:
		fmt.Println("Float: ", value)
	default:
		fmt.Println("Unknown type")
	}

	// fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if (err != nil) {
		fmt.Printf("error saving : %v\n",err)
		return err
	}

	fmt.Println("saved successfully")
	return nil
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