package main

import (
	"fmt"
)

type Item struct {
	name     string
	price    float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func main() {
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	fmt.Println(shirt)
	shoes := Item{name: "Sports Shoes", price: 30000}
	fmt.Println(shoes)
	dress := Item{name: "Stripe Shift Dress", quantity: 2}
	fmt.Println(dress)

	p := &Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	fmt.Println(p)

	/*
		구조체의 포인터를 생성하는 방법은 두 가지이다.
		1. new(Type)
		2. &Type{}

		두 구문 모두 주어진 타입에 맞는 메모리 공간을 할당해 제로값으로 초기화한 후 새로 만들어진 메모리의
		주소를 반환한다. 구조체 포인터를 &Type{}로 생성하면 초깃값이 할당된 구조체의 포인터를 생성할 수
		있어서 활용도가 높다.
	*/

	type rect struct{ w, h float64 }
	r1 := rect{1, 2}
	r2 := new(rect)
	r2.w, r2.h = 3, 4
	r3 := &rect{}
	r4 := &rect{5, 6}

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)

	// 익명 구조체
	// 구조체를 타입으로 정의하지 않고 익명으로 사용할 수도 있다.

	rects := []struct{ w, h int }{{1, 2}, {}, {-1, 2}}
	for _, r := range rects {
		fmt.Printf(" (%d, %d)\n", r.w, r.h)
	}

	// 내부 필드 접근
	var t Item
	t.name = "Men's Slim-Fit Shirt"
	t.price = 25000
	t.quantity = 3
	fmt.Println(t.name)
	fmt.Println(t.price)
	fmt.Println(t.quantity)
	fmt.Println(t.Cost())

}
