package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//example0()
	//example1()
	//example2()
	//example3()
	//example4()
	example5()
}

func example0() {
	count("foo")
	count("bar")
}

func example1() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("foo")
		wg.Done()
	}()

	wg.Wait()
}

func count(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, text)
		time.Sleep(time.Millisecond * 500)
	}
}

func example2() {
	c := make(chan string)
	go countChannel("foo", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func countChannel(text string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- text
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}

func example3() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2s"
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}

func example4() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2s"
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func example5() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
