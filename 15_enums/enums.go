package main

import "fmt"

// Enums in Go
// Go does not have a built-in enum type like some other languages.
// However, you can achieve similar functionality using constants and iota.

// type OrderStatus int

// const (
// 	Pending OrderStatus = iota
// 	Shipped
// 	Delivered
// 	Cancelled
// )

// for strings
type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Shipped               = "Shipped"
	Delivered             = "Delivered"
	Cancelled             = "Cancelled"
)

func changeStatus(status OrderStatus) {
	fmt.Println("chnaging order staus to ", status)
}

func main() {
	changeStatus(Shipped)
}
