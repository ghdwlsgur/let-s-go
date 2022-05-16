package main

import (
	"fmt"
	"strconv"
)

func area(w, h int) int {
	return w * h
}

func multiply(w, h int) (int, int) {
	return w * 2, h * 2
}

// ! 함수에서 값을 너무 많이 반환하면 가독성이 떨어질 수 있다.
// ! 반환 값은 보통 세 개 이내로만 사용하는 게 좋다.
// 값을 네 개 이상 반환할 때는 배열이나 구조체로 변환하는 것을 권장한다.

func main() {
	v := area(3, 4)
	fmt.Println(v)

	w, h := multiply(3, 4)
	fmt.Println(w, h)

	displayInt("two")
	displayInt("2")

	/*
		Go에는 값을 두 개 이상 반환하는 함수가 많습니다 (수행결과와 에러상태를 반환)
		정상적으로 함수가 수행된 후에는 결과값과 널(null)을 의미하는 nil을 반환하고, 정상적으로 수행되지
		않았을 때는 제로값 (false, 0, "", nil)과 에러 상태(error)를 반환한다.
		Go의 기본 라이브러리 패키지에 있는 함수들이 대부분 이런 방식으로 작성되어 있다.
	*/
	test, err := strconv.Atoi("2")
	fmt.Println("test :", test, "err :", err)
}

/*
	strconv 패키지의 Atoi 함수는 문자열을 정수로 변환하여 수행 결과와 에러 상태를 반환한다.
*/
func displayInt(s string) {
	if v, err := strconv.Atoi(s); err != nil {
		fmt.Printf("%s는 정수가 아닙니다.\n", s)
	} else {
		fmt.Printf("정수 값은 %d입니다.\n", v)
	}

	// i, _ := myFunc()
	_, i := myFunc()
	fmt.Println(i)
}

/*
 값을 두 개 이상 반환하는 함수를 사용할 때 반환 값 중 하나만 필요할 때가 있다.
 반환 값 중 하나만 필요할 떄는 사용하지 않을 값을 임시로 받아줄 공간이 필요하다.
 이때 빈 식별자(_)를 사용한다.
*/

func myFunc() (num int, content string) {
	return 2, "hi"
}
