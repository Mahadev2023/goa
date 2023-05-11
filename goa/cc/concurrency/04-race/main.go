package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

// var mu sync.Mutex
var mu sync.RWMutex
var counter int

func main() {
	wg.Add(2)
	go incrementor("FOO:")
	go incrementor("BARR:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		// mu.Lock()
		mu.Lock()
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3) * int(time.Microsecond)))
		counter = x
		fmt.Println(s, i, "Counter", counter)

		mu.Unlock()

		// mu.RLock()
		// fmt.Println(s, i, "Counter", counter)
		// mu.RUnlock()
	}
	wg.Done()
}
