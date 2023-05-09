package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() string {
	return fmt.Sprint("Speaking person", *p)
}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println("Speaking from human", h.speak())
	h.speak()
}
func main2() {
	p := person{
		"Mahadev", "Godbole", 22,
	}

	//  saySomething(p);
	saySomething(&p)

	fmt.Println(p.speak())
}
