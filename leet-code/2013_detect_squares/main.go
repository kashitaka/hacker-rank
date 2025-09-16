package _013_detect_squares

type DetectSquares struct {
	hash map[int]map[int]int // x:y:num of points
}

func Constructor() DetectSquares {
	return DetectSquares{
		hash: make(map[int]map[int]int),
	}
}

func (d *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if _, ok := d.hash[x]; !ok {
		d.hash[x] = make(map[int]int)
	}
	d.hash[x][y]++
}

func (d *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	res := 0
	ys := d.hash[x]
	for y1, _ := range ys {
		l := dist(y, y1)
		if l == 0 {
			continue
		}
		// left
		res += d.hash[x][y1] * d.hash[x-l][y1] * d.hash[x-l][y]

		// right
		res += d.hash[x][y1] * d.hash[x+l][y1] * d.hash[x+l][y]
	}
	return res
}

func dist(i, j int) int {
	res := i - j
	if res < 0 {
		return -res
	}
	return res
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
