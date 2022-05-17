package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(intToUint8(100))
	test, _ := intToUint8(100)
	fmt.Println(test)
	fmt.Println(intToUint8(1000))
}

// intToUint8() 함수는 int 값을 받아 그 값이 uint8의 범위에 있으면 uint8로 변환한 값과
// nil을 반환하고, 범위에 없으면 0과 error를 반환한다.
func intToUint8(i int) (uint8, error) {
	if 0 <= i && i <= math.MaxUint8 {
		return uint8(i), nil
	}
	return 0, fmt.Errorf("%d cannot convert to uint8 type", i)
}
