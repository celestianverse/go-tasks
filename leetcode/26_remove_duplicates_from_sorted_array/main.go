package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	k := removeDuplicates(input)

	fmt.Println(k, input[:k])
}

func removeDuplicates(nums []int) int {
	k := 0

	for i, v := range nums {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			continue
		} else {
			nums[k] = v
			k++
		}
	}

	return k
}
