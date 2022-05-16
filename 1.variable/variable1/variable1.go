package main

// 문자열 출력
import "fmt"

func main() {
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)
}

/* 모듈 설치
go mod init GO-LANG/variable
*/

/* 빌드 진행
go build
*/

/* 파일 실행
./variable
*/
