package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func (u User) OutputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name, and birthdate are required for a user")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(firstName, email, password string) (*Admin, error) {
	if email == "" || password == ""{
		return nil, errors.New("email and password are required for an admin")
	}
	return &Admin{
		email: email,
		password: password,
		User: User {
			firstName: firstName,
			lastName: "Admin",
			birthDate: " **/ **/ 19**",
			createdAt: time.Now(),
		},
	}, nil
}