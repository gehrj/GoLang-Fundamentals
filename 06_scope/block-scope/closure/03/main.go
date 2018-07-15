package main

import "fmt"

func main() {
	x := 0
	// also an example of an anonynomous func and func expression, needs to be this way to put a func within a func
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
