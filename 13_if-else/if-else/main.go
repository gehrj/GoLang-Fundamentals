package main

import "fmt"

func main() {
	if false {
		fmt.Println("this will not run first print statement")
	} else {
		fmt.Println("this will run second print statement")
	}
}
