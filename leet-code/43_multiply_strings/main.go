package _3_multiply_strings

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	intNum1 := make([]int, len(num1))
	intNum2 := make([]int, len(num2))
	for i, v := range num1 {
		intNum1[i] = int(v - '0')
	}
	for i, v := range num2 {
		intNum2[i] = int(v - '0')
	}

	res := make([]int, len(intNum1)+len(intNum2))

	for i := len(intNum1) - 1; i >= 0; i-- {
		n1 := intNum1[i]
		for j := len(intNum2) - 1; j >= 0; j-- {
			n2 := intNum2[j]
			pos := i + j + 1
			res[pos] += n1 * n2
			res[pos-1] += res[pos] / 10
			res[pos] = res[pos] % 10
		}
	}

	finishRemoveZeros := false
	resByte := []byte{}
	for _, v := range res {
		if v == 0 && !finishRemoveZeros {
			continue
		}
		if !finishRemoveZeros {
			finishRemoveZeros = true
		}
		resByte = append(resByte, byte(v)+'0')
	}
	return string(resByte)
}
