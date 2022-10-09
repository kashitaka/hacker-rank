package main

func palindromeIndex(s string) int32 {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			if isPalindrome(s[0:i] + s[i+1:]) {
				return int32(i)
			} else {
				return int32(len(s) - 1 - i)
			}
		}
	}
	return -1
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
