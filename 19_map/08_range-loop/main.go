package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good",
		1: "morning",
		2: "justin",
		3: "gehr",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
