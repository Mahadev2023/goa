package packageName

import (
	"fmt"
	"sort"
)

func main() {
	customSort()
}

func forIO() {
	// fmt.Fprintln(os.Stdout, "Hello From GO")
	// // os.Stdout.Write([]byte("hello"))
	// io.WriteString(os.Stdout, "GO TO WRITE")

	// // b := [2]byte{}
	// // os.Stdin.Read(b)
	// // fmt.Println(b)
}

func sortIntro() {

	xi := []int{5, 2, 1, 9, 4}

	sort.Ints(xi)
	fmt.Println(xi)

	s := []string{"mango", "apple", "orange", "pineApple", "A"}

	sort.Strings(s)
	fmt.Println(s)
}

// custom sort
type person struct {
	first string
	age   int
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func (p person) String() string {
	return fmt.Sprintf("%s:%d", p.first, p.age)
}

func customSort() {
	people := []person{
		{"Mahadev", 22},
		{"Abhijeet", 33},
		{"Ram", 12},
		{"Raja", 13},
	}
	fmt.Println(people)

	// sort.Sort(ByAge(people))

	fmt.Println(people)

	// sort.Slice(people, func(i, j int) bool { return people[i].age < people[j].age })
	// sort.Slice(people, func(i, j int) bool { return people[i].first < people[j].first })
	fmt.Println(people)
	// sort.Sort(people,func(a,b person) bool {a.age>b.age})
	// for _, per := range people {
	// 	fmt.Println(per.String())
	// }

}
