// package main

// import (
// 	"errors"
// 	"log"
// 	"math"
// )

// func main() {
// 	_, err := sqrt(-10)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		return 0, errors.New("norgate math:square root of negative number")
// 	}
// 	return math.Sqrt(f), nil
// }

// 33333333333333333333333333333

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var MyErr = errors.New("Sqrt of negative number")

func main() {
	fmt.Printf("Type of MyErr :%T\n", MyErr) //:*errors.errorString
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, MyErr
		return 0, fmt.Errorf("square root of -number : %v", f)
	}
	return math.Sqrt(f), nil
}
