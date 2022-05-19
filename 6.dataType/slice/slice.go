package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int
	b := []int{}
	c := []int{1, 2, 3}
	d := [][]int{{1, 2}, {3, 4, 5}}
	e := make([]int, 0)    // make 함수로 길이와 용량이 0인 슬라이스 생성
	f := make([]int, 3, 5) // make 함수로 길이가 3이고 용량이 5인 슬라이스 생성

	fmt.Printf("%-10T %d %d %v\n", a, len(a), cap(a), a)
	fmt.Printf("%-10T %d %d %v\n", b, len(b), cap(b), b)
	fmt.Printf("%-10T %d %d %v\n", c, len(c), cap(c), c)
	fmt.Printf("%-10T %d %d %v\n", d, len(d), cap(d), d)
	fmt.Printf("%-10T %d %d %v\n", e, len(e), cap(e), e)
	fmt.Printf("%-10T %d %d %v\n", f, len(f), cap(f), f)

	numbers := []int{3, 4, 5, 6, 7, 8, 4, 6, 8, 32, 4}
	// 인덱스 사용
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// 인덱스 미사용
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println("sum: ", sum)

	// 내부 요소 변경
	total := 0
	for i := range numbers {
		numbers[i] *= 2
		total += numbers[i]
	}
	fmt.Println("numbers: ", numbers)
	fmt.Println("total: ", total)

	// 부분 슬라이스 추출
	fmt.Println(numbers, "=", numbers[:3], numbers[3:5], numbers[5:])

	// 슬라이스 추가
	ns1 := []int{1, 2, 3}
	ns2 := []int{6, 7, 8}
	ns3 := []int{8, 9, 10, 11}

	// 슬라이스의 각 요소를 개별로 추가할 때는 ... 연산자를 사용한다.
	ns1 = append(ns1, 4, 5)
	fmt.Println(ns1)
	ns1 = append(ns1, ns2...)
	fmt.Println(ns1)
	ns1 = append(ns1, ns3[1:3]...)
	fmt.Println(ns1)

	s := make([]int, 0, 3)
	for i := 0; i < 9; i++ {
		s = append(s, i)
		fmt.Printf("len: %d, cap: %d, %v\n", len(s), cap(s), s)
	}

	// 슬라이스 삽입
	k := []int{1, 2, 3, 4, 5}
	// k = insert2(k, []int{-3, -2}, 0)
	// fmt.Println(k)

	k = insert2(k, []int{7}, 2)
	// fmt.Println(k)

	// k = insert(k, []int{6, 7}, len(s))
	// fmt.Println(k)

	// 슬라이스 정렬
	/*
		Todo: Float64
		- sort.Float64s(a)
		[]float64 슬라이스를 오름차순으로 정렬

		- sort.Float64sAreSorted(a)
		[]float64 슬라이스 정렬되어 있는지 확인

		Todo: Int
		- sort.Ints(a)
		[]int 슬라이스를 오름차순으로 정렬

		- sort.IntsAreSorted(a)
		[]int 슬라이스가 정렬되어 있는지 확인

		Todo: String
		- sort.Strings(a)
		[]string 슬라이스를 오름차순으로 정렬

		- sort.StringsAreSorted(a)
		[]string 슬라이스가 정렬되어 있는지 확인

		Todo: Interface
		- sort.Sort(data)
		sort.Interface 타입 슬라이스를 오름차순으로 정렬

		- sort.IsSorted(data)
		sort.Interface 타입 슬라이스가 정렬되어 있는지 확인
	*/

	kk := []int{5, 2, 6, 3, 1, 4}
	fmt.Println(kk)
	fmt.Println(sort.IntsAreSorted(kk))
	sort.Ints(kk)
	fmt.Println(kk)
	fmt.Println(sort.IntsAreSorted(kk))

}

// 슬라이스 삽입
func insert(s, new []int, index int) []int {
	return append(s[:index], append(new, s[index:]...)...)
}

// append() 함수 사용하지 않고 insert 기능 구현
func insert2(s, new []int, index int) []int {
	result := make([]int, len(s)+len(new))
	fmt.Println(result)
	position := copy(result, s[:index])
	fmt.Println(result)
	fmt.Println("position1: ", position)
	position += copy(result[position:], new)
	fmt.Println(result)
	fmt.Println("position2: ", position)
	copy(result[position:], s[index:])
	fmt.Println(result)
	return result
}
