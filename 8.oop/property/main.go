package main

import "fmt"

type dimension struct {
	width, height, length float64
}

type Item struct {
	name             string
	price            float64
	quantity         int
	packageDimension dimension
	itemDimension    dimension
}

func main() {
	shoes := Item{
		"Sports Shoes", 30000, 2,
		dimension{30, 270, 20},
		dimension{50, 300, 30},
	}

	// 구조체 값을 출력할 때 필드명과 값을 함께 출력하려면 %#v를 사용하면 된다.
	fmt.Printf("%#v\n", shoes.itemDimension)
	fmt.Printf("%#v\n", shoes.packageDimension)

	fmt.Println(shoes.packageDimension.width)
	fmt.Println(shoes.packageDimension.height)
	fmt.Println(shoes.packageDimension.length)
}
