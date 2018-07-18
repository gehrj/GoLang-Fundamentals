package main

import (
	"fmt"
	"sort"
)

func greatestNum(n ...int) int {
	sort.Ints(n)
	return n[len(n)-1]
}

func main() {
	fmt.Println(greatestNum(3, 4, 5, 2, 99, 6, 7, 90, 95))
}
