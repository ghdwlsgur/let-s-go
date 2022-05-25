package main

import "fmt"

func main() {
	var h, m int
	fmt.Scanf("%d %d", &h, &m)
	getMinutes := h*60 + m - 45
	hh := getMinutes / 60
	mm := getMinutes - (hh * 60)
	if getMinutes < 0 {
		fmt.Printf("%d %d", 23, mm+60)
		return
	}
	fmt.Printf("%d %d", hh, mm)
}
