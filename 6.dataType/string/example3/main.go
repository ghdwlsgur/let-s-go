package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "hello" // 5byte
	s2 := "안녕하세요" // 15byte

	fmt.Println(len(s1))                    // 5
	fmt.Println(len(s2))                    // 15
	fmt.Println(utf8.RuneCountInString(s2)) // 5
	fmt.Println(len([]rune(s2)))            // 5

	fmt.Println(s1[1:2])
	fmt.Println(s1[1:])
	fmt.Println(s1[:2])

	fmt.Println(s2[:3])
	fmt.Println(s2[3:6])
	fmt.Println(s2[6:9])
	fmt.Println(s2[9:12])
	fmt.Println(s2[12:])

	text := "Go is an open source programming language" + " that makes it easy to build simple, reliable, and efficient software"
	text += " GO is expressive, concise, clean and efficient."
	fmt.Println(text)
}
