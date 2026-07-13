package main

import "fmt"

func main() {
	input := 2147483647

	fmt.Println(mySqrt(input))
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left := 1
	right := x / 2

	for left <= right {
		mid := left + (right-left)/2

		if mid <= x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
