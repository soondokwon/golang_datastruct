package main

import (
	"fmt"
)

func main() {
	a := 4
	b := 2

	/* 변수 선언 예제
	var c int
	c = 5
	var d = 6
	var e = 3.14
	var f int = 8
	*/

	// bit 연산자 : and, or, xor
	fmt.Printf("a=[%v], b=[%v]\n", a, b)
	fmt.Printf("a&b=[%v]\n", a&b) // bit and
	fmt.Printf("a|b=[%v]\n", a|b) // bit or
	fmt.Printf("a^b=[%v]\n", a^b) // bit xor

	// bit 연산자 응용 : 자리수 찾기, clear
	c := 256
	d := c % 10
	e := c / 10 % 10
	f := c / 100 % 10
	clear := c &^ c // clear 연산

	fmt.Printf("일의 자리=[%v], 십의 자리=[%v] 백의 자리=[%v]\n", d, e, f)
	fmt.Printf("clear=[%v]\n", clear)

	// bit shift 연산자 : <<(x2), >>(/2) -> x, /보다 shift가 빠르다
	g := 4
	fmt.Printf("4 << = [%v]\n", g<<1)
	fmt.Printf("4 >> = [%v]\n", g>>1)
}
