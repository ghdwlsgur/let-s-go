package main

import "fmt"

// ! 원시 타입
type quantity int

func (q quantity) greaterThan(i int) bool {
	return int(q) > i
}

func (q *quantity) increment() {
	*q++
}

func (q *quantity) decrement() {
	*q--
}

// ! 참조 타입
type numberMap map[string]int

func (m numberMap) add(key string, value int) {
	m[key] = value
}

func (m numberMap) remove(key string) {
	delete(m, key)
}

func main() {
	q := quantity(3)
	q.increment()
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))
	q.decrement()
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))

	m := make(numberMap)
	m["one"] = 1
	m["two"] = 2
	m.add("three", 3)
	fmt.Println(m)
	m.remove("two")
	fmt.Println(m)
}
