package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var rd *bufio.Reader = bufio.NewReader(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, index int
	var max int = 1

	for i := 0; i < 9; i++ {
		fmt.Fscanln(rd, &n)
		if max < n {
			max = n
			index = i + 1
		}
	}
	wr.WriteString(strconv.Itoa(max))
	wr.WriteByte('\n')
	wr.WriteString(strconv.Itoa(index))
}

// 4ms
func test1() {
	var rd *bufio.Reader = bufio.NewReader(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var arr []int
	var n int

	for i := 0; i < 9; i++ {
		fmt.Fscanln(rd, &n)
		arr = append(arr, n)
	}

	var max int = 1
	var index int

	for i, v := range arr {
		if max < v {
			max = v
			index = i + 1
		}
	}
	wr.WriteString(strconv.Itoa(max))
	wr.WriteByte('\n')
	wr.WriteString(strconv.Itoa(index))
}
