package main

import (
	"fmt"
	"time"
)

func main() {

	i := 5

	switch i {
	case 1:
		fmt.Println("first line")
	case 2:
		fmt.Println("second line")
	default:
		fmt.Println("other")
	}

	// multiple switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is a weekend")
	default:
		fmt.Println("it is a weekday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("it is an integer")
		case string:
			fmt.Println("it is a string")
		case bool:
			fmt.Println("it is a bool")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI("hello")
	whoAmI(10)
	whoAmI(false)
	whoAmI(12.42)
}
