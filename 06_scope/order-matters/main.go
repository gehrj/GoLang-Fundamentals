package main

import "fmt"

func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42
}
// notice order does not matter for y since its scope is enclosing the inner scope of main
var y = 42