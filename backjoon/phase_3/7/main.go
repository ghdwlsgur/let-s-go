package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// rd := bufio.NewReaderSize(os.Stdin, 1)
	// line, _, _ := rd.ReadLine()
	// n, _ := strconv.Atoi(string(line))

	// wr := bufio.NewWriterSize(os.Stdout, 1<<n*10)
	// defer wr.Flush()

	// for i := 0; i < n; i++ {
	// 	raw, _, _ := rd.ReadLine()
	// 	l := bytes.Fields(raw)

	// 	a, _ := strconv.Atoi(string(l[0]))
	// 	b, _ := strconv.Atoi(string(l[1]))

	// 	wr.WriteString("Case #")
	// 	wr.WriteString(strconv.Itoa(i + 1))
	// 	wr.WriteString(": ")
	// 	wr.WriteString(strconv.Itoa(a + b))
	// 	wr.WriteByte('\n')
	// }

	// BufferedSize := wr.Buffered()
	// fmt.Printf("Buffered size : %d\n", BufferedSize)

	// var rd *bufio.Reader
	// var wr *bufio.Writer
	// rd = bufio.NewReaderSize(os.Stdin, 1)

	// var n, a, b int
	// fmt.Fscanln(rd, &n)
	// wr = bufio.NewWriterSize(os.Stdout, 1<<n*10)
	// defer wr.Flush()

	// for i := 0; i < n; i++ {
	// 	fmt.Fscanln(rd, &a, &b)
	// 	tmp := fmt.Sprint("Case #", i+1, ": ", a+b)
	// 	fmt.Fprintln(wr, tmp)
	// }

	var rd *bufio.Reader = bufio.NewReaderSize(os.Stdin, 1)

	var n, a, b int
	fmt.Fscanln(rd, &n)
	var wr *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1<<n*10)
	defer wr.Flush()

	for i := 1; i <= n; i++ {
		fmt.Fscanln(rd, &a, &b)
		tmp := fmt.Sprint("Case #", i, ": ", a+b)
		fmt.Fprintln(wr, tmp)
	}
}

func refactoring() {
	var rd *bufio.Reader
	var wr *bufio.Writer
	rd = bufio.NewReaderSize(os.Stdin, 1)

	var n, a, b int
	fmt.Fscanln(rd, &n)
	wr = bufio.NewWriterSize(os.Stdout, 1<<n*10)
	defer wr.Flush()

	for i := 0; i < n; i++ {
		fmt.Fscanln(rd, &a, &b)
		tmp := fmt.Sprint("Case #", i+1, ": ", a+b)
		fmt.Fprintln(wr, tmp)
	}
}

func refactoring2() {
	var rd *bufio.Reader = bufio.NewReaderSize(os.Stdin, 1)

	var n, a, b int
	fmt.Fscanln(rd, &n)
	var wr *bufio.Writer = bufio.NewWriterSize(os.Stdout, 1<<n*10)
	defer wr.Flush()

	for i := 1; i <= n; i++ {
		fmt.Fscanln(rd, &a, &b)
		tmp := fmt.Sprint("Case #", i, ": ", a+b)
		fmt.Fprintln(wr, tmp)
	}
}
