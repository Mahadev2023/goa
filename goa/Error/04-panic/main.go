package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// fmt.Println() print error//stdout
	//log.Println  //write of your choice like any file
	//log.Fatalln() //game over

	//log.Panicln()  deferred functions run //can use "recover"
	//panic()

	// _, err := os.Open("no-file.txt")
	// if err != nil {
	// 	// fmt.Println("err happened", err)
	// 	log.Println("err happened", err)
	// }

	// f, err := os.Create("log.txt")
	defer deferFunction()
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")

	if err != nil {
		// log.Println("err happened", err)
		// log.Fatalln(err) // program exit immediately
		// log.Panicln(err)
		panic(err)
	}

	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		log.Println(i)
	}

}

func deferFunction() {
	fmt.Println("Print from defer function")
}
