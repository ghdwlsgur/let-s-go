package main

import "fmt"

func main() {
	/*
		Go에서는 모든 반복문을 for문으로 작성한다.

		for 초기화구문; 조건식; 후속작업 {

		}
	*/

	/*
		for 문에 레이블로 식별자를 붙일 수 있다.
		for문 앞에 콜론(:)으로 끝나는 문자가 있으면 레이블로 인식한다.
		continue, break 레이블을 함께 사용하면 반복문을 유연하게 제어할 수 있다.
	*/

	x := 7
	table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}

LOOP:
	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[row]); col++ {
			if table[row][col] == x {
				fmt.Printf("found %d(row: %d, col: %d)\n", x, row, col)
				break LOOP
			}
		}
	}

	y := 5
	t := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}

LOOP2:
	for row := 0; row < len(t); row++ {
		for col := 0; col < len(t); col++ {
			if t[row][col] == y {
				fmt.Printf("found %d(row: %d, col: %d\n", y, row, col)
				continue LOOP2
			}
		}
	}

}
