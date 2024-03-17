package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var theUser User	
	theUser = User {
		"Leroy",
		"Royal",
		"01/01/2002",
		time.Now(),
	}

	var anotherUser User = User {
		firstName : userFirstName,
		lastName : userLastName,
		birthDate : userBirthdate,
		createdAt : time.Now(),
	}

	outputUserData(theUser)
	outputUserData(anotherUser)

}

func outputUserData(u User) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}