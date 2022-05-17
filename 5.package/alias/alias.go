package main

// 패키지 이름에 빈 식별자로 별칭을 주면 컴파일 에러가 발생하지 않는다. 디버깅시 유용하게 사용
import (
	// mylib "GO-LANG/5.package/lib"
	_ "GO-LANG/5.package/lib"
	"fmt"
)

func main() {
	fmt.Println("lib 패키지 사용 코드 임시 제거")
	// fmt.Println(mylib.IsDigit('1'))
	// fmt.Println(mylib.IsDigit('a'))
}
