// This sample program demonstrates how the goroutine scheduler
// will time slice goroutines on a single thread.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// main is the entry point for all Go programs.
// Notes:
// Create Goroutines
// Waiting To Finish
// B:2
// B:3
// ...
// B:4583
// B:4591
// A:3		** Goroutines Swapped
// A:5
// ...
// A:4561
// A:4567
// B:4603	** Goroutines Swapped
// B:4621
// ...
// Completed B
// A:4457	** Goroutines Swapped
// A:4463
// ...
// A:4993
// A:4999
// Completed A
// Terminating Program
func main() {
	// Allocate 1 logical processors for the scheduler to use.
	runtime.GOMAXPROCS(1)

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Terminating Program")
}

// printPrime displays prime numbers for the first 5000 numbers.
func printPrime(prefix string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
