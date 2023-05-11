// // package main

// // import (
// // 	"fmt"
// // 	"sync"
// // )

// // func push(c chan int, wg *sync.WaitGroup) {
// // 	for i := 0; i < 5; i++ {
// // 		c <- i
// // 	}
// // 	wg.Done()
// // }

// // func pull(c chan int, wg *sync.WaitGroup) {
// // 	for i := 0; i < 5; i++ {
// // 		result, ok := <-c
// // 		fmt.Println(result, ok)
// // 	}
// // 	wg.Done()
// // }

// // func main() {
// // 	var wg sync.WaitGroup
// // 	wg.Add(2)
// // 	c := make(chan int)

// // 	go push(c, &wg)
// // 	go pull(c, &wg)

// // 	wg.Wait()
// // }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	requests := make(chan int, 5)
// 	for i := 1; i <= 5; i++ {
// 		requests <- i
// 	}
// 	close(requests)
// 	limiter := time.Tick(200 * time.Millisecond)
// 	for req := range requests {
// 		<-limiter
// 		fmt.Println("request", req, time.Now())
// 	}
// 	burstyLimiter := make(chan time.Time, 3)

// 	for i := 0; i < 3; i++ {
// 		burstyLimiter <- time.Now()
// 	}

// 	go func() {
// 		for t := range time.Tick(200 * time.Millisecond) {
// 			burstyLimiter <- t
// 		}
// 	}()
// 	burstyRequests := make(chan int, 5)
// 	for i := 1; i <= 5; i++ {
// 		burstyRequests <- i
// 	}
// 	close(burstyRequests)
// 	for req := range burstyRequests {
// 		<-burstyLimiter
// 		fmt.Println("request", req, time.Now())
// 	}
// }

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	numOfGoRoutines := 10
	wg.Add(numOfGoRoutines)
	ch := make(chan int, numOfGoRoutines)

	go func(wg *sync.WaitGroup) {
		for i := 0; i < numOfGoRoutines; i++ {
			a := i
			go sqr(ch, a, wg)
		}
		wg.Wait()
		close(ch)
	}(&wg)
	fmt.Println("After WAIT")

	var res int
	for i := range ch {
		res += i
	}
	ch = nil
	fmt.Println("result = ", res)

}

func sqr(ch chan int, val int, wg *sync.WaitGroup) {
	fmt.Println("go - ", val)
	s := val * val
	ch <- s
	wg.Done()
}
