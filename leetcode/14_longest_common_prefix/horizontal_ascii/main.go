package main

import (
	"fmt"
)

func main() {
	strs := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, str := range strs[1:] {
		i := 0

		for i < len(prefix) && i < len(str) && prefix[i] == str[i] {
			i++
		}

		prefix = prefix[:i]

		if prefix == "" {
			return ""
		}
	}

	return prefix
}
