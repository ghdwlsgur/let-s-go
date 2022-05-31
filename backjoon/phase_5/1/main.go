package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(arr))

}

func sum(a []int) int {
	var tot int
	for _, v := range a {
		tot += v
	}
	return tot
}
