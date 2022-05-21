package main

import (
	item "GO-LANG/8.oop/pkg"
	"fmt"
)

func main() {
	shirt := item.New("Mes's Slim-Fit Shirt", 25000, 3)
	fmt.Println(shirt)

	shirt.SetPrice(30000)
	shirt.SetQuantity(2)

	fmt.Println("Name: ", shirt.Name())
	fmt.Println("Price: ", shirt.Price())
	fmt.Println("Quantity: ", shirt.Quantity())
}
