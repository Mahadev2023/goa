package main

import (
	"fmt"
)

type customErr struct {
	name  string
	error string
}

func (ce customErr) Error() string {
	return "this is my new error " + ce.name + ce.error
}

func main() {
	c1 := customErr{"Something bad", " Error is occured in program"}
	foo(c1)

}

func foo(e error) {
	// fmt.Println("foo ran - ", e.info())not work

	fmt.Println("foo ran - ", e.(customErr).name) //assertion

	//conversion
	type hotdog int
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Printf("%T\n", y)

}
