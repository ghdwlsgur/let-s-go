Go는 변수를 선언할 때 초깃값을 지정하지 않으면 제로값(zero value)으로 초기화된다. 이는 변수에 값이 할당되기 전에 가비지(garbage) 값이 할당되어 예상치 못한 결과가 발생하는 것을 방지한다.

- 정수 타입의 제로값: 0
- 실수(부동소수점) 타입의 제로값: 0.0
- 문자열 타입의 제로값: ""

변수 이름에는 공백을 포함하지 않은 문자, 숫자, 밑줄(\_)을 사용할 수 있고, 숫자는 첫 글자에 사용할 수 없다. Go는 대소문자를 구분하므로 NUMBER, Number, number는 모두 다른 변수이다. 변수 이름에는 알파벳뿐만 아니라 유니코드 문자도 사용할 수 있다. 수학 기호를 사용하는 프로그램에서 변수 이름에 알파나 베타 또는 파이 같은 기호를 사용하면 프로그램을 좀 더 이해하기 쉽게 작성할 수 있다.

### Go 키워드

- break
- default
- func
- interface
- select
- case
- defer
- go
- map
- struct
- chan
- else
- goto
- package
- switch
- const
- fallthrough
- if
- range
- type
- continue
- for
- import
- return
- var

### Go 예약어

- append
- copy
- nil
- true
- bool
- delete
- panic
- byte
- error
- print
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- uintptr
- complex
- complex64
- complex128
- float32
- float64
- cap
- false
- println
- close
- iota
- real
- len
- recover
- imag
- make
- rune
- new
- string
