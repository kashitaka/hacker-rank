package main

const (
	C = 67
	D = 68
	M = 77
	X = 88
	L = 76
	I = 73
	V = 86
)

func romanToInt(s string) int {
	res := 0
	i := 0
	for i < len(s) {
		if i+1 < len(s) {
			if s[i] == C && s[i+1] == D {
				res += 400
				i += 2
				continue
			} else if s[i] == C && s[i+1] == M {
				res += 900
				i += 2
				continue
			} else if s[i] == X && s[i+1] == L {
				res += 40
				i += 2
				continue
			} else if s[i] == X && s[i+1] == C {
				res += 90
				i += 2
				continue
			} else if s[i] == I && s[i+1] == V {
				res += 4
				i += 2
				continue
			} else if s[i] == I && s[i+1] == X {
				res += 9
				i += 2
				continue
			}
		}
		switch s[i] {
		case I:
			res += 1
		case V:
			res += 5
		case X:
			res += 10
		case L:
			res += 50
		case C:
			res += 100
		case D:
			res += 500
		case M:
			res += 1000
		}
		i++
	}
	return res
}
