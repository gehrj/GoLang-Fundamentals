package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10) this will start at 1 in binary and then shift over 10 places.
	MB = 1 << (iota * 10) // 1 << (2 * 10) this will start at 1 in binary and then shift over 20 places.
	GB = 1 << (iota * 10) // 1 << (3 * 10) this will start at 1 in binary and then shift over 30 places.
)

func main() {
	fmt.Println("binary \t\t decimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

}
