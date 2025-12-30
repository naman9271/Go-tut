package main

import "fmt"

func main() {

	//maps are key value pairs

	m := make(map[string]int) //empty map
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(m["one"])
	fmt.Println(len(m))

	//delete function
	delete(m, "two")
	fmt.Println(m)

	//clear map
	m = make(map[string]int)
	fmt.Println(m)

	// find value
	val, ok := m["three"]
	fmt.Println(val, ok) //0 false

	if v, ok := m["three"]; ok {
		fmt.Println("value found:", v)
	} else {
		fmt.Println("value not found")
	}

}
