package main

import "fmt"

// 메서드 오버라이딩 (Overriding)

type Item struct {
	name     string
	price    float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

type DiscountItem struct {
	Item
	discountRate float64
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

func main() {
	shoes := Item{"Women's Walking Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Sports Shoes", 50000, 3},
		10.00,
	}

	fmt.Println(shoes.Cost())
	fmt.Println(eventShoes.Cost())
	fmt.Println(eventShoes.Item.Cost())
}
