package main

import "fmt"

func main() {
	n := HashBucket("Go", 12)
	fmt.Println(n)
}

func HashBucket(word string, buckets int) int {
	// letter := rune(word[0])  change everything to int32 and comment out other letter assignment to see this version work
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
