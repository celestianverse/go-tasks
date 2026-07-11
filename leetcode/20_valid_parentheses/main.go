package main

import "fmt"

func main() {
	input := "{()[]}"

	fmt.Println(isValid(input))
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	pairs := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, v := range s {
		if expected, ok := pairs[v]; ok {
			stack = append(stack, expected)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
