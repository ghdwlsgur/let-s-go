package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var val, n, m int
	fmt.Scanln(&n, &m)

	rd := bufio.NewReaderSize(os.Stdin, n<<3)
	wr := bufio.NewWriterSize(os.Stdout, m<<1)
	defer wr.Flush()

	// 3 4 5 6 입력했다면
	// [51 32 52 32 53 32 54]
	// 32는 공백
	// string(line) => 3 4 5 6
	line, is_prefix, _ := rd.ReadLine()
	// strings.Fields() 공백을 기준으로 슬라이스
	arr := strings.Fields(string(line))

	// 버퍼 초과할 경우
	for is_prefix {
		line, is_prefix, _ = rd.ReadLine()
		arr = append(arr, strings.Fields(string(line))...)
	}

	for _, v := range arr {
		val, _ = strconv.Atoi(v)
		if val < m {
			wr.WriteString(v)
			wr.WriteByte(' ')
		}
	}
}

func refactoring() {
	var val, n, m int
	fmt.Scanln(&n, &m)

	rd := bufio.NewReaderSize(os.Stdin, n<<3)
	wr := bufio.NewWriterSize(os.Stdout, m<<1)
	defer wr.Flush()

	line, is_prefix, _ := rd.ReadLine()
	arr := strings.Fields(string(line))

	for is_prefix {
		line, is_prefix, _ = rd.ReadLine()
		arr = append(arr, strings.Fields(string(line))...)
	}

	for _, v := range arr {
		val, _ = strconv.Atoi(v)
		if val < m {
			wr.WriteString(v)
			wr.WriteByte(' ')
		}
	}
}

func typing() {
	var val, n, m int
	fmt.Scanln(&n, &m)

	rd := bufio.NewReaderSize(os.Stdin, n<<3)
	wr := bufio.NewWriterSize(os.Stdout, m<<1)
	defer wr.Flush()

	line, is_prefix, _ := rd.ReadLine()
	arr := strings.Fields(string(line))

	for is_prefix {
		line, is_prefix, _ = rd.ReadLine()
		arr = append(arr, strings.Fields(string(line))...)
	}

	for _, v := range arr {
		val, _ = strconv.Atoi(v)
		if val < m {
			wr.WriteString(v)
			wr.WriteByte(' ')
		}
	}
}

func typing2() {
	var val, n, m int
	fmt.Scanln(&n, &m)

	rd := bufio.NewReaderSize(os.Stdin, n<<3)
	wr := bufio.NewWriterSize(os.Stdout, m<<1)
	defer wr.Flush()

	line, is_prefix, _ := rd.ReadLine()
	arr := strings.Fields(string(line))

	for is_prefix {
		line, is_prefix, _ = rd.ReadLine()
		arr = append(arr, strings.Fields(string(line))...)
	}

	for _, v := range arr {
		val, _ = strconv.Atoi(v)
		if val < m {
			wr.WriteString(v)
			wr.WriteByte(' ')
		}
	}
}

func typing3() {
	var val, n, m int
	fmt.Scanln(&n, &m)

	rd := bufio.NewReaderSize(os.Stdin, n<<3)
	wr := bufio.NewWriterSize(os.Stdout, m<<1)
	defer wr.Flush()

	line, is_prefix, _ := rd.ReadLine()
	arr := strings.Fields(string(line))

	for is_prefix {
		line, is_prefix, _ = rd.ReadLine()
		arr = append(arr, strings.Fields(string(line))...)
	}

	for _, v := range arr {
		val, _ = strconv.Atoi(v)
		if val < m {
			wr.WriteString(v)
			wr.WriteByte(' ')
		}
	}
}

func typing4() {
	var val, n, m int
	fmt.Scanln(&n, &m)

	rd := bufio.NewReaderSize(os.Stdin, n<<3)
	wr := bufio.NewWriterSize(os.Stdout, m<<1)
	defer wr.Flush()

	line, is_prefix, _ := rd.ReadLine()
	arr := strings.Fields(string(line))

	for is_prefix {
		line, is_prefix, _ = rd.ReadLine()
		arr = append(arr, strings.Fields(string(line))...)
	}

	for _, v := range arr {
		val, _ = strconv.Atoi(v)
		if val < m {
			wr.WriteString(v)
			wr.WriteByte(' ')
		}
	}

}
