package main

import (
	"fmt"
)

func main() {
	input := []int{4, 1, 2, 1, 2}

	fmt.Println(singleNumber(input))
}

func singleNumber(nums []int) int {
	numsMap := make(map[int]bool, len(nums)/2+1)

	for _, num := range nums {
		if numsMap[num] {
			delete(numsMap, num)
		} else {
			numsMap[num] = true
		}
	}

	for k := range numsMap {
		return k
	}

	return 0
}
