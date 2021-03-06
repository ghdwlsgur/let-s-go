package main

import "fmt"

func enter(s string) string {
	fmt.Println("entering:", s)
	return s
}

func leave(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer leave(enter("a"))
	fmt.Println("in a")
}

func b() {
	defer leave(enter("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
}

/* 예상 출력 결과
entering: b
in b
entering: a
in a
leaving: a
leaving: b
*/
