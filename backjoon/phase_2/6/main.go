package main

import "fmt"

func main() {
	var h, m, add int
	fmt.Scanf("%d %d\n%d", &h, &m, &add)
	getMinutes := h*60 + m + add
	hh := getMinutes / 60
	mm := getMinutes - (hh * 60)
	if hh >= 24 {
		fmt.Printf("%d %d", hh-24, mm)
		return
	}
	fmt.Printf("%d %d", hh, mm)
}
