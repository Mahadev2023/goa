package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {

	// f := factorial(7000000000)
	num := big.NewInt(70)
	f := factorial(num)

	for n := range f {
		fmt.Println(time.Now())
		fmt.Println("Total:", n)
	}

}

func factorial(n *big.Int) chan *big.Int {
	fmt.Println(time.Now())
	out := make(chan *big.Int)
	go func() {
		one := big.NewInt(1)
		total := big.NewInt(1)
		// fmt.Println("n", n)
		// fmt.Println(one.Cmp(n) < 0)
		for i := big.NewInt(1); i.Cmp(n) <= 0; i.Add(i, one) {
			total.Mul(total, i)
			// fmt.Println(total)
		}
		fmt.Println("out", total)
		out <- total
		fmt.Println("chan going to close", time.Now())

		close(out)
	}()
	return out
}
