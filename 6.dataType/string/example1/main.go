package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(s[0])        // 104: 문자열 s의 첫 번째 문자 'h'의 코드값
	fmt.Println(s[len(s)-1]) // 111: 문자열 s의 마지막 문자 'o'의 코드값

	s2 := "안녕a세a"

	r1 := []rune(s)
	r2 := []rune(s2)
	r3 := []byte(s2)

	// 문자열은 UTF-8 인코딩을 사용한 유니코드 문자의 집합이다. 유니코드는 문자에 따라 바이트 수가
	// 다르므로 인덱스로 문자열 내부에 접근할 때는 주의해야 한다.

	fmt.Printf("s2: %c %c %c %c %c\n", s2[0], s2[1], s2[2], s2[3], s2[4])
	fmt.Printf("s1: %c %c %c %c %c\n", s[0], s[1], s[2], s[3], s[4])

	fmt.Printf("r1: %c %c %c %c %c\n", r1[0], r1[1], r1[2], r1[3], r1[4])
	fmt.Printf("r2: %c %c %c %c %c\n", r2[0], r2[1], r2[2], r2[3], r2[4])
	fmt.Printf("r3: %c %c %c %c %c\n", r3[0], r3[1], r3[2], r3[3], r3[4])

	for i, c := range s2 {
		fmt.Printf("%c(%d)\t", c, i)
	}
	fmt.Println()
	for i, c := range s {
		fmt.Printf("%c(%d)\t", c, i)
	}
	// 한글은 3byte이며 나머지는 1byte이므로 한글의 경우 인덱스가 3씩 커지는 것을 볼 수 있다.

	/*
		안(0) 녕(3) a(6) 세(7) 요(10)
		3byte(0 ~ 3) 3byte(3 ~ 6) 1byte(6 ~ 7) 3byte(7 ~ 10) 3byte
	*/

}
