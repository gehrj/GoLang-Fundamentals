package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

// when you attach a String() string method to a type like this it will be used when you try and print something of that type
func (p person) String() string {
	return fmt.Sprintf("YAYAYA %s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []person based on the Age field.

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// func (a ByAge) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people[0])
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}

// String() string
// https://golang.org/doc/effective_go.html#printing
