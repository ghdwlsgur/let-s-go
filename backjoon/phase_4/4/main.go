package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var rd *bufio.Reader = bufio.NewReader(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var arr = make([]int, 42)

	var n, cnt int
	for i := 0; i < 10; i++ {
		fmt.Fscanln(rd, &n)
		arr[n%42]++
	}

	for i := range arr {
		if arr[i] > 0 {
			cnt++
		}
	}
	fmt.Fprintln(wr, cnt)
}
