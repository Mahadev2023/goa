package main

import "fmt"

func main() {
	q := make(chan int)

	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(q)

	}()
	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("value form c ", v)
		case v := <-q:
			fmt.Println("q value get", v)
			return
		}
	}
}
