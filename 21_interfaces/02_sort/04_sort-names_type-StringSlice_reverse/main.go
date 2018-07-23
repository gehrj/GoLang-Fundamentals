package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	// sort.Reverse takes an interface and returns an interface that we then use in sort to get reverse order

	// everything below is for experimentation and to help visualize how the above is happening
	// uncomment and comment out the above to see these tests

	// sort.Sort(sort.StringSlice(s))
	// fmt.Println(s)
	// fmt.Printf("just s: %T\n", s)
	// fmt.Printf("s converted to StringSlice: %T\n", sort.StringSlice(s))
	// fmt.Printf("s reversed: %T\n", sort.Reverse(sort.StringSlice(s)))
	// i := sort.Reverse(sort.StringSlice(s))
	// fmt.Println(i)
	// fmt.Printf("%T\n", i)
	// sort.Sort(i)
	// fmt.Println(s)
}
