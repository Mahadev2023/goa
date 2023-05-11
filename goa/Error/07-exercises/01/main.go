// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type person struct {
// 	First   string
// 	Last    string
// 	Sayings []string
// }

// func main() {
// 	p1 := person{
// 		First:   "James",
// 		Last:    "Bond",
// 		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
// 	}

// 	bs, err := json.Marshal(p1)
// 	if err != nil {
// 		fmt.Println(err) //
// 		log.Fatalln("JSON did not marshal", err)
// 	}
// 	fmt.Println(string(bs))

// }
// //
// ####################################
//ex2

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		// fmt.Println(err)
		// return
		// or
		log.Fatalln(err)

	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		// msg := fmt.Errorf("Someting happen during marshaling %v", err)

		// return bs,msg
		// return []byte{}, fmt.Println("there was an error in to JSON",err) not work as println return two values

		//return []byte{}, errors.New(fmt.Sprint("there was an error in toJSON", err))
		return []byte{}, errors.New(fmt.Sprintf("there was an error in toJSON: %v", err))

	}
	return bs, err
}
