package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var rd *bufio.Reader = bufio.NewReaderSize(os.Stdin, 1)

	var n, a, b int
	fmt.Fscanln(rd, &n)
	var wr *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1<<n*10)
	defer wr.Flush()

	for i := 1; i <= n; i++ {
		fmt.Fscanln(rd, &a, &b)
		tmp := fmt.Sprint("Case #", i, ": ", a, " + ", b, " = ", a+b)
		fmt.Fprintln(wr, tmp)
	}
}
