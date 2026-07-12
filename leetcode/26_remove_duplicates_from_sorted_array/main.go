package main

import "fmt"

func main() {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	k := removeDuplicates(input)

	fmt.Println(k, input[:k])
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
