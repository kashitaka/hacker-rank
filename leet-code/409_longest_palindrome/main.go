package _09_longest_palindrome

func longestPalindrome(s string) int {
	charMap := make(map[int32]int)
	result := 0
	for _, v := range s {
		cnt := charMap[v]
		if cnt == 1 {
			charMap[v] = 0
			result += 2
			continue
		}
		charMap[v]++
	}
	for _, v := range charMap {
		if v != 0 {
			result++
			break
		}
	}
	return result
}
