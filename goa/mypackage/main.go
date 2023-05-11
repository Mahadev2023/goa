package main

import (
	"fmt"

	"github.com/Mahadev2023/mypackage/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.Mg())
}
