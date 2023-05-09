package main

import "fmt"

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func receive(ch <-chan int) {
	for v := range ch {
		fmt.Println("value is ", v)
	}

}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) //this executed after for condition false
	}()

	// close(c) //this will execute as soon as first c<-i is put into channel
	return c
}
