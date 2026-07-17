package main

import "fmt"

func main() {
	input := 7

	fmt.Println(generate(input))
}

func generate(numRows int) [][]int {
	rows := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1

		for j := 1; j < i; j++ {
			row[j] = rows[i-1][j-1] + rows[i-1][j]
		}

		rows[i] = row
	}

	return rows
}
