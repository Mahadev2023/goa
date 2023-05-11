package main

import "fmt"

func main() {
	//ex1
	// c := make(chan int)
	// // c <- 1
	// go func() {
	// 	c <- 1
	// }()
	// fmt.Println(<-c)
	// ###########################3
	// ex2
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

}
