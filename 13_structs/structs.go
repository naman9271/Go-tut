package main

import (
	"fmt"
	"time"
)

// Structs in Go
// A struct is a composite data type that groups together variables under a single name.
// Each variable in a struct is called a field.

type Customer struct {
	name string
	age  int
}

type Order struct {
	id        int
	amount    float64
	status    string
	createdAt time.Time
	Customer  //struct embedding
}

func NewOrder(id int, amount float64, status string) *Order {
	myOrder := Order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// have to use pointer receiver to modify the original struct
func (o *Order) changeStatus(newStatus string) {
	o.status = newStatus
}

// pointer is not needed here as we are just reading the value
func getAmount(o Order) float64 {
	return o.amount
}

func main() {
	// Creating an instance of Order struct
	order := Order{
		id:     1,
		amount: 99.99,
		status: "Pending",
		Customer: Customer{ //struct embedding
			name: "John Doe",
			age:  30,
		},
	}
	fmt.Println(order)
	order.createdAt = time.Now()

	order.changeStatus("Shipped")
	amount := getAmount(order)
	fmt.Println("Order Amount:", amount)
	// // Accessing struct fields
	// fmt.Println("Order ID:", order.id)
	// fmt.Println("Amount:", order.amount)
	// fmt.Println("Status:", order.status)
	// fmt.Println("Created At:", order.createdAt)

	order2 := NewOrder(2, 49.99, "Processing")
	fmt.Println(order2)

	//struct without initialization
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}
