package main

import "fmt"

// variadic function for adding numbers
func add(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// variadic function are those which can take variable number of arguments
func main() {
	fmt.Println(1, 2, 4, 5, 6, "hello") // Println is variadic function

	// using our variadic add function
	sum1 := add(1, 2)
	sum2 := add(1, 2, 3, 4, 5)
	fmt.Println("Sum1:", sum1)
	fmt.Println("Sum2:", sum2)

	// passing slice to variadic function
	numbers := []int{10, 20, 30, 40}
	sum3 := add(numbers...) // unpacking slice
	fmt.Println("Sum3:", sum3)

}
