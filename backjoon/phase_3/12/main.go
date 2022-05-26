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

	var a, b int = -1, -1
	for i := 0; (a != 0) || (b != 0); i++ {
		fmt.Fscanln(rd, &a, &b)
		if (a != 0) || (b != 0) {
			fmt.Fprintln(wr, a+b)
		}
	}
}

func typing() {
	var rd *bufio.Reader = bufio.NewReader(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var a, b int = -1, -1
	for i := 0; (a != 0) || (b != 0); i++ {
		fmt.Fscanln(rd, &a, &b)
		if (a != 0) || (b != 0) {
			fmt.Fprintln(wr, a+b)
		}
	}
}

func typing2() {
	var rd *bufio.Reader = bufio.NewReader(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var a, b int = -1, -1
	for i := 0; (a != 0) || (b != 0); i++ {
		fmt.Fscanln(rd, &a, &b)
		if (a != 0) || (b != 0) {
			fmt.Fprintln(wr, a+b)
		}
	}
}
