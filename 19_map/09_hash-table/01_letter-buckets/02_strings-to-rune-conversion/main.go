package main

import "fmt"

func main() {
	letter := rune("A"[0]) // string is slice of bytes, rune are bytes so using rune on string will give you the slice of bytes
	fmt.Println(letter)
}
