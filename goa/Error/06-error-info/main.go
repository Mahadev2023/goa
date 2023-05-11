package main

import (
	"fmt"
	"log"
	"math"
)

type MyErr struct {
	lat  string
	long string
	err  error
}

func (n MyErr) Error() string {
	return fmt.Sprintf("a norgate math error occured : %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		name := fmt.Errorf("norgate math redux: square root of negative number")

		return 0, MyErr{"645.233 N", "3344.3 w", name}
	}
	return math.Sqrt(f), nil
}
