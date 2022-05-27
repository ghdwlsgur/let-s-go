package main

import (
	"fmt"
	"time"
)

// 채널은 고루틴끼리 정보를 교환하고 실행의 흐름을 동긱화하기 위해 사용한다.

func main() {
	fmt.Println("main 함쑤 시작", time.Now())

	done := make(chan bool)
	go long(done)
	go short(done)

	<-done
	<-done
	fmt.Println("main 함수 종료", time.Now())
}

func long(done chan bool) {
	fmt.Println("long 함수 시작", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("long 함수 종료", time.Now())
	done <- true
}

func short(done chan bool) {
	fmt.Println("short 함수 시작", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("short 함수 종료", time.Now())
	done <- true
}
