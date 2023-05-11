package _09_longest_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, 7, longestPalindrome("abccccdd"))
	assert.Equal(t, 1, longestPalindrome("aA"))
}
