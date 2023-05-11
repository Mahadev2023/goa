package main

import (
	"fmt"
)

func main() {

	// n, err := fmt.Println("Hello")//also here ln => new line character
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(n)//result 6

	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sports: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)

}
