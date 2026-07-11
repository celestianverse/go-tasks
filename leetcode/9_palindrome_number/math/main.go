package main

import "fmt"

func main() {
	x := 121

	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}
	if x%10 == 0 {
		return false
	}

	reversed := 0

	for x > reversed {
		reversed = x%10 + reversed*10
		x /= 10
	}

	return x == reversed || x == reversed/10
}
