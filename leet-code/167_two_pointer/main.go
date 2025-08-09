package _67_two_pointer

// Time O(n)
// Space O(1)
func twoSumTwoPointer(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if sum < target {
			// need to be larger. so shift l to right
			l++
		} else {
			// otherwise, need to be smaller. so shift r to left
			r--
		}
	}
	panic("error")
}

// Time O(nlogn)
// Space O(1)
func twoSumBinarySearch(numbers []int, target int) []int {
	var binarySearch func(l int, target int) int
	binarySearch = func(l int, t int) int {
		r := len(numbers) - 1
		for l <= r {
			mid := (l + r) / 2
			if numbers[mid] == t {
				return mid
			}
			if numbers[mid] < t {
				// search right
				l = mid + 1
			} else {
				// search left
				r = mid - 1
			}
		}
		return -1
	}

	for idx1, num := range numbers {
		sub := target - num
		idx2 := binarySearch(idx1+1, sub)
		if idx2 != -1 {
			return []int{idx1 + 1, idx2 + 1}
		}
	}

	panic("error")
}
