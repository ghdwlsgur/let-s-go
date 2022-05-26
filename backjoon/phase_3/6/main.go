package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

	rd := bufio.NewReaderSize(os.Stdin, 1)
	line, _, _ := rd.ReadLine()

	n, _ := strconv.Atoi(string(line))
	wr := bufio.NewWriterSize(os.Stdout, n)

	for i := n; i >= 1; i-- {
		wr.WriteString(strconv.Itoa(i))
		wr.WriteByte('\n')
	}

	defer wr.Flush()
}
