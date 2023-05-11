package main

import (
	"fmt"
	"strconv"
)

func main() {

	c := incrementor(2)

	var count int

	for n := range c {
		count++
		fmt.Println(n)
	}
	fmt.Println("Final Count", count)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprintf("Process:" + strconv.Itoa(i) + " printing: " + strconv.Itoa(k))
			}
			done <- true
		}(i)
	}
	go func(done chan bool, n int) {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}(done, n)
	return c

}
