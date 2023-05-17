package _99_bulls_and_cows

import "fmt"

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	secretCount, guessCount := make(map[uint8]int),make(map[uint8]int)
	for i:=0; i<len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
			continue
		}
		secretCount[secret[i]]++
		guessCount[guess[i]]++
	}
	for k, secCnt := range secretCount {
		guessCnt, ok := guessCount[k]
		if ok {
			cows += min(secCnt, guessCnt)
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}