package main

import "fmt"

func main() {
	name := "Justin"
	fmt.Println(name) // Justin

	changeMe(name)

	fmt.Println(name) // Justin
}

func changeMe(z string) {
	fmt.Println(z) // Justin
	z = "Is Amazzzzzing"
	fmt.Println(z) // Is Amazzzzzzing
}
