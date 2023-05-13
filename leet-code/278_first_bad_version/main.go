package _78_first_bad_version

func isBadVersion(version int) bool {
	return version == 1
}

func firstBadVersion(n int) int {
	left := 0
	right := n
	for left+1 < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
