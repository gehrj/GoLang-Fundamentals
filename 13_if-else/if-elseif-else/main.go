package main

import "fmt"

func main() {
	if false {
		fmt.Println("This wont run")
	} else if true {
		fmt.Println("This will run")
	} else {
		fmt.Println("We wont even make it here")
	}
}
