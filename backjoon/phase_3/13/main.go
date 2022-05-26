package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var rd *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var a, b int
	var set []string

	for {
		if rd.Scan() {
			set = strings.Fields(rd.Text())
			a, _ = strconv.Atoi(set[0])
			b, _ = strconv.Atoi(set[1])
			fmt.Fprintln(wr, a+b)
		} else {
			break
		}
	}
}

func typing() {
	var rd *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var a, b int
	var set []string
	for {
		if rd.Scan() {
			set = strings.Fields(rd.Text())
			a, _ = strconv.Atoi(set[0])
			b, _ = strconv.Atoi(set[1])
			fmt.Fprintln(wr, a+b)
		} else {
			break
		}
	}
}

func typing2() {
	var rd *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)

	var a, b int
	var set []string

	for {
		if rd.Scan() {
			set = strings.Fields(rd.Text())
			a, _ = strconv.Atoi(set[0])
			b, _ = strconv.Atoi(set[1])
			fmt.Fprintln(wr, a+b)
		} else {
			break
		}
	}
}
