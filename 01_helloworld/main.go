package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	//simple values
	// int
	fmt.Println(1 + 1)

	// string
	fmt.Println("Hello" + " " + "GoLang")

	// bool
	fmt.Println(true && false)

	// float
	fmt.Println(3.14 * 2)

	// variables
	var name string = "Gopher"
	// type inference // short hand
	age := 5
	isCool := true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Cool:", isCool)

	//constants	const pi = 3.14159
	const petName = "Gopher"

	// petName="New Gopher" // This will cause an error
	fmt.Println("Constant Name:", petName)

}
