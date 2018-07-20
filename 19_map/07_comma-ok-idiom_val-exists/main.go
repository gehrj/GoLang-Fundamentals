package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"zero":  "Good",
		"one":   "HI",
		"two":   "Foo",
		"three": "bar",
	}
	fmt.Println(myGreeting)

	delete(myGreeting, "two") // this is where we are deleting "two" from map, comment it out to see how if statement changes

	// this inline if will set values of val and exists and then run if statement on exists. this is good for scoping
	if val, exists := myGreeting["two"]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
	fmt.Println(myGreeting)
}
