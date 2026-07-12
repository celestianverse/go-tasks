package main

import "fmt"

func main() {
	haystack := "mississippi"
	needle := "issip"

	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i] == needle[0] {
			match := true

			for j := 0; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					match = false
					break
				}
			}

			if match {
				return i
			}
		}
	}

	return -1
}
