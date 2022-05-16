package main

import "fmt"

func main() {
	// const (
	// 	Sunday   = 0
	// 	Monday   = 1
	// 	Tuesday  = 2
	// 	Thursday = 3
	// 	Friday   = 4
	// 	Saturday = 5
	// )

	// fmt.Println(Sunday, Monday, Tuesday, Thursday, Friday, Saturday)

	/* 열거형
	상수를 열겨형으로 선언할 때 iota 예약어를 사용하면 편리하다. 상수를 그룹으로 묶어서 선언할 때
	const 그룹에서 iota의 값은 0이고, 이후로는 1씩 증가한다.
	*/

	const (
		Sunday = iota
		Monday
		Tuesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Thursday, Friday, Saturday)

	type Color int

	const (
		RED Color = iota
		ORANGE
		YELLOW
		GREEN
	)

	fmt.Println(RED, ORANGE, YELLOW, GREEN)

	type ByteSize int64

	const (
		_           = iota             // 첫줄 무시 // iota = 0
		KB ByteSize = 1 << (10 * iota) // 1을 비트 왼쪽으로 10칸 즉, 2**10, iota = 1
		MB                             // 1을 비트 왼쪽으로 20칸 즉, 2**20, iota = 2
		GB                             // 2**30, iota = 3
		TB                             // 2**40, iota = 4
		PB                             // 2**50, iota = 5
		EB                             // 2**60, iota = 6
	)

	fmt.Println(KB, MB, GB, TB, PB, EB)

	const (
		_            = iota
		DEFAULT_RATE = 5 + 0.3*iota
		GREEN_RATE
		SILVER_RATE
		GOLD_RATE
	)

	fmt.Println(DEFAULT_RATE, GREEN_RATE, SILVER_RATE, GOLD_RATE)
}
