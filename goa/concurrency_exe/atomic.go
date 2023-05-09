package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	var gr = 100
	wg.Add(gr)

	var incrementor int64

	fmt.Println("Number Goroutine Before", runtime.NumGoroutine())

	for i := 0; i < gr; i++ {
		go func() {

			atomic.AddInt64(&incrementor, 1)
			fmt.Println(atomic.LoadInt64(&incrementor))
			wg.Done()
		}()
	}
	fmt.Println("Number Goroutine mid", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("Number Goroutine After", runtime.NumGoroutine())
	fmt.Println("counter", incrementor)

}
