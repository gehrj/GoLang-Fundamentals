package main

import "fmt"

func main() {
	var small int
	var large int

	fmt.Printf("Please enter a large number: ")
	fmt.Scan(&large)
	fmt.Printf("Now eneter a small number: ")
	fmt.Scan(&small)
	fmt.Printf("The remainder of %d divided by %d is %d\n", large, small, large%small)
}
