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

	var n, k, result int = 0, 0, 1

	for i := 0; i < 3; i++ {
		fmt.Fscanln(rd, &n)
		result *= n
	}

	arr := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for result > 0 {
		k = result % 10
		result = result / 10
		arr[k]++
		// switch {
		// case k == 0:
		// 	arr[0]++
		// case k == 1:
		// 	arr[1]++
		// case k == 2:
		// 	arr[2]++
		// case k == 3:
		// 	arr[3]++
		// case k == 4:
		// 	arr[4]++
		// case k == 5:
		// 	arr[5]++
		// case k == 6:
		// 	arr[6]++
		// case k == 7:
		// 	arr[7]++
		// case k == 8:
		// 	arr[8]++
		// case k == 9:
		// 	arr[9]++
		// default:
		// 	break
		// }
	}

	for _, v := range arr {
		fmt.Fprintln(wr, v)
	}
}
