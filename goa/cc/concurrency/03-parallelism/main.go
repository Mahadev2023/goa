package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var wg sync.WaitGroup

	go foo(&wg)
	go bar(&wg)
	fmt.Println("GOroutine no ", runtime.NumGoroutine())
	wg.Add(runtime.NumGoroutine() - 1)
	wg.Wait()

}

func foo(referWG *sync.WaitGroup) {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(3 * time.Microsecond))
	}
	fmt.Println("DONE foo")
	referWG.Done()
}

func bar(referWG *sync.WaitGroup) {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(3 * time.Microsecond))

	}
	fmt.Println("DONE bar")
	referWG.Done()
}
