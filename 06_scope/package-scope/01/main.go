package main

import "fmt"

// x is accessible in whole package since it is declared outside of {}
var x int = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
