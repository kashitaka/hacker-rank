package _51_Reverse_Words_in_a_String_

func reverseWords(s string) string {
	stack := []string{""}
	for _, v := range s {
		if v != ' ' {
			stack[len(stack)-1]  = stack[len(stack)-1] + string(v)
			continue
		}
		pop := stack[len(stack)-1]
		if pop == "" {
			continue
		}
		stack = append(stack, "")
	}
	if stack[len(stack)-1] == "" {
		stack = stack[:len(stack)-1]
	}

	res := ""
	for i:= len(stack)-1; i >= 0; i-- {
		res += stack[i] + " "
	}
	return res[:len(res)-1]
}