package main

import "fmt"

/*
	defer 키워드는 함수가 종료되기 전까지 특정 구문의 실행을 지연시켰다가, 함수가 종료되기 직전에
	지연시켰던 구문을 수행한다. 자바나 c#의 finally 같은 개념이다. 주로 리소스를 해제시키거나
	클렌징 작업이 필요할 때 사용한다.
*/

func main() {

	f1()
	f()

}

func f1() {
	fmt.Println("f1 - start")
	defer f2()
	fmt.Println("f1 - end")
}

func f2() {
	fmt.Printf("f2 - deferred")
}

// defer를 여러 개 사용
/*
	함수 하나에 defer 키워드를 여러 개 사용하면 defer로 지정한 각 구문은 스택(stack)에 쌓였다가
	가장 나중에 쌓인 defer 구문부터 수행된다.
*/
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}

/* defer 사용예제
- 파일 스트림 닫기
file, _ := os.Open(path)
defer file.Close()

- 리소스 잠금 해제하기
mu.Lock()
defer mu.Unlock()

- 레포트에서 푸터 출력하기
printHeader()
defer printFooter()

- 데이터베이스 커넥션 닫기
con, _ := Connection()
defer con.Close()
*/
