package _50_evaluate_reserve_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, op1+op2)
		case "-":
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, op1-op2)
		case "*":
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, op1*op2)
		case "/":
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, op1/op2)
		default:
			intToken, _ := strconv.Atoi(token)
			stack = append(stack, intToken)
		}
	}
	return stack[0]
}
