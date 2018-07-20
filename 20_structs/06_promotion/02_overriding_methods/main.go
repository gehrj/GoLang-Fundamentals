package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
	p1 := Person{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := DoubleZero{
		Person: Person{
			Name: "James Bond",
			Age:  20,
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()

	// just like with overriding values the outer method will overwrite the inner method when accessing outer object method directly ex. p2.Greeting (outer) p2.Person.Greeting() (inner)
}
