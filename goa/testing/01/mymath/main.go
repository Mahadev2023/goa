// Package mymath provides ACME inc math solutions
package mymath

//Sum adds an unlimited number of values of type int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// kanaka@DESKTOP-N833H7D MINGW64 ~/OneDrive - Kanaka Software Consulting Pvt Ltd/Desktop/goo/goa/doc/01/mymath (main)
// $ go doc
//  go doc fmt.Printf

// https://stackoverflow.com/questions/63442354/godoc-command-not-found
//go install golang.org/x/tools/cmd/godoc
//godoc -http :8080
