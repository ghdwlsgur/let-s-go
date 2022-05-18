package main

/*
	패키지
	패키지는 코드를 구조화하고 재사용하기 위한 단위로, 다른 언어의 모듈이나 라이브러리와 유사한 개념이다.
	모든 Go 프로그램은 패키지로 구성되어 있고, 한 패키지에서 다른 패키지를 임포트하여 사용할 수 있다.
	Go는 항상 패키지 선언(package 키워드를 사용하여 선언)으로 코드를 시작한다.
	패키지 이름과 디렉터리 이름은 같아야 하고, 같은 패키지에 있는 소스 파일은 모두 같은 디렉터리에
	있어야 한다.
*/

import (
	"fmt"
	"os"
)

func main() {
	var name string

	fmt.Print(" 이름을 입력하세요: ")
	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hello %s\n", name)
}

/*
패키지는 크게 두 가지로 나눌 수 있다. 하나는 명령 프롬프트에서 명령을 내려 실행할 수 있는
실행 가능한 프로그램이다. 다른 하나는 다른 프로그램에서 호출하여 사용할 수 있도록 연관된 작업을
하는 코드 묶음인 라이브러리이다.

패키지 이름이 main이면 Go는 실행 가능한 프로그램으로 인식한다.
main 패키지를 빌드하면 디렉터리 이름과 같은 이름으로 실행 파일이 생성되고, 프로그램을 실행하면
main 패키지의 main() 함수를 찾아서 실행한다.
*/