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
	star := "*"

	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			wr.WriteString(" ")
		}
		wr.WriteString(star)
		wr.WriteByte('\n')
		star += "*"
	}
}
