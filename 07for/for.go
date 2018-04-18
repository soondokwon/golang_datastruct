package main

import (
	"fmt"
)

func forSimpleLoop() {
	i := 0       // Loop 지나고 사용할 때 구조
	for i < 10 { // for 조건 {}
		fmt.Println(i)
		i++
	}
}

func forLoop(size int) {
	// Loop 내에서만 사용할 때 구조
	for i := 0; i < size; i++ { // for 전처리 ; 조건 ; 후처리 {}
		fmt.Println(i)
	}
}

func forUnlimitedLoop() {
	i := 0

	for { // 무한 루프
		fmt.Println("forUnlimitedLoop() called... i=", i)
		if i == 5 {
			fmt.Println("continue... i=", i)

			i += 2   // 6 통과
			continue // 현재 loop의 끝으로 이동 후, 다음 loop 준비
		}

		if i >= 10 { // 종료 조건
			break // 현재 loop를 종료
		}

		i++
	}
}

func main() {
	forSimpleLoop()
	fmt.Println("----------------------------")
	forLoop(10)
	forUnlimitedLoop()
}
