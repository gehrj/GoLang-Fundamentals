package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Miss",
			Last:  "Moneypenny",
			Age:   18,
		},
		LicenseToKill: false,
	}
	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.Person.First, p2.Last, p2.Person.Age, p2.LicenseToKill)
}
