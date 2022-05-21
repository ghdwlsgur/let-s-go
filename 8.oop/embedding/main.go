package main

import "fmt"

/*
	보편적으로 상속을 통해 코드를 재사용하고, 객체 지향 언어 대부분은 상속을 지원한다.
	하지만 객체 지향 프로그래밍이 등장한 지 수십 년이 지난 지금, 상속의 단점이 많이 드러났다.
	한 클래스가 다른 클래스를 상속하여 클래스를 만들면 클래스 간에 부모-자식 관계가 생긴다.
	프로그램 구조가 커져 상속 관계가 점점 깊어지면 클래스들이 거대한 트리 구조가 되고 이는 여러
	문제를 유발한다. 그래서 디자인 패턴에서는 상속보다는 조합을 강조한다.

	Go는 언어 차원에서 상속을 없애고 사용자 정의 타입을 조합하여 구조체를 정의하는 방식으로
	겍체를 재사용한다. 이처럼 사용자 정의 타입을 구조체의 필드로 지정하는 것을 임베딩이라 한다.
*/

type Option struct {
	name  string
	value string
}

type Item struct {
	name     string
	price    float64
	quantity int
	Option   // 임베디드 필드
}

func main() {
	shoes := Item{"Sports Shoes", 30000, 2, Option{"color", "red"}}

	fmt.Println(shoes)

	fmt.Println(shoes.value)

	fmt.Println(shoes.name)
	fmt.Println(shoes.Option.name)
}
