package main

import "fmt"

func main() {
	s := "MCMXCIV"

	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	values := [256]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0

	for i := range s {
		j := i + 1
		current := values[s[i]]

		if j < len(s) && current < values[s[j]] {
			result -= current
		} else {
			result += current
		}
	}

	return result
}
