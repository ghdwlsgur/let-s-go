package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	i  int64
	mu sync.Mutex
}

func (c *counter) increment() {
	c.mu.Lock()   // i 값을 변경하는 부분(임계 영역)을 뮤텍스로 잠금
	c.i += 1      // 공유 데이터 변경
	c.mu.Unlock() // i 값을 변경 완료한 후 뮤텍스 잠금 해제
}

func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	// 모든 CPU 점유
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}          // 카운터 생성
	done := make(chan struct{}) // 완료 신호 수신용 채널

	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()
			done <- struct{}{}
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}

	c.display()

}
