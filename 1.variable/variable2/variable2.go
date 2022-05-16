package main

import "fmt"

func main() {
	// 같은 타입의 변수 여러 개를 한 번에 선언
	var name, id, address string

	// 제로값 확인 (공백 출력)
	fmt.Println(name, id, address)

	name = "홍진혁"
	id = "1"
	address = "Seoul"

	fmt.Println(name, id, address)

	// 다른 타입의 변수 여러 개를 한 번에 선언
	var (
		consumer string
		age      int
		weight   float32
	)

	consumer = "홍진호"
	age = 19
	weight = 70.5

	fmt.Println(consumer, age, weight)

	// 변수 타입 생략
	// 변수 선언과 동시에 값을 할당할 때는 타입 생략 가능
	var c = true

	fmt.Println(c)

	// 짧은 선언
	/*
		변수 선언과 동시에 값을 할당할 때는 var 생략 가능 var를 생략할 때는 := 연산자 사용
		그러나 이미 선언된 변수에 값을 할당할 때는 := 연산자를 사용할 수 없다.
		또한 짧은 선언은 함수 안에서만 가능하다.

		Go는 제한된 범위 내에서만 사용하는 변수를 선언할 때 변수 타입을 생략하고 := 연산자로
		짧게 선언하는 패턴을 자주 사용한다. 짧은 선언은 변수의 선언 방식이나 타입 보다는 코드의
		패턴과 흐름에 집중하여 코드를 작성할 수 있게 해준다.
	*/

	start := 1

	fmt.Println(start)

}
