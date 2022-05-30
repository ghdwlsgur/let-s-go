package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var rd *bufio.Reader = bufio.NewReader(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, num int
	var avg float64

	fmt.Fscanln(rd, &n)

	for i := 0; i < n; i++ {
		raw, _, _ := rd.ReadLine()
		input := bytes.Fields(raw)

		arr := []int{}
		cnt := 0
		tot := 0

		for i, v := range input {
			num, _ = strconv.Atoi(string(input[0]))
			arr := []int{}
			if i != 0 {
				n, _ := strconv.Atoi(string(v))
				arr = append(arr, n)
				tot += n
			}
		}
		avg = float64(tot) / float64(num)

		for _, v := range arr {
			if avg < float64(v) {
				cnt++
			}
		}
		result := float64(cnt) / float64(num) * 100
		fmt.Fprintf(wr, "%.3f%s\n", result, "%")
	}
}
