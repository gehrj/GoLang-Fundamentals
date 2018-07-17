package main

import "fmt"

func main() {
	switch "Stephanie" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Justin":
		fmt.Println("Wassup Justin")
	case "Stephanie":
		fmt.Println("Wassup Stephanie")
		fallthrough
	case "Zach":
		fmt.Println("Wassup Zach")
		fallthrough
	case "Doreen":
		fmt.Println("Wassup Doreen")
	case "Brandon":
		fmt.Println("Wassup Brandon")
	}
}

// this will print out stephanie, zach, and doreen and then stop
// it prints out the extra two cases because of the fallthroughs
