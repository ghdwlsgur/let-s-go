package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		wr.WriteString(strconv.Itoa(i + 1))
		wr.WriteByte('\n')
	}
}
