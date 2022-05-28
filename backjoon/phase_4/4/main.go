package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var rd *bufio.Reader = bufio.NewReader(os.Stdin)
	var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var arr = make([]int, 10)
	var comparison = make([]int, len(arr))

	var n, cnt int = 0, 1
	for i := 0; i < 10; i++ {
		fmt.Fscanln(rd, &n)
		k := n % 42
		arr[i] += k
	}

	fmt.Println(arr)
	// 39 40 41 0 1 2 40 41 0 1
	// [k] [k]
	// 39 40 41 0 1 2 40 41 0 1
	// 40
	// 39 41 0 1 2 41 0 1
	// 41
	// 39 0 1 2 0 1
	// 0
	// 39 1 2 1
	// 1
	// 39 2

	// 39
	//
	// 39 -> 40 41 0 1 2 ...
	// 40 -> 41 0 1 2 40 41 0 1
	// 41 -> 0 1 2 40 41 0 1
	//

	copy(comparison, arr)
	fmt.Println(comparison)

	for i,v := range arr {
		for j := i + 1; i <= 9 j++ {
			
		}
	}
	// arr[]

	// for i, v := range arr {
	// 	fmt.Println("before", arr)
	// 	for j := i + 1; j <= 9; j++ {
	// 		if v == arr[j] && j != 9 {
	// 			arr = append(arr[:j], arr[j+1:]...)
	// 			arr = append(arr[:i], arr[i+1:]...)
	// 			j--
	// 			i--
	// 			fmt.Println("after", arr)
	// 		} else if j == 9 {
	// 			arr = append(arr[:j], arr[len(arr)-1:]...)
	// 		}
	// 	}
	// 	cnt++
	// }

	wr.WriteString(strconv.Itoa(cnt))
}
