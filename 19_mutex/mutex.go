package main

import (
	"fmt"
	"sync"
)

// Mutexes in Go
// A mutex is a mutual exclusion lock that can be used to protect shared data from concurrent access.
// The sync package provides the Mutex type for this purpose.

// A mutex has two main methods:
// Lock()   // locks the mutex
// Unlock() // unlocks the mutex

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.views++
	// p.mu.Unlock() -> it is better to use defer for unlocking to avoid deadlocks in case of panic or return
}

func main() {

	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println("Total views:", myPost.views)
}
