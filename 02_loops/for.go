package main

import "fmt"

func main() {

	//while loop
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// infinit loop
	// for {
	// fmt.Println("1")
	// }

	// for loop
	for i := 0; i < 5; i++ {
		println("Iteration:", i)
	}

	// for range loop

	for i := range 3 {
		fmt.Println(i)
	}
}
