package main

import "fmt"

type rect struct{ width, height float64 }
type circle struct{ radius float64 }

func display(s interface{ show() }) {
	s.show()
}

func (r rect) show() {
	fmt.Printf(" width: %f, height: %f\n", r.width, r.height)
}

func (c circle) show() {
	fmt.Printf(" radius: %f\n", c.radius)
}

func main() {
	r := rect{3, 4}
	c := circle{2.5}
	display(r)
	display(c)
}
