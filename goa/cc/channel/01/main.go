package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			// time.Sleep(time.Second - (time.Nanosecond * 100))
			c <- i
		}

	}()

	go func() {
		for {
			fmt.Println(<-c)

		}
	}()

	time.Sleep(time.Second)

}
