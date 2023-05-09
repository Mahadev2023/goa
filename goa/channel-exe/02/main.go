package main

import "fmt"

func main() {
	// c := make(chan<- int)
	// c := make(<-chan int)
	c := make(chan int)

	go func() {
		c <- 42

	}()

	// ch:=(<-chan int)c
	fmt.Println(<-c)

	fmt.Println("###########")
	fmt.Printf("c \t %T\n", c)
}
