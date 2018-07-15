package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	// the # adds in 0x
	fmt.Printf("%d - %b - %#x\n", 42, 42, 42)
	// the captial makes the output capital
	fmt.Printf("%d - %b - %#X\n", 42, 42, 42)
	// \t is a tab	
	fmt.Printf("%d \t  %b \t %#X\n", 42, 42, 42)
}
