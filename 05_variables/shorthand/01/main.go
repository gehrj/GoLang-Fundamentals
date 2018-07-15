package main

import "fmt"

// when using := for assignment it will take on the type of what you are assinging variable too. this has to be used inside of a func
func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
