### sync.RWMutex

> sync.RWMutex는 sync.Mutex와 동작 방식이 유사하다. 하지만 sync.RWMutex는 읽기 동작과 쓰기 동작을 나누어 잠금 처리할 수 있다.

- #### 읽기 잠금
  읽기 잠금은 읽기 작업에 한해서 공유 데이터가 변하지 않음을 보장해준다. 즉, 읽기 잠금으로 잠금 처리해도 다른 고루틴에서 잠금 처리한 데이터를 읽을 수는 있지만, 변경할 수는 없다.
- #### 쓰기 잠금

  쓰기 잠금은 공유 데이터에 쓰기 작업을 보장하기 위한 것으로, 쓰기 잠금으로 잠금 처리하면 다른 고루틴에서 읽기와 쓰기 작업 모두 할 수 없다.

- `func (rw *RWMutex) Lock()` `쓰기 잠금`
- `func (rw *RWMutex) UnLock()` `쓰기 잠금 해제`
- `func(rw *RWMutex) RLock()` `읽기 잠금`
- `func(rw *RWMutex) RUnlock()` `읽기 잠금 해제`
