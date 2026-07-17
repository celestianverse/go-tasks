package main

import "fmt"

func main() {
	input := 5

	fmt.Println(getRow(input))
}

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}

	return row
}
