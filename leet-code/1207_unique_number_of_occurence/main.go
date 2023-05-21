package _207_unique_number_of_occurence

func uniqueOccurrences(arr []int) bool {
	freqByNum := make(map[int]int)
	for _, v := range arr {
		freqByNum[v]++
	}
	numOccurrence := make(map[int]struct{})
	for _, v := range freqByNum {
		_, ok := numOccurrence[v]
		if ok {
			return false
		}
		numOccurrence[v] = struct{}{}
	}
	return true
}