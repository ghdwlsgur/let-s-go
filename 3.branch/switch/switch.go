package main

import "fmt"

func main() {

	i := -1
	// switch i {
	// case -1, -2:
	// 	fmt.Println(i, "는 음수입니다.")
	// case 1, 2:
	// 	fmt.Println(i, "는 양수입니다.")
	// }

	switch i {
	case -1, -2:
		fmt.Println("pass")
		fallthrough // fallthrough를 사용하면 다음 case로 이동하여 추가 작업을 수행할 수 있다.
	case 1, 2:
		fmt.Println(i, "는 숫자입니다.")
	}
}
