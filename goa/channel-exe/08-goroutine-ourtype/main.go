// write a program that
// ○ launches 10 goroutines
// ■ each goroutine adds 10 numbers to a channel
// ○ pull the numbers off the channel and print them

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// c:=make(chan int)
	c := make(chan int)

	go gen(c)
	var ff []int
	var count = 1
	for v := range c {
		// fmt.Println("value ", v, count)
		ff = append(ff, v)
		count++
	}
	fmt.Println(ff)

}

func gen(c chan int) {
	var wg sync.WaitGroup
	// var mu sync.Mutex
	const goroutine = 10
	wg.Add(goroutine)

	var counter int64

	// var counter = 1
	// go func() {
	for i := 1; i <= goroutine; i++ {
		// fmt.Println("hello")
		go func(j int) {
			for v := 1; v <= 10; v++ {

				atomic.AddInt64(&counter, 1)
				c <- int(atomic.LoadInt64(&counter))

				//handle sync using mutex
				// mu.Lock()

				// c <- counter
				// counter++
				// mu.Unlock()

				fmt.Println("helloo", v*j)
			}
			wg.Done()

		}(i)

	}
	// }()
	// fmt.Println("hello before wait ")
	wg.Wait()
	// fmt.Println("hello after")

	close(c)

}
