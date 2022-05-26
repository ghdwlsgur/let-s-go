package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var rd *bufio.Reader = bufio.NewReaderSize(os.Stdin, 1)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscanln(rd, &n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			wr.WriteString("*")
		}
		wr.WriteByte('\n')
	}
}
