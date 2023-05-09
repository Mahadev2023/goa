// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type user struct {
// 	First   string
// 	Last    string
// 	Age     int
// 	Sayings []string
// }

// func main() {
// 	u1 := user{
// 		First: "James",
// 		Last:  "Bond",
// 		Age:   32,
// 		Sayings: []string{
// 			"Shaken, not stirred",
// 			"Youth is no guarantee of innovation",
// 			"In his majesty's royal service",
// 		},
// 	}

// 	u2 := user{
// 		First: "Miss",
// 		Last:  "Moneypenny",
// 		Age:   27,
// 		Sayings: []string{
// 			"James, it is soo good to see you",
// 			"Would you like me to take care of that for you, James?",
// 			"I would really prefer to be a secret agent myself.",
// 		},
// 	}

// 	u3 := user{
// 		First: "M",
// 		Last:  "Hmmmm",
// 		Age:   54,
// 		Sayings: []string{
// 			"Oh, James. You didn't.",
// 			"Dear God, what has James done now?",
// 			"Can someone please tell me where James Bond is?",
// 		},
// 	}

// 	users := []user{u1, u2, u3}

// 	fmt.Println(users)
// 	fmt.Println("##############################")

// 	json.NewEncoder(os.Stdout).Encode(users)

// 	// json.NewEncoder().Encode(users)
// 	// os.WriteFile("/temp.txt",)

// 	// your code goes here

// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
// 	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

// 	fmt.Println(xi)
// 	// sort xi
// 	sort.Slice(xi, func(i, j int) bool { return xi[i] < xi[j] })
// 	fmt.Println(xi)

// 	fmt.Println(xs)
// 	// sort xs
// 	sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
// 	fmt.Println(xs)
// }

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// fmt.Println(users)

	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		for _, v := range user.Sayings {
			fmt.Println("\t", v)
		}
	}

	sort.Sort(ByAge(users))
	fmt.Println("#################")

	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		sort.Strings(user.Sayings)
		for _, v := range user.Sayings {
			fmt.Println("\t", v)
		}
	}

	sort.Sort(ByLast(users))
	fmt.Println("#################")

	for _, user := range users {
		fmt.Println(user.First, user.Last, user.Age)
		sort.Strings(user.Sayings)
		for _, v := range user.Sayings {
			fmt.Println("\t", v)
		}
	}

	// your code goes here
	// sort.Slice(users, func(i, j int) bool { return users[i].Age < users[j].Age })
	// fmt.Println(users)

}
