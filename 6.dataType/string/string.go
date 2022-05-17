package main

import "fmt"

/*
	엄밀히 말하자면 Go에는 문자(character)타입이 없다. 문자를 표현하려면 정수 타입인 rune(int32의 별칭)
	8으로 문자의 코드값을 사용해야 한다.

	문자는 작은 따옴표(')로 감싸서 표현한다. 문자 'A'의 아스키 코드값은 65이고, 문자 'A'를 16진수
	로 작성하면 41, 8진수로는 101이다. 따라서 문자 'A'는 다음과 같이 나타날 수 있다.

	var ch1 = 'A'
	var ch2 = 65
	var ch3 = '\x41'
	var ch4 = '\101'

	16진수로 표현할 때는 앞에 \x를 붙이고, 8진수로 표현할 때는 앞에 \를 붙인다.

	유니코드(UTF-8)의 코드값으로 문자를 표현할 때는 앞에 \u나 \U를 붙인다. 코드 포인트가 4자리인
	유니코드 문자는 \u를 붙이고, 코드 포인트가 8자리인 유니코드 문자는 \U를 붙인다.
*/

func main() {
	var ch1 int = '\u0041'
	var ch2 int = '\uAC00'
	var ch3 int = '\U00101234'
	fmt.Printf("%8d - %8d - %8d\n", ch1, ch2, ch3) // 정수
	fmt.Printf("%8c - %8c - %8c\n", ch1, ch2, ch3) // 문자
	fmt.Printf("%8X - %8X - %8X\n", ch1, ch2, ch3) // UTF-8 바이트 수
	fmt.Printf("%8U - %8U - %8U\n", ch1, ch2, ch3) // UTF-8 코드값
	// fmt.Printf("%")
}
