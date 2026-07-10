package solution

import (
	"strconv"
)

func isPalindromeStr(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)

	for i, j := 0, len(str)-1; i <= j; {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}
