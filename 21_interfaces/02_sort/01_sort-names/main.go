package main

import (
	"fmt"
	"sort"
)

type people []string

// notice since we are creating methods off of slice which is already a reference type it is just (p people) not (p *people) like we would with struct
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

// we had to attach these methods to be able to use the sort.Sort(data Interface) method since it expects an interface that has those methods

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface
