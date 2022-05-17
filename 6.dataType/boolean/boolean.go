package main

import "fmt"

func main() {

	v := 10

	b1 := true
	b2 := v == 5
	b3 := v == 10
	b4 := b1 && b2
	b5 := b1 || b2
	b6 := !b1

	fmt.Println(b1, b2, b3, b4, b5, b6)

}
