package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

var wg sync.WaitGroup

// var mu sync.Mutex
// var mu sync.RWMutex
// var counter int

func main() {
	wg.Add(2)
	go incrementor("FOO:")
	go incrementor("BARR:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {

		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "counter", atomic.LoadInt64(&counter))
		// fmt.Println(s, i, "Counter", counter)reading race condition

		//mutex code
		// mu.Lock()
		// x := counter
		// x++
		// time.Sleep(time.Duration(rand.Intn(3) * int(time.Microsecond)))
		// counter = x
		// fmt.Println(s, i, "Counter", counter)

		// mu.Unlock()

	}
	wg.Done()
}
