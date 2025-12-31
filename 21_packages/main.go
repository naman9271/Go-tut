package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/naman9271/Go-tut/auth"
	"github.com/naman9271/Go-tut/user"
)

// Everything has first letter capitalized to be accessible from other packages
// Functions, Structs, Variables, Constants
// if small then it will be private to that package only

func main() {
	auth.LoginWithCredentials("hello", "naman")
	session := auth.GetSession()
	fmt.Println("Session ID:", session)

	user := user.User{
		Email: "hello@gmail.com",
		Name:  "Naman",
	}
	fmt.Println("User Name:", user.Name)
	fmt.Println("User Email:", user.Email)

	// colorful terminal output
	color.Red("This is red color text")
	c := color.New(color.BgCyan)
	c.Println("This has cyan background")
}
