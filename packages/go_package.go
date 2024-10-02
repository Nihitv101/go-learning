package main

import (
	"fmt"

	"github.com/Nihitv101/packup/auth"
	"github.com/Nihitv101/packup/user"
	"github.com/fatih/color"
)

// Packages - reusable code

func main() {

	auth.LoginWithCredentials("nihit", "123")
	session := auth.GetSession()
	fmt.Println(session)

	user1 := user.User{
		Email: "user@gmail.com",
		Name:  "john doe",
	}

	// fmt.Println(user1.Email, user1.Name)
	color.Red(user1.Email)
	color.Blue(user1.Name)

}
