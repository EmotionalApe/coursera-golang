package main

import (
	"fmt"
	"sync"
)

type fork struct {
	sync.Mutex
}

type philosopher struct {
	id                  int
	leftFork, rightFork *fork
}

func (p philosopher) eat(wg *sync.WaitGroup, permission chan bool) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		permission <- true // Request permission to eat
		p.leftFork.Lock()
		p.rightFork.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		fmt.Printf("finishing eating %d\n", p.id)
		fmt.Println()

		p.rightFork.Unlock()
		p.leftFork.Unlock()
		<-permission // Release permission after eating
	}
}

func main() {
	var wg sync.WaitGroup
	permission := make(chan bool, 2) // Channel to control the number of philosophers eating concurrently

	Forks := make([]*fork, 5)
	for i := 0; i < 5; i++ {
		Forks[i] = new(fork)
	}

	Philosophers := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		Philosophers[i] = &philosopher{i + 1, Forks[i], Forks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Philosophers[i].eat(&wg, permission)
	}

	wg.Wait()
}
