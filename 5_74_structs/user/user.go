package user

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
func (u *User) PrintUser() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("not allow empty value")
	}

	return &User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}, nil
}

type Admin struct{
	email string
	password string
	User
}
func NewAdmain(u *User, email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: u.firstName,
			lastName:  u.lastName,
			birthDate: u.birthDate,
			createdAt: u.createdAt,
		},
	}
}