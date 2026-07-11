package main

import (
	"fmt"
)

func main() {
	strs := []string{"облако", "область", "облик"}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := []rune(strs[0])

	for _, str := range strs[1:] {
		i := 0
		s := []rune(str)

		for i < len(prefix) && i < len(s) && prefix[i] == s[i] {
			i++
		}

		prefix = prefix[:i]

		if len(prefix) == 0 {
			return ""
		}
	}

	return string(prefix)
}
