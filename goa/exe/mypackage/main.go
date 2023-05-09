package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
}

type customer struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	fmt.Println("Hello, Modules!")
	users := []user{
		{"mahadev", 12},
		{"Abhi", 22},
	}
	jdata, _ := json.Marshal(users)

	fmt.Println(string(jdata))

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var customers []customer

	error := json.Unmarshal([]byte(s), &customers)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(customers)

	for i, customer := range customers {
		fmt.Println("customer $", i)
		fmt.Println(customer)
	}

}
