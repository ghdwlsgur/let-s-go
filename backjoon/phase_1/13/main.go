package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d\n%d", &a, &b)
	fmt.Printf("%d\n", a*(b%10))
	fmt.Printf("%d\n", a*((b%100)/10))
	fmt.Printf("%d\n", a*(b/100))
	fmt.Printf("%d\n", a*b)
}
