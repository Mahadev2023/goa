package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main1() {

	var wg sync.WaitGroup

	var mu sync.Mutex
	// var mw sync.RWMutex

	var gr = 100

	wg.Add(gr)

	var counter = 0

	fmt.Println("Number Goroutine Before", runtime.NumGoroutine())

	for i := 0; i < gr; i++ {
		go func() {
			mu.Lock()
			// mw.RLock()
			var v = counter

			// runtime.Gosched()//Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
			v++
			counter = v
			// mw.RUnlock()
			fmt.Println("counter ", counter)
			mu.Unlock()

			wg.Done()
		}()
	}
	fmt.Println("Number Goroutine mid", runtime.NumGoroutine())
	fmt.Println("counter mid", counter)
	wg.Wait()
	fmt.Println("Number Goroutine After", runtime.NumGoroutine())
	fmt.Println("counter", counter)

}
