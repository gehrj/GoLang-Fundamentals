package main

import "fmt"

func main() {
	expression := (true && false) || (false && true) || !(false && false) // this should result in being true
	fmt.Println(expression)
}
