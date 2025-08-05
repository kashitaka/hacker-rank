package _0_valid_parentheses

func isValid(s string) bool {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		switch c {
		case ')':
			if pop != '(' {
				return false
			}
		case '}':
			if pop != '{' {
				return false
			}
		case ']':
			if pop != '[' {
				return false
			}
		}
	}
	return len(stack) == 0
}
