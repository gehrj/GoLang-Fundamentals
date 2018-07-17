package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))
}

// 'i' is a rune. its codepoint would be 105
// type of rune is int32 since rune is alias for int32
