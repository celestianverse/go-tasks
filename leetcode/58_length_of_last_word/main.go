package main

import "fmt"

func main() {
	input := "   fly me   to   the moon  "

	fmt.Println(lengthOfLastWord(input))
}

func lengthOfLastWord(s string) int {
	found := false
	start := 0
	end := 0

	for i := len(s) - 1; i >= 0; i-- {
		space := s[i] == ' '

		if !found && !space {
			found = true
			end = i + 1
		}

		if found {
			if space {
				start = i + 1

				return end - start
			}
			if i == 0 {
				return end
			}
		}
	}

	return 0
}
