package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if _, isOpening := mapping[char]; isOpening {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	// Contoh penggunaan
	s1 := "()"
	result1 := isValid(s1)
	fmt.Println("Output 1:", result1)

	s2 := "()[]{}"
	result2 := isValid(s2)
	fmt.Println("Output 2:", result2)

	s3 := "(]"
	result3 := isValid(s3)
	fmt.Println("Output 3:", result3)
}
