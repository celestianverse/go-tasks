package main

import (
	"fmt"
	"unicode"
)

func main() {
	input := "A man, a plan, a canal: Panama"

	fmt.Println(isPalindrome(input))
}

func isPalindrome(s string) bool {
	runes := []rune(s)

	i, j := 0, len(runes)-1

	for i < j {
		for i < j && !isAlphanumeric(runes[i]) {
			i++
		}
		for i < j && !isAlphanumeric(runes[j]) {
			j--
		}

		if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
