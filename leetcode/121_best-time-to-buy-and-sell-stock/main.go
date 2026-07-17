package main

import "fmt"

func main() {
	input := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(input))
}

func maxProfit(prices []int) int {
	profit := 0
	low := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < low {
			low = prices[i]
		}
		if prices[i]-low > profit {
			profit = prices[i] - low
		}
	}

	return profit
}
