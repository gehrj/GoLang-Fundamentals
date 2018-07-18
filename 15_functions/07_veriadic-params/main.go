package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// the ... can be used on a type if it is last argument in func. it means any number of 0 to infinity arguments can be used of that type there
// it puts them in a slice so sf is a slice of float64
