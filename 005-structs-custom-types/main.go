package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func (u User) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) clearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name, and birthdate are required")
	}
	return &User {
		firstName : firstName,
		lastName : lastName,
		birthDate : birthDate,
		createdAt : time.Now(),
	}, nil
} 

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var theUser *User	
	theUser, err := newUser("Leroy", "Royal", "01/01/2002")
	anotherUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	theUser.outputUserData()
	anotherUser.outputUserData()
	theUser.clearUsername()
	anotherUser.clearUsername()
	anotherUser.outputUserData()
	theUser.outputUserData()

}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}