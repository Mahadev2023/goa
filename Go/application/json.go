package packageName

import (
	"encoding/json"
	"fmt"
)

//as first is not public so make it public . if marshal output is [{},{}]
// type person struct {
// 	first string
// 	last  string
// 	age   int
// }

type person struct {
	First string
	Last  string
	Age   int
}

// pass at time of unmarshal
type uPerson struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	// m()
	unm()

}

// marshal the data
func m() {
	p1 := person{
		First: "Mahadev",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Abhi",
		Last:  "Panchal",
		Age:   43,
	}

	people := []person{
		p1,
		p2,
	}
	fmt.Println("Data:", people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}

// unmarshal data
func unm() {
	s := `[{"First":"Mahadev","Last":"Bond","Age":32},{"First":"Abhi","Last":"Panchal","Age":43}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// people:=[]uPerson{}

	var people []uPerson

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data\n", people)

	for i, v := range people {
		fmt.Println("Person Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
