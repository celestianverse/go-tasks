package main

import "fmt"

func main() {
	a := "1010"
	b := "1011"

	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	maxLen := max(len(a), len(b))
	result := make([]byte, maxLen+1)
	carry := 0

	for i := 0; i < maxLen; i++ {
		ai := len(a) - 1 - i
		bi := len(b) - 1 - i
		sum := carry

		if ai >= 0 {
			sum += int(a[ai] - '0')
		}
		if bi >= 0 {
			sum += int(b[bi] - '0')
		}

		result[maxLen-i] = byte(sum%2 + '0')
		carry = sum / 2
	}

	if carry == 1 {
		result[0] = '1'

		return string(result)
	}

	return string(result[1:])
}
