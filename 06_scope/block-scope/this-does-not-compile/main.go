package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// does not have access to x
	fmt.Println(x)
}
