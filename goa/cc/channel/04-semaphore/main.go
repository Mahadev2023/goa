package main

import "fmt"

//using semaphore
func main() {

	c := make(chan int)

	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done   //wait for first single
		<-done   //wait for second value
		close(c) //close the c
	}()
	/*//this is wrong
	<-done//wait for first single
		<-done//wait for second value
		close(c)//close the c
	*/
	for n := range c {
		fmt.Println(n)
	}

}

//example 2

// package main

// import "fmt"

// func main() {
// 	n := 10

// 	c := make(chan int)
// 	done := make(chan bool)

// 	for i := 0; i < n; i++ {
// 		go func() {
// 			for i := 0; i < 10; i++ {
// 				c <- i
// 			}
// 			done <- true
// 		}()
// 	}

// 	go func() {
// 		for i := 0; i < n; i++ {
// 			<-done
// 		}
// 		close(c)
// 	}()

// 	for n := range c {
// 		fmt.Println(n)
// 	}
// }
