package main

import (
	"fmt"
)

func main() {
	/*
		slice는 한번에 메모리 확보하는 방법에 대한
		생각이 필요하다.
		여유 분을 확보하면 성능에 도움이 된다.
		- 확보 공간 : capacity
		- 현재 실제 사용하는 공간이 몇 개 인가 : 크기(size)
	*/

	fmt.Println("--------------------------------------")
	//---------------------------------------------------
	// 1. slice 확인
	var a []int // 동적 배열 선언(빈 공간의 동적 배열)
	fmt.Printf("len(a)=[%v], cap(a)=[%v]\n", len(a), cap(a))

	aa := []int{1, 2, 3, 4} // 동적 배열 4개 공간 확보하고 값을 초기화 함
	fmt.Printf("len(aa)=[%v], cap(aa)=[%v]\n", len(aa), cap(aa))

	aaa := make([]int, 3) // 초기 크기(size) : 3
	fmt.Printf("len(aaa)=[%v], cap(aaa)=[%v], aaa=[%v]\n", len(aaa), cap(aaa), aaa)

	aaaa := make([]int, 0, 8) // 크기(size) : 0, capacity : 8
	fmt.Printf("len(aaaa)=[%v], cap(aaaa)=[%v], aaaa=[%v]\n", len(aaaa), cap(aaaa), aaaa)

	aaaaa := make([]int, 3, 8) // 8개 공간을 확보하고, 3개의 배열에 0으로 초기화 함.
	fmt.Printf("len(aaaaa)=[%v], cap(aaaaa)=[%v], aaaaa=[%v]\n", len(aaaaa), cap(aaaaa), aaaaa)

	sliceAppend(&a)
	fmt.Printf("len(a)=[%v], cap(a)=[%v], a=[%v]\n", len(a), cap(a), a)
	a = append(a, 200)
	fmt.Printf("len(a)=[%v], cap(a)=[%v], a=[%v]\n", len(a), cap(a), a)

	appendResultAddress()
	makeSlice()
	simpleSlice()
	forceCopySlice()
	/*
		var f [10]int // 고정 배열 10개
	*/

}

func sliceAppend(input *[]int) {
	fmt.Println("--------------------------------------")
	//---------------------------------------------------
	// 2. slice에 추가 : append()
	*input = append(*input, 100)
	fmt.Printf("len(*input)=[%v], cap(*input)=[%v], *input=[%v]\n", len(*input), cap(*input), *input)

}

func appendResultAddress() {
	fmt.Println("--------------------------------------")
	//---------------------------------------------------
	a := []int{1, 2}
	b := append(a, 3) // 이전 크기의 2배로 cap()을 만든다.

	fmt.Printf("a=[%p], b=[%p]\n", a, b)

	fmt.Println("--------------------------------------")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v, ", a[i])
	}

	fmt.Println("\n--------------------------------------")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%v, ", b[i])
	}

	fmt.Println("\n--------------------------------------")
	fmt.Printf("cap(a)=[%v], cap(b)=[%v]\n", cap(a), cap(b))
}

func makeSlice() {
	fmt.Println("--------------------------------------")
	a := make([]int, 2, 4) // len = 2, cap = 4
	a[0] = 1
	a[1] = 2
	b := append(a, 3)

	fmt.Printf("a=[%p], b=[%p]\n", a, b) // a와 b는 같은 주소를 가르키고 있다.
	fmt.Printf("a=[%v], b=[%v]\n", a, b)

	b[0] = 4
	b[1] = 5
	fmt.Printf("a=[%v], b=[%v]\n", a, b)
}

func simpleSlice() {
	fmt.Println("--------------------------------------")
	a := []int{1, 2}  // len = 2, cap = 4
	b := append(a, 3) // 여유 공간이 없으면, 현재 공간의 2배를 할당한다.

	fmt.Printf("a=[%p], b=[%p]\n", a, b) // a와 b는 같은 주소를 가르키고 있다.
	fmt.Printf("a=[%v], b=[%v]\n", a, b)

	b[0] = 4
	b[1] = 5
	fmt.Printf("a=[%v], b=[%v]\n", a, b)
}

func forceCopySlice() {
	fmt.Println("--------------------------------------")
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2

	b := make([]int, len(a))
	for i := 0; i < len(a); i++ { // 복사 코드
		b[i] = a[i]
	}

	b = append(b, 3)
	b[0] = 10
	b[1] = 20

	fmt.Printf("a=[%p], b=[%p]\n", a, b)
	fmt.Printf("a=[%v], b=[%v]\n", a, b)
}
