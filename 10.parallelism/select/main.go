package main

import "fmt"

/*
	select문은 하나의 고루틴이 여러 채널과 통신할 때 사용한다. case로 여러 채널을 대기시키고 있다가
	실행 가능 상태가 된 채널이 있으면 해당 케이스를 수행한다.
*/

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// x 값을 c 채널로 송신
		case c <- x:
			x, y = y, x+y
			/*
				0, 1
				1, 1
				1, 2
				2, 3
				3, 5
				5, 8
				8, 13
				13, 21
				21, 34
				34, 55
			*/
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)    // int 타입의 데이터를 송수신할 수 있는 c 채널 생성
	quit := make(chan int) // int 타입의 데이터를 송수신할 수 있는 quit 채널 생성

	// 하나의 고루틴에서 두 개의 채널과 송수신
	// 두 개의 채널은 fibonacci 함수에서 select문 안에 로직으로 구성
	go func() {
		// 0부터 9까지, 10번 채널 c로부터 수신받은 데이터를 출력한다.
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		// 반복문이 종료되었을 때 0을 quit으로 전송하며
		// quit 채널에 데이터를 수신할 경우 반복문이 종료되도록 로직 작성
		quit <- 0 // 0을 quit 채널로 전송
	}()

	fibonacci(c, quit)
}
