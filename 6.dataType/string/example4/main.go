package main

/*
	문자열은 한 번 생성되면 그 값을 변경할 수 없으므로 +로 문자열을 조합하면 조합할 때마다 항상
	새로운 문자열을 생성한다. 그래서 문자열을 +로 조합하는 것은 효율적이지 않다. 문자열을 조합할 때
	+ 대신 strings.Join() 함수를 사용하거나 bytes.Buffer 타입을 사용하면 더 빠르고 효율적이다.
*/

// ! 문자열은 한 번 생성되면 그 값을 변경할 수 없으므로 + 로 문자열을 조합하면 조합할 때마다 항상 새로운
// ! 문자열을 생성한다.

import (
	"fmt"
	"math"
	"unicode"
)

// + 연산자 사용
func main() {
	var str string
	for i := 0; i < math.MaxUint8; i++ {
		if s, ok := nextString(i); ok {
			str += s
		}
	}
	fmt.Print(str, "\n")
}

func nextString(i int) (s string, ok bool) {
	if unicode.IsLetter(rune(i)) {
		return string(i), true
	}
	return "", false
}
