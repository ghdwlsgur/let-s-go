package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main 함수 시작", time.Now())

	go long()
	go short()

	time.Sleep(5 * time.Second)
	/*
		고루틴을 사용할 때 주의할 점이 있다. 실행 중인 고루틴이 있어도 메인 함수가 종료되면 프로그램이
		종료된다. 그래서 아직 실행 중인 고루틴이 있다면 메인 함수각 종료되지 않게 해야 한다.
		만약 메인 함수의 대기 시간을 2초로 변경한다면 프로그램은 long 함수가 종료되기 전에 프로그램이
		종료되고 만다. 프로그램이 비정상적으로 종료되는 것을 피하기 위해서는 메인 함수에 충분히 긴 시간
		동안 대기하게 할 수도 있지만 각 고루틴이 얼마나 오랫동안 실행될지 알 수 없어서 이 방법은
		효율적이지 않다. Go에서 제안하는 방법은 메인 함수에서 고루틴의 종료 상황을 확인할 수 있는
		채널을 만들고, 다른 채널을 통해 종료 메시지를 대기시키는 것이다.
	*/
	fmt.Println("main 함수 종료", time.Now())
}

func long() {
	fmt.Println("long 함수 시작", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("long 함수 종료", time.Now())
}

func short() {
	fmt.Println("short 함수 시작", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("short 함수 종료", time.Now())
}
