package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type counter struct {
	i int64
}

// counter의 값을 1씩 증가
func (c *counter) increment() {
	atomic.AddInt64(&c.i, 1)
}

// counter의 값을 출력
func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}     // 카운터 생성
	wg := sync.WaitGroup{} // WaitGroup 생성

	for i := 0; i < 1000; i++ {
		wg.Add(1) // WaitGroup의 고루틴 개수 1 증가
		go func() {
			defer wg.Done() // 고루틴 종료 시 Done() 처리
			c.increment()   // 카운터 값을 1 증가시킴
		}()
	}

	wg.Wait()   // 모든 고루틴이 종료될 떄까지 대기
	c.display() // c의 값 출력
}
