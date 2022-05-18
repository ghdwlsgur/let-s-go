### 정수 타입과 범위

|  타입   |                            크기                             |                           범위                           |
| :-----: | :---------------------------------------------------------: | :------------------------------------------------------: |
|  int8   |                           1바이트                           |                      `-128` ~ `127`                      |
|  int16  |                           2바이트                           |                    `-32,768`~`32,767`                    |
|  int32  |                           4바이트                           |             `-2,147,483,648`~`2,147,483,647`             |
|  int64  |                           8바이트                           | `-9,223,372,036,854,775,808`~`9,223,372,036,854,775,807` |
|   int   |  32비트 머신에서는 `int32`</br> 64비트 머신에서는 `int64`   |                                                          |
|  uint8  |                           1바이트                           |                        `0`~`255`                         |
| uint16  |                           2바이트                           |                       `0`~`65,535`                       |
| uint32  |                           4바이트                           |                   `0`~`4,294,967,295`                    |
| uint64  |                           8바이트                           |             `0`~`18,446,744,073,709,551,615`             |
|  uint   | 32비트 머신에서는 `uint32` </br> 64비트 머신에서는 `uint64` |                                                          |
|  byte   |                        uint8의 별칭                         |                                                          |
|  rune   |                        int32의 별칭                         |                                                          |
| uintptr |                         uint와 같음                         |                 포인터를 저장할 때 사용                  |

### 실수 타입과 범위

|  타입   |            범위            |  전체 자릿수  |
| :-----: | :------------------------: | :-----------: |
| float32 |   ±10^-453~±3.4 \* 10^38   | 소수점 7자리  |
| float64 | ±5 _ 10^-324~±1.7 _ 10^308 | 소수점 15자리 |

### 복소수 타입과 범위

|    타입    |           범위           |
| :--------: | :----------------------: |
| complex64  | 32비트의 실수부와 허수부 |
| complex128 | 64비트의 실수부와 허수부 |

### 문자나 문자열에서 사용하는 이스케이프

| 이스케이프 |                       설명                       |
| :--------: | :----------------------------------------------: |
|     \\     |                   역슬래기(\)                    |
|     \'     |    작은 따옴표('), 문자 리터럴에만 사용 가능     |
|     \"     |     큰따옴표("), 문자열 리터럴에만 사용 가능     |
|     \a     |      벨(ASCII 코드7, BEL), 콘솔의 벨을 울림      |
|     \b     |                    백스페이스                    |
|     \f     |                     쪽 바꿈                      |
|     \n     |                     줄 바꿈                      |
|     \r     |                       복귀                       |
|     \t     |                        탭                        |
|     \v     |                     수직 탭                      |
|    \xhh    | 유니코드 문자, 16진수 코드 포인트를 2자리로 표현 |
|    \ooo    | 유니코드 문자, 8진수 코드 포인트를 3자리로 표현  |
|   \uhhhh   | 유니코드 문자, 16진수 코드 포인트를 4자리로 표현 |
| \Uhhhhhhhh | 유니코드 문자, 16진수 코드 포인트를 8자리로 표현 |

#### 자주 사용하는 문자열 관련 함수

|           함수            |                설명                |
| :-----------------------: | :--------------------------------: |
|          len(s)           |        문자열 s의 바이트 수        |
|      len([]rune(s))       |         문자열 s의 문자 수         |
| utf8.RuneCountInString(s) |         문자열 s의 문자 수         |
|      strconv.Atoi(s)      | 문자열 s를 정수(int) 타입으로 변환 |
|      strconv.Itoa(i)      |       정수 i를 문자열로 변환       |