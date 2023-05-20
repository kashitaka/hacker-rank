package _431_Kids_With_the_Greatest_Number_of_Candies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// find max O(N)
	max := 0
	for _, v := range candies {
		if v > max {
			max = v
		}
	}
	res := make([]bool, len(candies))

	// O(N)
	for i, v := range candies {
		if v + extraCandies >=  max {
			res[i] = true
			continue
		}
		res[i] = false
	}
	return res
}