package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	sync.WaitGroup은 모든 고루틴이 죵료될 때까지 대기해야 할 때 사용한다.

	func (wg *WaitGroup) Add(delta int): WaitGroup에 대기 중인 고루틴 개수 추가
	func (wg *WaitGroup) Done(): 대기 중인 고루틴의 수행이 종료되는 것을 알려줌
	func (wg *WaitGroup) Wait(): 모든 고루틴이 종료될 때까지 대기
*/

type counter struct {
	i  int64
	mu sync.Mutex // 공유 데이터 i를 보호하기 위한 뮤텍스
}

func (c *counter) increment() {
	c.mu.Lock()
	c.i += 1
	c.mu.Unlock()
}

// counter의 값을 출력
func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	// 모든 CPU를 사용하게 함
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}     // 카운터 생성
	wg := sync.WaitGroup{} // WaitGroup 생성

	// c.increment()를 실행하는 고루틴 1000개 실행
	for i := 0; i < 1000; i++ {
		wg.Add(1) // WaitGroup의 고루틴 개수 1 증가
		go func() {
			defer wg.Done() // 고루틴 종료 시 Done() 처리
			c.increment()   // 카운터 값을 1 증가시킴
		}()
	}

	wg.Wait() // 모든 고루틴이 종료될 때까지 대기

	c.display() // c의 값 출력

	// ! 주의할 점은 Add 메서드로 추가한 고루틴의 개수와 Done 메서드를 호출한 횟수는 같아야 한다.

}
