package main

import "fmt"

func main() {
	var sum int
	for i := 3; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("The sum is:", sum)
}
