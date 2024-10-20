package main

import (
	"fmt"

	"example.com/structs/user"
)



func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.PrintUser()
	appUser.ClearUserName()
	outputUserDetails(appUser)

	beAdmin := user.NewAdmain(appUser, "admin@admin.com.tw", "wasd")
	// fmt.Println(*beAdmin)
	beAdmin.PrintUser()
}

func outputUserDetails(u *user.User) {
	// for struct, it's equal to (*u).firstName...
	// fmt.Println(u.firstName, u.lastName, u.birthDate)

	// there is not work for imported struct
	fmt.Println(*u)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
