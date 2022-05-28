package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var rd *bufio.Reader = bufio.NewReader(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, sum, addCnt int
	var input string
	fmt.Fscanln(rd, &n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(rd, &input)
		k := strings.Fields(input)

		for _, v := range k {
			for i := 0; i < len(v); i++ {
				if v[i] == 79 {
					addCnt++
				} else {
					addCnt = 0
				}
				sum += addCnt
			}
			wr.WriteString(strconv.Itoa(sum))
			wr.WriteByte('\n')
			sum = 0
			addCnt = 0
		}
	}

}
