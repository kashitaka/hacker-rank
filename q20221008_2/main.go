package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var mem map[int32][]int32

/*
 * Complete the 'numberOfItems' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER_ARRAY startIndices
 *  3. INTEGER_ARRAY endIndices
 */
func numberOfItems(s string, startIndices []int32, endIndices []int32) []int32 {
	fmt.Println(startIndices, endIndices)
	res := make([]int32, len(startIndices))
	for i := 0; i < len(startIndices); i++ {
		start := startIndices[i]
		end := endIndices[i]
		res[i] = count(s, start, end)
	}
	return res
}

func createMem(s string, mem map[int32][]int32) {
	prev := int32(0)
	end := int32(0)
	start := false
	for i, v := range s {
		if v == 124 {
			// | が見つかった。
			// 前回のprevからのitem数を格納する
			end = int32(i)
			j := prev
			if !start {
				for j <= end {
					mem[j] = []int32{end, 0}
					j++
				}
			} else {
				mem[prev] = []int32{end, end - prev - 1}
				j++
				for j <= end {
					mem[j] = []int32{end, 0}
					j++
				}
			}
			start = true
			prev = end
		}
	}
	// 最後のend以降の処理
	for end < int32(len(s)) {
		mem[end] = []int32{math.MaxInt32, 0}
		end++
	}
}

func count2(startIdx int32, endIdx int32, mem map[int32][]int32) int32 {
	v := mem[startIdx]
	if v[0] > endIdx {
		return 0
	}
	return v[1] + count2(v[0], endIdx, mem)
}

func count(s string, startIdx int32, endIdx int32) int32 {
	res := int32(0)
	split := s[startIdx-1 : endIdx]
	start := false
	tmp := int32(0)
	for _, v := range split {
		if v == 124 {
			start = true
		}
		if v == 42 && start {
			tmp++
		}
		if v == 124 && start {
			res += tmp
			tmp = 0
		}
	}
	fmt.Println(res)
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	startIndicesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var startIndices []int32

	for i := 0; i < int(startIndicesCount); i++ {
		startIndicesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		startIndicesItem := int32(startIndicesItemTemp)
		startIndices = append(startIndices, startIndicesItem)
	}

	endIndicesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var endIndices []int32

	for i := 0; i < int(endIndicesCount); i++ {
		endIndicesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		endIndicesItem := int32(endIndicesItemTemp)
		endIndices = append(endIndices, endIndicesItem)
	}

	result := numberOfItems(s, startIndices, endIndices)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
