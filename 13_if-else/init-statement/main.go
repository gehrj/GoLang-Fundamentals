package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}
}

// food will be "Chocolate" if b is true
