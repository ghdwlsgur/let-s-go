package main

import (
	"fmt"
	"image/color"
)

func main() {
	// var vRect *rect
	// var pInt *int
	// var pFloat *float64
	// var pComplex *complex128

	var p *int

	i := 1
	p = &i
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(*p)
	fmt.Println(p)

	var k *int
	var kk **int

	// j <-- k <--- kk
	j := 1
	k = &j  // k는 j의 메모리 주소값을 참조
	kk = &k // kk는 k의 메모리 주소값을 참조
	fmt.Println(j, *k, **kk)

	j += 1
	fmt.Println(j, *k, **kk)

	*k++
	fmt.Println(j, *k, **kk)

	**kk++
	fmt.Println(j, *k, **kk)

	type rect struct{ w, h float64 }
	var g int = 1
	var w *int = &g
	var s *rect = &rect{1, 2}

	fmt.Println(w)
	fmt.Println(s)

	a := new(int)
	*a = 1
	fmt.Println(a)
	fmt.Println(*a)

	type rect2 struct{ w, h float64 }
	r := new(rect2)
	r.w, r.h = 3, 4
	fmt.Println(r)
	fmt.Println(*r)

	/*
		Go에서는 함수나 메서드를 호출할 때 매개변수 값을 복사해서 함수나 메서드 내부로 전달한다.
		그래서 함수나 메서드 내부에서는 전달된 매개변수의 원본 값을 변경할 수 없다.
		함수나 메서드 내에서 원본 값을 변경하려면 포인터를 사용해야 한다.

		포인터 값의 크기는 64비트 머신에서는 8바이트, 32비트 머신에서는 4바이트이다. 함수나 메서드로
		포인터를 전달하면 실제 포인터가 가리키고 있는 값에 상관없이 아주 적은 양의 값만 전달된다.

		불 타입이나 숫자 타입 값은 크기가 1바이트에서 8바이트 정도이므로 함수나 메서드 내부로 값을 전달할 때
		값을 복사하는 작업이 시스템에 부담을 주지 않는다. 문자열을 전달하는 것도 가볍게 동작한다.

		Go 런타임은 문자열을 전달할 때 실제 그 값이 얼마나 큰지에 상관없이 데이터를 아주 적은 양만
		전달하도록 최적화되어 있기 때문이다. (문자열 하나를 함수나 메서드에 전달할 때 실제 전달되는 값은
		64비트 머신에서는 16바이트, 32비트 머신에서는 8바이트이다.) 물론 += 같은 작업을 할 때는 문자열
		전체를 복사한다.

		C나 C++와 달리 Go의 배열은 값을 복사해서 함수나 메서드 내부로 전달한다.
		그래서 큰 배열을 함수나 메서드로 전달하면 시스템에 부담이 많이 된다. 길이가 큰 배열을 함수나 메서드로
		전달할 때는 배열 대신 슬라이스를 사용하는 것이 좋다. 함수나 메서드로 슬라이스를 전달하면 참조가
		전달되므로 시스템에 덜 부담된다.
	*/

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	multiply(numbers, 5)
	fmt.Println(numbers)

	/*
		크기가 큰 구조체를 함수나 메서드에 전달할 때는 값 전체를 복사해서 전달하므로 시스템에 부담을 많이 준다.
		이럴 때는 구조체 값을 직접 전달하지 않고 포인터로 전달하면 시스템에 주는 부담을 줄일 수 있다.
	*/

	type rect3 struct {
		x0, y0, x1, y1 int
		color.RGBA
	}

	rect4 := rect3{2, 4, 10, 20, color.RGBA{0xFF, 0, 0, 0xFF}}
	fmt.Println(rect4)

}

func multiply(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}
