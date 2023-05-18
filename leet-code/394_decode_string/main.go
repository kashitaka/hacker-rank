package _94_decode_string

import "strconv"

func decodeString(s string) string {
	var stack []string
	for i:=0; i<len(s); i++ {
		char := s[i:i+1]
		if char != "]" {
			stack = append(stack, char)
			continue
		}
		// if "]"
		var repeatedString string
		var pop string
		// make string to repeat
		for{
			pop = stack[len(stack)-1:len(stack)][0]
			stack = stack[0:len(stack)-1]
			if pop == "[" {
				break
			}
			repeatedString = pop + repeatedString
		}
		// make repeat count
		var repeatNum string
		for len(stack) > 0{
			pop = stack[len(stack)-1:len(stack)][0]
			_, err := strconv.Atoi(pop)
			if err != nil {
				break
			}
			stack = stack[0:len(stack)-1]
			repeatNum = pop + repeatNum
		}
		num, _ := strconv.Atoi(repeatNum)
		addToStack := ""
		for j := 0; j < num; j++ {
			addToStack += repeatedString
		}
		stack = append(stack, addToStack)
	}
	var res string
	for _, v := range stack {
		res += v
	}
	return res
}