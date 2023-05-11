package main

import "fmt"

func main() {

	c1 := incrementor("FOO:")
	c2 := incrementor("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	a := <-c3
	b := <-c4
	fmt.Println("Final Counter", a, b, a+b)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
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
