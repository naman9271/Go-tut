package main

import "fmt"

func main() {
	var nums [4]int

	fmt.Println(len(nums))

	nums[0] = 10
	fmt.Println(nums[0])
	fmt.Println(nums)

	var vals [4]bool
	fmt.Println(len(vals))
	fmt.Println(vals)

	var name [4]string
	fmt.Println(len(name))
	fmt.Println(name)

	num := [4]int{1, 2, 3, 4}
	fmt.Println(num)

	//2d array
	// array2d := [2][3]int{{1, 2, 3}, {6, 7, 8}}
	var array2d = [2][3]int{{1, 2, 3}, {6, 7, 8}}
	fmt.Println(array2d)

}
