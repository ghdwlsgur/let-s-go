```Go
func main() {
  // rd와 wr의 버퍼 사이즈
  // 1<<24 => 2**24 => 16,777,216

	rd := bufio.NewReaderSize(os.Stdin, 1<<24)
	wr := bufio.NewWriterSize(os.Stdout, 1<<24)

  // defer 구문에 오는 함수는 가장 마지막에 실행 (try catch finally)
	defer wr.Flush()

  // 반복문을 몇 번 돌 것인지 사용자로부터 입력받는다.
	raw, _, _ := rd.ReadLine()
  // raw를 숫자형으로 반환한다.
	t, _ := strconv.Atoi(string(raw))


	for tc := 1; tc <= t; tc++ {
		raw, _, _ = rd.ReadLine()

    // 한 줄에 사용자로부터 입력받은 raw를 아스키코드값으로 반환한다.
    /*
      1 1을 입력받았을 경우 [[49] [49]] 출력
    */
		line := bytes.Fields(raw)

    // 전달받은 문자열을 숫자로 반환 , 두번째 err은 _ 처리
		a, _ := strconv.Atoi(string(line[0])) // [49]
		b, _ := strconv.Atoi(string(line[1])) // [49]

    // Itoa 숫자형 자료형을 문자열 타입으로 바꿔준다.
		wr.WriteString(strconv.Itoa(a + b))
		wr.WriteByte('\n')
	}
```
