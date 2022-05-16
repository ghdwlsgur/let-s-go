package main

import "fmt"

func main() {

	// 상수는 불, 숫자, 문자열 타입으로만 선언 가능
	const limit = 64
	fmt.Println(limit)

	// 상수를 선언할 때 타입은 표기하지 않아도 된다.
	// 할당되는 값의 타입에 따라 컴파일러가 자동으로 상수의 타입을 결정한다.
	// const max uint64 = 1024
	const max = 1024 * 1024
	fmt.Println(max)

	// 여러 개 선언
	const (
		RED    = 0
		ORANGE = 1
		YELLOW = 2
	)

	fmt.Println(RED, ORANGE, YELLOW)
}
