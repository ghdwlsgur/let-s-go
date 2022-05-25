package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if a == b && b == c && a == c {
		fmt.Printf("%d", 10000+(a*1000))
		return
	} else if a == b || b == c {
		fmt.Printf("%d", 1000+b*100)
		return
	} else if a == c {
		fmt.Printf("%d", 1000+c*100)
		return
	}

	if (a > b) && (a > c) {
		fmt.Printf("%d", a*100)
	} else if (b > a) && (b > c) {
		fmt.Printf("%d", b*100)
	} else {
		fmt.Printf("%d", c*100)
	}
}
