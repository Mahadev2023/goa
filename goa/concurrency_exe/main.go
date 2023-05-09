package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main3() {
	wg.Add(2)

	fmt.Println("Begin cpu", runtime.NumCPU())
	fmt.Println("Begin Goroutine", runtime.NumGoroutine())

	go func() {
		fmt.Println("Hello from thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from thing two")
		wg.Done()
	}()
	// go goR1()
	// go goR2()

	fmt.Println("Mid CPU", runtime.NumCPU())
	fmt.Println("Mid Goroutine", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Hello From main goroutine")

	fmt.Println("End CPU", runtime.NumCPU())
	fmt.Println("End Goroutine", runtime.NumGoroutine())

}

func goR1() {
	defer wg.Done()
	// time.Sleep(time.Second)
	fmt.Println("Message by Goroutine 1")

}
func goR2() {
	defer wg.Done()
	fmt.Println("Message by Goroutine 2")
}
