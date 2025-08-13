package _4_search_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	t, b := 0, len(matrix)-1
	search := -1
	for t <= b {
		mid := (t + b) / 2
		if matrix[mid][0] <= target && target <= matrix[mid][len(matrix[mid])-1] {
			search = mid
			break
		}
		if target < matrix[mid][0] {
			b = mid - 1
		} else {
			t = mid + 1
		}
	}
	if search == -1 {
		return false
	}

	// t is the row to search
	l, r := 0, len(matrix[0])-1
	for l <= r {
		mid := (l + r) / 2
		if matrix[search][mid] == target {
			return true
		}
		if matrix[search][mid] < target {
			// search right
			l = mid + 1
		} else {
			// search left
			r = mid - 1
		}
	}
	return false
}
