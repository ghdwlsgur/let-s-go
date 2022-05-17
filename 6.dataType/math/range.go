package main

import "fmt"

func main() {
	const (
		MaxInt8   = 1<<7 - 1
		MinInt8   = -1 << 7
		MaxInt16  = 1<<15 - 1
		MinInt16  = -1 << 15
		MaxInt32  = 1<<31 - 1
		MinInt32  = -1 << 31
		MaxInt64  = 1<<64 - 1
		MinInt64  = -1 << 63
		MaxUint8  = 1<<8 - 1
		MaxUint16 = 1<<16 - 1
		MaxUint32 = 1<<32 - 1
		MaxUint64 = 1<<64 - 1
	)

	fmt.Printf("MaxInt8: %d, MinInt8: %d\n", MaxInt8, MinInt8)
	fmt.Printf("MaxInt16: %d, MinInt16: %d\n", MaxInt16, MinInt16)
	fmt.Printf("MaxInt32: %d, MinInt32: %d\n", MaxInt32, MinInt32)
	// fmt.Printf("MaxInt64: %d, MinInt64: %d\n", MaxInt64, MinInt64)

	/*
		증감 연산자(++, --)는 후치 연산으로만 사용할 수 있고 반환 값은 없다. 이는 증감 연산자의
		과도한 사용으로 코드의 가독성이 떨어지는 것을 막기 위해서다. 표현식에 증감 연산자를 사용하거나
		증감 연산자로 전치 연산을 하면 코드가 이해하기 어려워진다.

		정수 타입에만 사용할 수 있는 연산자
		- ^X
		- x %= y
		- x &= y
		- x |= y
		- x ^= y
		- x &^= y
		- x >>= u
		- x <<= u
		- x % y
		- x & y
		- x | y
		- x ^ y
		- x &^ y
		- x << u
		- x >> u
	*/
}
