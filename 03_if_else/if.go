package main

import "fmt"

func main() {

	age := 16

	if age >= 18 {
		fmt.Println("you are an adult")
	} else if age >= 12 {
		fmt.Println("a teenager")
	} else {
		fmt.Println("not an adult")
	}

}
