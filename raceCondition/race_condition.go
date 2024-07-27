package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		counter++
	}
}

func decrement(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		counter--
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go increment(&wg)
	go decrement(&wg)

	wg.Wait()
	fmt.Println("Final counter : ", counter)
}

/*
Concurrent Access:
The increment function increments the counter variable 50000 times.
The decrement function decrements the counter variable 50000 times.
Both functions run concurrently, accessing and modifying the shared counter variable.

Interleaving of Operations:
Since there is no synchronization, the operations on counter by the two goroutines can interleave in many ways.
Each read-modify-write operation on counter is not atomic. This means that one goroutine can read the counter value while another goroutine is in the process of updating it.

Non-Deterministic Results:
The final value of counter depends on the exact order of operations, which is determined by the Go scheduler and can vary with each execution.
As the number of operations increases, the likelihood of more interleaved and conflicting operations increases, leading to more variability in the output.
*/
