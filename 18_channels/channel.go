package main

import (
	"fmt"
	"time"
	// "time"
	// "math/rand"
)

// sending data to channel
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing num", num)
		time.Sleep(time.Second)
	}

	// here go do somethjing internally like this while using range we don't need to use <- operator
	//for {
	// num, ok := <-numChan   // receive
	// if !ok {               // channel closed?
	// break
	// }
	// ...use num...
	// }

}

// receiving data from channel
func sum(resultChan chan int, num1 int, num2 int) {
	result := num1 + num2
	resultChan <- result
}

// doing the thing concurrently without waitgroup
func task(done chan bool) {
	defer func() { done <- true }() // notify that task is done
	fmt.Println("Task started")
}

func main() {

	//buffered channel -> channel with capacity

	emailChan := make(chan string, 2) // buffered channel with capacity 2

	emailChan <- "user1@example.com"
	emailChan <- "user2@example.com"

	// The following line would block because the channel is full
	// emailChan <- "user3@example.com"
	close(emailChan) // close the channel when done sending

	// receiving data from buffered channel
	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)

	// unbuffered channel
	// done := make(chan bool)
	//
	// go task(done)
	//
	// <-done // wait for task to complete
	// fmt.Println("Task completed")

	// resultChan := make(chan int)

	// go sum(resultChan, 10, 20)
	// // recieve data from channel
	// result := <-resultChan
	// fmt.Println("Sum:", result)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// messageChan := make(chan string)

	// messageChan <- "Hello, Channel!"
	// msg := <-messageChan

	// fmt.Println(msg)
}
