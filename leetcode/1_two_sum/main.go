package main

import "fmt"

func main() {
	nums := []int{2, 9, 7, 15, 11}
	target := 18

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))

	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}

	return nil
}
