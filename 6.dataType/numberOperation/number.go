package main

import "fmt"

/*
숫자 연산(산술 연산, 비교 연산)은 타입이 같은 숫자끼리만 할 수 있다. 타입이 다른 숫자끼리
연산하려면 반드시 타입을 변환해주어야 한다. 타입 변환 코드는 다음과 같다. 타입명(값)

타입이 다른 숫자끼리 연산하려고 하면 컴파일 에러가 발생한다.
*/

func main() {
	i := 100000
	j := int16(10000)
	// k := uint8(100)

	// compile Error
	// fmt.Println(i + j)
	// fmt.Println(i + k)
	// fmt.Println(j > k)

	fmt.Println(i + int(j))
	fmt.Println(j + int16(i))
}
