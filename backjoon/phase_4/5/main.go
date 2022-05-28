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

	var n, max, tot int
	var avg float64
	fmt.Scanln(&n)

	line, _, _ := rd.ReadLine()
	input := bytes.Fields(line)

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(string(input[i]))
		tot += num
		if max < num {
			max = num
		}
	}
	avg = ((float64(tot) / float64(max)) * 100)
	fmt.Fprintln(wr, avg/float64(n))
}

func test() {
	var rd *bufio.Reader = bufio.NewReader(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, max, tot int
	var avg float64
	fmt.Scanln(&n)

	line, _, _ := rd.ReadLine()
	input := bytes.Fields(line)

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(string(input[i]))
		tot += num
		if max < num {
			max = num
		}
	}

	avg = (float64(tot) / float64(max)) * 100
	fmt.Fprintln(wr, avg/float64(n))
}
