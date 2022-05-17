package main

import (
	"GO-LANG/5.package/lib"
	"fmt"
)

var v rune

func init() {
	v = '1'
}

func main() {
	fmt.Println(lib.IsDigit(v))
	fmt.Println("Hello")
	fmt.Println(lib.IsDigit('1'))
	fmt.Println(lib.IsDigit('a'))
	// 빌드 오류 // 식별자 소문자로 시작하므로 패키지 외부에서 접근 불가
	// fmt.Println(lib.isSpace('\t'))
}
