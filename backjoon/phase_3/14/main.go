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

	var n, count int = 0, 1
	fmt.Fscanln(rd, &n)

	result := reg(n)
	for result != n {
		result = reg(result)
		count++
	}
	fmt.Fprint(wr, count)
}

func reg(n int) int {
	val := (n / 10) + (n % 10)
	nn := (n%10)*10 + (val % 10)
	return nn
}
