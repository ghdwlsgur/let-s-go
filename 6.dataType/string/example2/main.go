package main

import "fmt"

/*
문자열은 유니코드 문자의 코드값을 정수로 표현한 값(rune 또는 int32)의 시퀀스이므로 []rune(또는
[]int32)으로 변환할 수 있다. 1바이트로 표현할 수 있는 아스키 문자열은 []byte(또는 []uint8)로
변환할 수 있다. 아스키가 아닌 문자열을 []byte로 변환하면 잘못된 코드값으로 변환될 수 있으니 주의한다.

그렇다면 왜 int32, rune를 사용할까 ?
그 이유는 go-lang은 문자 인코딩 방식으로 UTF-8을 사용하고 있는데 UTF-8은 한 문자를 나타내기 위해
1바이트에서 4바이트까지를 사용한다. 따라서 int32, 8비트, 4바이트의 데이터 타입을 사용한다.

또한 아스키 코드에서 지원하는 문자는 []byte 또는 []int8 데이터 타입을 사용하여 표시할 수 있다.

- []rune(s)
문자열 s를 유니코드 문자의 코드값 배열([]rune 또는 []int32)로 변환

- []byte(s)
문자열 s를 바이트 배열([]byte 또는 []uint)로 변환

- string(chars)
유니코드 문자의 코드값 배열([]rune 또는 []int32)을 문자열로 변환

- string(i)
유니코드의 코드 포인트 i를 문자열로 변환 (i가 65라면 "A"로 변환)
*/

func main() {
	s1 := "Ahello"

	fmt.Println([]rune(s1))
	fmt.Println([]byte(s1))

	fmt.Println(string([]byte{65, 104, 101, 108, 108, 111}))
	fmt.Println(string([]rune{65, 104, 101, 108, 108, 111}))

	s2 := "안녕하세요"
	fmt.Println([]rune(s2))
	fmt.Println([]byte(s2))

	fmt.Println(string([]rune{50504, 45397, 54616, 49464, 50836}))
	fmt.Println(string([]byte{236, 149, 136, 235, 133, 149, 237, 149, 152, 236, 132, 184, 236, 154, 148}))
	fmt.Println(string([]byte{236, 149, 136})) // 안
	fmt.Println(string([]byte{235, 133, 149})) // 녕
	fmt.Println(string([]byte{237, 149, 152})) // 하
	fmt.Println(string([]byte{236, 132, 184})) // 세
	fmt.Println(string([]byte{236, 154, 148})) // 요
}
