package main

import "fmt"

// same as loop from numeral systems which showed us decimal, binary, and hexadecimal
// but added %q which shows us UTF-8 also

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
