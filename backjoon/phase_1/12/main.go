package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Printf("%d\n", (a+b)%c)
	fmt.Printf("%d\n", ((a%c)+(b%c))%c)
	fmt.Printf("%d\n", (a*b)%c)
	fmt.Printf("%d\n", ((a%c)*(b%c))%c)
}
