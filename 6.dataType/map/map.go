package main

import "fmt"

func main() {
	numberMap := map[string]int{}
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	fmt.Println(numberMap)
	fmt.Println(numberMap["one"])
	fmt.Println(numberMap["two"])
	fmt.Println(numberMap["three"])

	numberMap2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println(numberMap2)

	numberMap3 := make(map[string]int, 3)
	numberMap3["one"] = 1
	numberMap3["two"] = 2
	numberMap3["three"] = 3
	fmt.Println(numberMap3)

	for k, v := range numberMap3 {
		fmt.Println(k, v)
	}

	// 슬라이스는 참조 타입이므로 맵의 키로 사용할 수 없다.
	// 하지만 []byte 타입과 []int32 타입은 string으로 변환하면 맵의 키로 사용할 수 있다.

	groupMap := make(map[string]string)

	group1 := []int32{1, 4, 6}
	group2 := []int32{2, 4, 5}
	group3 := []int32{4, 6, 7}

	// groupMap[group1] = "first"
	// groupMap[group2] = "second"
	// groupMap[group3] = "third"

	groupMap[string(group1)] = "first"
	groupMap[string(group2)] = "second"
	groupMap[string(group3)] = "third"

	fmt.Println(groupMap[string(group1)])
	for i, v := range groupMap {
		fmt.Println(i, v)
	}

	/*
		문자열은 유니코드 문자의 코드값을 정수로 표현한 값(rune 또는 int32)의 시퀀스이므로,
		[]int32 타입을 문자열로 변환할 수 있다. 하지만 문자의 코드값이 올바르지 않다면 문자열 형태로
		출력되지 않을 수 있다. 이 코드에서는 단지 실행할 때 []int32 타입 값을 문자열로 인식시켜
		맵의 키로 사용했다.
	*/

	if v, ok := numberMap["zero"]; ok {
		fmt.Println(" 'three' is in numberMap. value: ", v)
	} else {
		fmt.Println(" 'three' is not in numberMap")
	}

	if _, ok := numberMap["four"]; ok {
		numberMap["four"] = 4
		fmt.Println(numberMap["four"])
	} else {
		numberMap["four"] = 4
		fmt.Println(numberMap["four"])
	}

	fmt.Println("numberMap['four']:", numberMap["four"])

	appendMap := map[int]string{}
	appendMap[1] = "one"
	appendMap[2] = "two"
	appendMap[3] = "three"
	fmt.Println(appendMap)
	appendMap[3] = "삼"
	fmt.Println(appendMap)

	delete(appendMap, 3)
	fmt.Println(appendMap)

	delete(appendMap, 2)
	fmt.Println(appendMap)

	delete(appendMap, 1)
	fmt.Println(appendMap)

}
