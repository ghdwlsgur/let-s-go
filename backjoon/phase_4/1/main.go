package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var getInt, n int
	var max, min = -1000000, 1000000
	var minus bool = false
	fmt.Scanln(&n)

	var reader *bufio.Reader = bufio.NewReaderSize(os.Stdin, n*1)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input, _ := reader.ReadBytes('\n')

	for _, num := range input {
		switch {
		// 유니코드 48 == 0, 유니코드 57 == 9
		case num >= 48 && num <= 57:
			// int(num - 48) byte에서 int자료형으로 형변환
			getInt = getInt*10 + int(num-48)
		// '-' = 45 마이너스일 경우
		case num == 45:
			minus = true
		default:
			// 비교숫자에 -1 곱셈
			if minus {
				getInt *= -1
			}
			// 최대값
			if max < getInt {
				max = getInt
			}
			// 최소값
			if min > getInt {
				min = getInt
			}
			// 초기화
			minus = false
			getInt = 0
		}
	}
	writer.WriteString(strconv.Itoa(min))
	writer.WriteByte(' ')
	writer.WriteString(strconv.Itoa(max))
}

// func test() {
// 	var a, n int
// 	var max, min = -1000000, 1000000
// 	fmt.Scanln(&n)

// 	var reader *bufio.Reader = bufio.NewReaderSize(os.Stdin, n*1)
// 	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
// 	defer writer.Flush()

// 	input, _ := reader.ReadBytes('\n')

// fmt.Println(input)
// fmt.Printf("%c\n", input) // 유니코드 문자 반환

// 	for _, c := range input {

// 		switch {
// 		case c >= 48 && c <= 57:
// 			a = a*10 + int(c-48)
// 		case c == 45:
// 			a *= -1
// 		default:
// 			if max < a {
// 				max = a
// 			}
// 			if min > a {
// 				min = a
// 			}
// 			if a < 0 {
// 				a *= -1
// 			}
// 			a = 0
// 		}
// 	}
// 	writer.WriteString(strconv.Itoa(min))
// 	writer.WriteByte(' ')
// 	writer.WriteString(strconv.Itoa(max))

// }

// func main() {
// 	rd := bufio.NewReader(os.Stdin)
// 	wr := bufio.NewWriter(os.Stdout)
// 	defer wr.Flush()

// 	var n, min int
// 	fmt.Fscanln(rd, &n)

// 	line, _, _ := rd.ReadLine()
// 	arr := strings.Fields(string(line))

// 	for _, v := range arr {
// 		min, _ = strconv.Atoi(arr[0])
// 		v, _ := strconv.Atoi(v)

// 		if min > v {
// 			min = v
// 			wr.WriteString(strconv.Itoa(min))
// 		}
// 	}
// 	fmt.Fprintln(wr, min)
// }
