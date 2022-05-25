package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	if n >= 90 && n <= 100 {
		fmt.Println("A")
	} else if n >= 80 && n <= 89 {
		fmt.Println("B")
	} else if n >= 70 && n <= 79 {
		fmt.Println("C")
	} else if n >= 60 && n <= 69 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
