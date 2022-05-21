package main

import "fmt"

type Coster interface {
	Cost() float64
}

func displayCost(c Coster) {
	fmt.Println(" cost: ", c.Cost())
}

type Item struct {
	name     string
	price    float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price
}

type DiscountItem struct {
	Item
	discountRate float64
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

/*
	Item 타입과 DiscountItem 타입은 Cost() float64 메서드를 가지므로 Coster 인터페이스로
	사용할 수 있다. Item과 DiscountItem 둘 다 displayCost() 함수를 통해 Cost()를 출력한다.
*/

func main() {
	shoes := Item{"Sports Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 3},
		10.00,
	}
	displayCost(shoes)
	displayCost(eventShoes)
}
