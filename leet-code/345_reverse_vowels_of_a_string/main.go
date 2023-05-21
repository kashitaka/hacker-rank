package _45_reverse_vowels_of_a_string

func reverseVowels(s string) string {
	byteS := []byte(s)
	i, j := 0, len(byteS)-1
	for i < j {
		if isVowels(byteS[i]) {
			for i < j {
				if isVowels(byteS[j]) {
					byteS[i], byteS[j] = byteS[j], byteS[i]
					j--
					break
				}
				j--
			}
			i++
			continue
		}
		i++
	}
	return string(byteS)
}

func isVowels(char uint8) bool {
	return char == 65 || char == 69 || char == 73 || char == 79 || char == 85 ||
		char == 97 || char == 101 || char == 105 || char == 111 || char == 117
}