package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

// call by value
func swap(x, y int) (int, int) { // 리턴을 2개 이상 받기
	return y, x
}

// 재귀 호출 1
func recursiveCall(x int) {
	if x == 0 { // 탈출 조건
		return
	}

	fmt.Printf("before call : x=%v\n", x)

	recursiveCall(x - 1)

	fmt.Printf("after call : x=%v\n", x)
}

// 재귀 호출 2
func sum(x int, result int) int {
	if x == 0 {
		return result // 최종 return
	}

	result += x

	return sum(x-1, result) // 최종 return을 제외한 return
}

// 재귀 호출 3
func sumPointerResult(x int, result *int) {
	if x == 0 {
		return
	}

	*result += x

	sumPointerResult(x-1, result)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}

	a := 10
	b := 20

	fmt.Printf("a=%v, b=%v\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("a=%v, b=%v\n", a, b)

	fmt.Println("-------------------------------")
	recursiveCall(10)

	fmt.Println("-------------------------------")

	result := sum(10, 0)
	fmt.Printf("result=%v\n", result)

	fmt.Println("-------------------------------")
	var pointerResult int
	sumPointerResult(10, &pointerResult)
	fmt.Printf("pointerResult=%v\n", pointerResult)

	fmt.Println("-------------------------------")
	fiboResult := fibo(10)
	fmt.Printf("fiboResult=%v\n", fiboResult)
}

func fibo(x int) int {
	if x == 0 || x == 1 {
		return 1
	}

	return fibo(x-1) + fibo(x-2)
}
