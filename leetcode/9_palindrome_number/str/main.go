package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 121

	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)

	for i, j := 0, len(str)-1; i <= j; {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}
