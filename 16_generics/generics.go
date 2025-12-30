package main

import "fmt"

// Generics in Go
// Generics allow you to write flexible and reusable code by defining functions, types, and data structures that can operate on different data types without sacrificing type safety.

// any == interface{} also we can do here [t int | string] for specific types
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}


func main(){
	nums := []int{1, 2, 3, 4, 5}
	strs := []string{"hello", "world", "generics", "in", "go"}

	fmt.Println("Printing integer slice:")
	PrintSlice(nums)

	fmt.Println("Printing string slice:")
	PrintSlice(strs)
}