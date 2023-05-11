// package main

// import "fmt"

// func main() {
// 	//setting pipeline
// 	c := gen(2, 3)
// 	out := sq(c)

// 	//output
// 	fmt.Println(<-out)
// 	fmt.Println(<-out)
// }

// func gen(nums ...int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, v := range nums {
// 			out <- v
// 		}
// 	}()
// 	return out
// }

// func sq(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out

// }

package main

import "fmt"

func main() {
	//setting pipeline
	c := gen(2, 3)
	out := sq(c)

	//output
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out

}
