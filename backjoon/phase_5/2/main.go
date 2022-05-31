package main

import "fmt"

func main() {
	for i := 20; i < 100; i++ {
		// if kaprekar(i) != i {

		// }
		fmt.Println(kaprekar(i))
	}
	// fmt.Println(kaprekar(11))
}

func kaprekar(n int) int {
	var a int = 1
	var self int

	self += n

	for {
		if n/a > 0 {
			a *= 10
		} else {
			break
		}
	}
	for a > 10 {
		self += n / a
		if n%a > 10 {
			n %= a
		}
		a /= 10
	}

	self += n%a + n/a
	return self
}
