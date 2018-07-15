package main

import "fmt"

// showing 0 through 199 in decimal, binary, and hexidecimal
func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
