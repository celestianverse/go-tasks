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

	runes := make([][]rune, len(strs))

	for i, str := range strs {
		runes[i] = []rune(str)
	}

	for i := 0; i < len(runes[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(runes[j]) || runes[j][i] != runes[0][i] {
				return string(runes[0][:i])
			}
		}
	}

	return string(runes[0])
}
