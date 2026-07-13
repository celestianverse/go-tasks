package main

import "fmt"

func main() {
	input := 6

	fmt.Println(climbStairs(input))
}

func climbStairs(n int) int {
	if n < 4 {
		return n
	}

	prev := 2
	current := 3

	for i := 4; i <= n; i++ {
		prev, current = current, prev+current
	}

	return current
}
