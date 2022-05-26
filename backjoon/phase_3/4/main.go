package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

func main() {
	bf := 1 << 24
	rd := bufio.NewReaderSize(os.Stdin, bf)
	wr := bufio.NewWriterSize(os.Stdout, bf)
	defer wr.Flush()

	one, _, _ := rd.ReadLine()
	n, _ := strconv.Atoi(string(one))

	for i := 1; i <= n; i++ {
		two, _, _ := rd.ReadLine()
		line := bytes.Fields(two)

		a, _ := strconv.Atoi(string(line[0]))
		b, _ := strconv.Atoi(string(line[1]))

		wr.WriteString(strconv.Itoa(a + b))
		wr.WriteByte('\n')
	}
}
