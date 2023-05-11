// package main

// import "fmt"

// func main() {
// 	c := incrementor()
// 	cSum := puller(c)
// 	for n := range cSum {
// 		fmt.Println(n)
// 	}
// }

// func incrementor() chan int {
// 	out := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			out <- i
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func puller(c chan int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		var sum int
// 		for n := range c {
// 			sum += n
// 		}
// 		out <- sum
// 		close(out)
// 	}()
// 	return out
// }

//channel direction

package main

import "fmt"

func main() {
	c := incrementor()
	// cSum := puller(c)
	for n := range puller(c) {
		fmt.Println(n)
	}
}

//return receiving channel
func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

//get receive channel and return also receive channel
func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
