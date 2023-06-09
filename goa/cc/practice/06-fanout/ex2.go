package main

import (
	"fmt"
	"time"
)

var workerId int = 0
var publisherId int = 0

// Publishers push data into a channel
func Publisher(out chan string) {
	publisherId += 1
	thisId := publisherId
	dataId := 0
	for {
		dataId += 1
		fmt.Printf("Publisher %d is pushing data\n", thisId)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisId, dataId)
		out <- data
	}
}

func WorkerProcess(in <-chan string) {
	workerId += 1
	thisId := workerId
	for {
		fmt.Printf("%d: waiting for input...\n", thisId)
		input := <-in
		fmt.Printf("%d: input is: %s\n", thisId, input)
	}
}

func main() {
	input := make(chan string)
	go WorkerProcess(input)
	go WorkerProcess(input)
	go WorkerProcess(input)
	go Publisher(input)
	go Publisher(input)
	go Publisher(input)
	go Publisher(input)
	time.Sleep(1 * time.Millisecond)
}
