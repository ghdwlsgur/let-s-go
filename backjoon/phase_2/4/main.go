package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanf("%d\n%d", &x, &y)
	if x > 0 && y > 0 {
		fmt.Print("1")
		return
	} else if x < 0 && y > 0 {
		fmt.Print("2")
		return
	} else if x < 0 && y < 0 {
		fmt.Print("3")
		return
	}
	fmt.Print("4")
}
