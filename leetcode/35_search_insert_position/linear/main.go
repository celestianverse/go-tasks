package main

import "fmt"

func main() {
	nums := []int{1, 3, 5}
	target := 4

	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if target == v {
			return i
		}
		if target < v {
			return i
		}
		if i == len(nums)-1 {
			return len(nums)
		}
	}

	return 0
}
