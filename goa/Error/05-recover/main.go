package main

import "fmt"

func main() {
	fmt.Println("Main start the learning")
	iCausePanics()
	fmt.Println("Main:Post Panic")
}

func iCausePanics() {
	defer tryToRecover()
	fmt.Println("about to panic")
	// panic("Func : cannot continue to execute")
	fmt.Println("I just Panicked")
}

func tryToRecover() {
	if err := recover(); err != nil {
		fmt.Println("Recovered from Error: ", err)
	}
}

// https://www.youtube.com/watch?v=kTNlQiRoI2w&list=PLa6iDxjj_9qXaVIemjejvomRgKu3r7t3J&index=30
