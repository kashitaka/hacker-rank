package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'lonelyinteger' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func main() {
	input := inputToIntSlice()
	t := input[0][0]
	for i := 0; i < t; i++ {
		in := input[2*i+2]
		findZigZagSequence(in)
	}
}

func findZigZagSequence(input []int) {
	n := len(input)
	out := make([]int, n)
	sort.Ints(input)
	k := n / 2
	for i := 0; i < n; i++ {
		if i < k {
			out[i] = input[i]
		} else if i == k {
			out[i] = input[n-1]
		} else {
			out[i] = input[n-1-(i-k)]
		}
	}
	printSlice(out)
}

func printSlice(in []int) {
	s := fmt.Sprint(in)
	fmt.Println(s[1 : len(s)-1])
}

func inputToIntSlice() [][]int {
	var res [][]int
	var sc = bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 65536) // max 65536. change if input may be longer
	for sc.Scan() {
		line := sc.Text()
		ss := strings.Split(line, " ")
		intSlice := make([]int, len(ss))
		for i, v := range ss {
			intV, _ := strconv.Atoi(v)
			intSlice[i] = intV
		}
		res = append(res, intSlice)
	}
	return res
}
