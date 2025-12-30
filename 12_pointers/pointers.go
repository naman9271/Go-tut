package main

// Pointers in Go
// A pointer is a variable that stores the memory address of another variable.
// Pointers are used to directly access and manipulate memory locations.

import "fmt"

// func chanageValue(num int) {
// 	num = 20 // dereferencing pointer to change the value at the address
// 	fmt.Println("Inside chanageValue, val:", num)
// }

func changeValueUsingPointer(numPtr *int) {
	*numPtr = 20 // dereferencing pointer to change the value at the address
	fmt.Println("Inside changeValueUsingPointer, val:", *numPtr)
}

func main() {
	num := 10

	// fmt.Println("Before chanageValue, num:", num)
	// chanageValue(num)
	// fmt.Println("After chanageValue, num:", num)

	// the above will not change the original value because num is passed by value

	// Using pointers to change the original value
	fmt.Println("Before changeValueUsingPointer, num:", num)
	changeValueUsingPointer(&num) // passing address of num
	fmt.Println("After changeValueUsingPointer, num:", num)

}
