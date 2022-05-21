package main

import "fmt"

type rect struct{ width, height float64 }

func (r rect) area() float64 {
	return r.width * r.height
}

func (r *rect) resize(w, h float64) {
	r.width += w
	r.height += h
}

func main() {
	r := rect{3, 4}
	fmt.Println("area: ", r.area())
	r.resize(10, 10)
	fmt.Println("area: ", r.area())

	// 함수 표현식
	areaFn := rect.area        // func(rect) float64
	resizeFn := (*rect).resize // func(*rect, float64, float64)

	fmt.Println("area: ", areaFn(r))
	resizeFn(&r, -10, -10)
	fmt.Println("area: ", areaFn(r))

}
