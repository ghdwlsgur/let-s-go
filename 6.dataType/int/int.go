package main

import "fmt"

/*
	Go는 정수 타입과 문자 타입을 구분하지 않는다. byte는 또는 rune 타입으로 문자의 코드값을 저장하여
	문자를 표기한다. byte는 uint8의 별칭으로 1바이트로 표현할 수 있는 아스키 문자를 표기할 수 있다.
	rune은 int32의 별칭으로 유니코드(UTF-8) 문자를 표기할 수 있다.

	! - byte = uint8 (1바이트)
	! - rune = int32 (4바이트)
*/
func main() {
	a := 365   // 10진수
	b := 0555  // 8진수 (8^2 * 5) + (8^1 * 5) + (8^0 * 5)
	c := 0x16D // 16진수 (16^2 * 1) + (16^1 * 6) + (16^0 * 13)

	fmt.Println(a, b, c)

	var ch1 byte = 65   // 65 (10진수)
	var ch2 byte = 0101 // 65 (8진수)
	var ch3 byte = 0x41 // 65 (16진수)

	var ch4 rune = 44032   // 44032 (10진수)
	var ch5 rune = 0126000 // 44032 (8진수)
	var ch6 rune = 0xAC00  // 44032 (16진수)

	fmt.Printf("%c %c %c\n", ch1, ch2, ch3)
	fmt.Printf("%c %c %c\n", ch4, ch5, ch6)

	// 문자 코드값이 아니라 문자 자체를 저장할 때는 작은 따옴표를 사용한다.
	var ch7 byte = 'A'
	var ch8 rune = '가'

	fmt.Printf("%c %c \n", ch7, ch8)
}
