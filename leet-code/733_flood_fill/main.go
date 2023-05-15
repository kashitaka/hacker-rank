package _33_flood_fill

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	m := len(image)
	n := len(image[0])
	targetColor := image[sr][sc]
	if color == targetColor {
		return image
	}
	fill(image, m, n, sr, sc, targetColor, color)
	return image
}

func fill(image [][]int, m, n, i, j, targetColor, color int) {
	if image[i][j] != targetColor {
		return
	}
	image[i][j] = color
	if i != 0 {
		fill(image, m, n, i-1, j, targetColor, color)
	}
	if j != 0 {
		fill(image, m, n, i, j-1, targetColor, color)
	}
	if i != m-1 {
		fill(image, m, n, i+1, j, targetColor, color)
	}
	if j != n-1 {
		fill(image, m, n, i, j+1, targetColor, color)
	}
}
