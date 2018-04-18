package main

import (
	"fmt"
)

func main() {
	var A [10]int
	B := [10]int{}
	C := [10]int{1, 2, 4, 8, 16}
	// var D [10]int = {1,2,4} // error

	fmt.Printf("A=%v\n", A)
	fmt.Printf("B=%v\n", B)
	fmt.Printf("C=%v\n", C)

	fmt.Println("---------------------------")

	for i := 0; i < len(A); i++ {
		A[i] = i * i
	}

	fmt.Printf("A=%v\n", A)
	fmt.Printf("B=%v\n", B)
	fmt.Printf("C=%v\n", C)

	fmt.Println("---------------------------")

	// UTF8은 1~3 byte가 될 수 있다.
	// string은 배열(array)이다.
	S := "Hello 월드" // 월 : 3 bytes, 드 : 3 bytes
	fmt.Printf("len(S)=%v\n", len(S))
	for i := 0; i < len(S); i++ {
		fmt.Print(string(S[i]), "=", S[i], ", ") // 문자 코드 출력
		if i == len(S)-1 {
			fmt.Println("")
		}
	}

	fmt.Println("---------------------------")
	/*
		1. [] byte : 무조건 1 byte 단위의 배열
		2. [] rune : UTF8을 한 개의 단위의 배열로 처리(1~3byte)
	*/

	// string -> []rune
	rS := []rune(S)
	fmt.Printf("len(rS)=%v\n", len(rS))
	for i := 0; i < len(rS); i++ {
		fmt.Print(string(rS[i]), "=", rS[i], ", ") // 문자 코드 출력
	}

}
