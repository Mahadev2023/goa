package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	//fan out multiple function reading from  same channel
	// c0 := factorial(in)
	// c1 := factorial(in)
	// c2 := factorial(in)
	// c3 := factorial(in)
	// c4 := factorial(in)
	// c5 := factorial(in)
	// c6 := factorial(in)
	// c7 := factorial(in)
	// c8 := factorial(in)
	// c9 := factorial(in)

	xc := fanOut(in, 10)

	//fan in
	var y int
	for n := range merge(xc...) {
		y++
		fmt.Println(y, "\t", n)
	}
	fmt.Println(y)
}
func fanOut(in <-chan int, n int) []<-chan int {
	var xc []<-chan int
	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in))
	}
	return xc
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 30; j < 40; j++ {
				out <- j

			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
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
