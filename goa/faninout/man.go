package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	//FAN OUT
	//Distribute the sq work across two goroutines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	//FAN IN
	//consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n) //4 then 9 ,0r 9 then 4
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	//output go routine for each input channel in cs
	//until closing c input place c input ont out chan completion call done
	output := func(c <-chan int) {
		// wg.Add(1)
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	//start goroutine to close out once all the output goroutines
	//done .this must start after the wg.Add call
	go func() {
		wg.Wait()
		close(out)

	}()
	return out

}

func mergemy(c1, c2 chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		// wg.Add(1)
		for n := range c1 {
			out <- n
		}
		wg.Done()
	}()
	go func() {
		// wg.Add(1)
		for n := range c2 {
			out <- n
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(out)

	}()
	return out

}
