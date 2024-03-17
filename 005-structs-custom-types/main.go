package main

import (
	"fmt"

	"leroy.africa/leroy/structs-custom-types/user"
)



func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var theUser *user.User	
	theUser, err := user.New("Leroy", "Royal", "01/01/2002")
	anotherUser, err := user.New(userFirstName, userLastName, userBirthdate)

	admin, err := user.NewAdmin("Leroy", "leroy@leroy.com", "the password")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	theUser.OutputUserData()
	anotherUser.OutputUserData()
	theUser.ClearUsername()
	anotherUser.ClearUsername()
	anotherUser.OutputUserData()
	theUser.OutputUserData()
	admin.OutputUserData()

}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}