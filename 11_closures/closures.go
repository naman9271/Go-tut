package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// Closures are functions that "close over" variables from their surrounding scope.
	// They can capture and remember the environment in which they were created.

	// Example of a closure that captures a variable from its surrounding scope
	increment := counter()
	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
	fmt.Println(increment()) // Output: 3
}
