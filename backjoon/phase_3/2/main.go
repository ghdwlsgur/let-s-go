package main

import "fmt"

func main() {
	var n int
	var a, b int

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &a, &b)
		fmt.Printf("%d\n", a+b)
	}

}
