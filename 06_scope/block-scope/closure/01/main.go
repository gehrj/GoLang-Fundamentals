package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The cridit belongs with the one who is in the ring."
		fmt.Println(y)
	}
	// fmt.Println(y) // outside scope of y so will not work
}
