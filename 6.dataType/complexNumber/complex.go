package main

import "fmt"

/*
 복소수는 내장 함수 complex()로 생성하거나 리터럴 표기법 (실수부 + 허수부)으로 직접 변수에 할당할 수 있다.
 복소수를 리터럴로 표현할 때는 허수부에 접미사 i를 붙인다. 실수부에 해당하는 값은 real() 함수로 얻어
 올 수 있고, 허수부에 해당하는 값은 imag() 함수로 얻어올 수 있다.
*/
func main() {
	c1 := 1 + 2i
	c2 := complex64(3 + 4i)
	c3 := complex(5, 6)

	fmt.Println(c1, real(c1), imag(c1))
	fmt.Println(c2, real(c2), imag(c2))
	fmt.Println(c3, real(c3), imag(c3))

}
