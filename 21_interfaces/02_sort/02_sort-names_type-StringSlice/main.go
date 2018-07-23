package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.StringSlice(s).Sort() // we are converting s to type StringSlice which has method Sort() on it which we then use
	// sort.Sort(sort.StringSlice(s)) // we are converting s to type StringSlice and passing it as argument in Sort method which expects Sort(data Interface) which is satisfied by converting to type StringSlice
	// both of these accomplish same thing just different ways to do it.
	fmt.Println(s)
}
