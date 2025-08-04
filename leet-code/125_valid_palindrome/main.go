package _25_valid_palindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for i <= j {
		if !isAlphanumericRune(rune(s[i])) {
			i++
			continue
		}
		if !isAlphanumericRune(rune(s[j])) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumericRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
