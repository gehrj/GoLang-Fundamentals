package main

import "fmt"

func main() {
	name := "Justin"
	fmt.Println(&name) // 0x8238984

	changeMe(&name)

	fmt.Println(&name) // 0x8238984
	fmt.Println(name)  // Dope
}

func changeMe(z *string) {
	fmt.Println(z)  // 0x8238984
	fmt.Println(*z) // Justin
	*z = "Dope"
	fmt.Println(z)  // 0x8238984
	fmt.Println(*z) // Dope
}
