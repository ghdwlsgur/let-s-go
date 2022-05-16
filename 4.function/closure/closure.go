package main

import "fmt"

/*
	때로는 함수 이름을 저장하지 않고 익명 함수 형태로 사용할 때가 있다. Go에서 함수는
	일급 객체(first-class object)이므로 변수의 값으로 사용할 수 있다. 다음과 같이 함수를
	변수에 할당하여 변수처럼 사용할 수 있다.
*/

func main() {
	fplus := func(x, y int) int {
		return x + y
	}

	fmt.Println(fplus(3, 4))

	func() {
		fmt.Println("익명함수 호출")
	}()

	func(x, y int) int {
		fmt.Println(x, y)
		return x + y
	}(3, 4)

}
