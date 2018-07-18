package main

import "fmt"

func main() {
	func() {
		fmt.Println("I am executing myself lol")
	}()
}
