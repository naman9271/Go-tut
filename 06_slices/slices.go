package main

import (
	"fmt"
	"slices"
)

// slices are dynamic arrays

func main() {

	// var s []int // nil slice
	// fmt.Println(s)
	// fmt.Println(len(s)) // current size
	// fmt.Println(cap(s)) //max size before resize

	// var nums = make([]int, 3, 5) // len=3, cap=5
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// nums = append(nums, 10)
	// nums = append(nums, 20)
	// fmt.Println(nums)

	// //copy function
	// nums := make([]int, 2, 5)
	// nums[0] = 1
	// nums2 := make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	//slice operator
	nums := []int{1, 2, 3}
	fmt.Println(nums[0:2])

	// slcie package functions

	nums1 := []int{1, 2}
	nums2 := []int{1, 3}

	fmt.Println(slices.Equal(nums1, nums2))

	//2d slices
	slice2d := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(slice2d)
}
