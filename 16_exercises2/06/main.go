package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(threeDigitProducts())
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	var i int
	j := len(s) - 1

	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func threeDigitProducts() int {
	var largest int
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if isPalindrome(i * j) {
				if largest < i*j {
					largest = i * j
				}
			}
		}
	}
	return largest
}

// the goal of this problem is to find the largest palindrome you can make by multiplying two 3 digit numbers together
