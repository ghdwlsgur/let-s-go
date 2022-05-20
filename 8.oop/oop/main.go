package main

import "fmt"

/*
	Go에서는 기본 타입 외에 사용자가 타입을 직접 정의할 수 있는 사용자 정의 타입이 있다.
	일반적으로는 구조체와 인터페이스를 사용자 정의 타입으로 지정해서 쓴다. 기본 타입이나 함수 서명을
	사용자 정의 타입으로 지정해서 쓰기도 한다.

	메서드는 사용자 정의 타입과 함수를 바인딩시키는 방식으로 정의한다. 메서드를 정의할 때는
	어떤 타입과 연결할지 리시버를 지정해준다. 리시버를 지정해주면 타입과 메서드가 연결되어 특정 사용자 정의
	타입에 대한 동작을 표현할 수 있게 된다.

	func (리시버명 리시버타입) 메서드명 (매개변수) (반환타입 또는 반환 값) {

	}

	func과 메서드명 사이에 리시버를 지정해주는 것만 제외하면 함수를 정의하는 방식과 같다.
	메서드를 호출하면 리시버 값이 메서드 내부로 전달되며, 메서드 안에서 리시버 값을 사용할 수 있게 된다.
*/

// ! 타입 정의
type Item struct {
	name     string
	price    float64
	quantity int
}

// ! 메서드 정의
func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

// ! 타입 정의
type rect struct {
	width  float64
	height float64
}

// ! 메서드 정의
func (r rect) area() float64 {
	return r.width * r.height
}

func main() {
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	fmt.Println(shirt.Cost())

	type quantity int
	type dozen []quantity

	var d dozen
	for i := quantity(1); i <= 12; i++ {
		d = append(d, i)
	}
	fmt.Println(d)

	r := rect{3, 4}
	fmt.Println("area: ", r.area())

}
