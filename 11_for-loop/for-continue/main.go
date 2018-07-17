package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		// continue means to move on to the next iteration of the loop so this program will only print odd numbers 1-51
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
