package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// will also work
func add(a, b int) int {
	return a + b
}

// func to return multiple values
func languages() (string, string, string) {
	return "Go", "Python", "JavaScript"
}

// func accepting functions
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func multiply(x, y int) int {
	return x * y
}

//func returning func

func operation(op string) func(int, int) int {
	switch op {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "multiply":
		return func(a, b int) int {
			return a * b
		}
	}
	return nil
}

func main() {
	sum := add(3, 5)
	fmt.Println("Sum:", sum)

	lang1, lang2, lang3 := languages()
	fmt.Println("Languages:", lang1, lang2, lang3)

	// ignoring values using _
	_, langB, _ := languages()
	fmt.Println("Second Language:", langB)

	// directly print all returned values
	fmt.Println(languages())

	result := operate(4, 6, multiply)
	fmt.Println("Multiplication Result:", result)

	addFunc := operation("add")
	if addFunc != nil {
		fmt.Println("Operation Result (Add):", addFunc(10, 15))
	}

	mulFunc := operation("multiply")
	if mulFunc != nil {
		fmt.Println("Operation Result (Multiply):", mulFunc(10, 15))
	}
}
