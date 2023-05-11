package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	in := gen()

	c := factorial(in)
	for n := range c {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 55; j < 65; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for n := range in {
			// out <- fact(n)
			wg.Add(1)
			go func(n1 int, wg *sync.WaitGroup) {
				out <- fact(n1)
				wg.Done()

			}(n, &wg)
		}
		fmt.Println("NUM OF GOr", runtime.NumGoroutine())
		wg.Wait()
		close(out)

	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
