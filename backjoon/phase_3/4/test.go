package test

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

func test() {
	rd := bufio.NewReaderSize(os.Stdin, 1<<24)
	wr := bufio.NewWriterSize(os.Stdout, 1<<24)
	defer wr.Flush()

	raw, _, _ := rd.ReadLine()
	t, _ := strconv.Atoi(string(raw))

	for tc := 1; tc <= t; tc++ {
		raw, _, _ = rd.ReadLine()
		line := bytes.Fields(raw)
		a, _ := strconv.Atoi(string(line[0]))
		b, _ := strconv.Atoi(string(line[1]))

		wr.WriteString(strconv.Itoa(a + b))
		wr.WriteByte('\n')
	}
}

func test2() {
	rd := bufio.NewReaderSize(os.Stdin, 1<<24)
	wr := bufio.NewWriterSize(os.Stdout, 1<<24)
	defer wr.Flush()

	raw, _, _ := rd.ReadLine()
	t, _ := strconv.Atoi(string(raw))

	for tc := 1; tc <= t; tc++ {
		raw, _, _ := rd.ReadLine()
		line := bytes.Fields(raw)
		a, _ := strconv.Atoi(string(line[0]))
		b, _ := strconv.Atoi(string(line[1]))

		wr.WriteString(strconv.Itoa(a + b))
		wr.WriteByte('\n')
	}
}
