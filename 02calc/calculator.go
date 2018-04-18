package main

import "fmt"

func main() {
	a := 10   // 변수 선언 및 대입
	var b int // 선언
	var c rune
	var d rune

	var str = "안녕하세요...!!!" // string == []rune
	str2 := "대입과 선언 동시 -> "
	var str3 string

	str3 = "선언 후 대입하여 사용하기"

	a = 3 // 값 바꾸기
	b = 4 // 선언 후 대입
	c = 'a'
	d = '한'

	fmt.Println(str, a+b)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println("a=", c)
	fmt.Println("한", d)

}
