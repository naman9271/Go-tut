package main

import "fmt"

// Interfaces in Go
// An interface is a type that defines a set of method signatures.
// A struct implements an interface by implementing its methods.

type paymenter interface {
	pay(amount float64)
}

type payment struct {
	gateway paymenter
}

func (p payment) processPayment(amount float64) {
	// razorpayment := razorpay{}
	// razorpayment.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float64) {
	fmt.Printf("Processing Razorpay payment of $%.2f\n", amount)
}

type stripe struct{}

func (s stripe) pay(amount float64) {
	fmt.Printf("Processing Stripe payment of $%.2f\n", amount)
}

func main() {
	// stripePaymentGw := stripe{}
	razorpayPaymentGw := razorpay{}
	newPayment := payment{gateway: razorpayPaymentGw}
	newPayment.processPayment(150.75)
}
