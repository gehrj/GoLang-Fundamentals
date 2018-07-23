package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{5, 3, 6, 7, 19, 233, 45, 89}
	sort.Ints(n)
	fmt.Println(n)
}
