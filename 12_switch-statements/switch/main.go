package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("You have no friends")
	}
}

// no default fallthrough
// fallthrough is optional
// you can specify fallthrough by explicity stating it
// break isn't needed like in other languages
