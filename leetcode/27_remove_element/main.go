package main

import "fmt"

func main() {
	input := []int{2, 9, 8, 2, 7, 6, 5, 2}

	k := removeElement(input, 2)

	fmt.Println(k, input[:k])
}

func removeElement(nums []int, val int) int {
	k := 0

	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}

	return k
}
