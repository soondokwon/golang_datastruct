package main

import (
	"fmt"
)

func main() {
	var a bool
	var b bool
	var c bool
	val := 80

	a = 3 > 4
	b = 3 < 4
	c = 3 < 4 && 4 < 5 // and

	fmt.Printf("a=[%v]\n", a) // false
	fmt.Printf("b=[%v]\n", b) // true
	fmt.Printf("c=[%v]\n", c) // true

	//--------------------------------
	// if
	if a {
		fmt.Println("a=참")
		return
	}
	fmt.Println("a=거짓")

	//--------------------------------
	// if - else
	if c {
		fmt.Println("c=참")
	} else {
		fmt.Println("c=거짓")
	}

	if d := true; d {
		fmt.Println("d=참")
	} else {
		fmt.Println("d=거짓")
	}

	//--------------------------------
	// if - else if - else
	if val >= 90 {
		fmt.Println("학점 : A")
	} else if val >= 80 {
		fmt.Println("학점 : B")
	} else if val >= 70 {
		fmt.Println("학점 : C")
	} else {
		fmt.Println("학점 : F")
	}

	if val = 91; val >= 90 {
		fmt.Println("학점 : A")
	} else if val >= 80 {
		fmt.Println("학점 : B")
	} else if val >= 70 {
		fmt.Println("학점 : C")
	} else {
		fmt.Println("학점 : F")
	}

	//--------------------------------
	if avg := 50; avg >= 75 && avg <= 85 {
		fmt.Println("중앙 점수 입니다.")
	} else {
		fmt.Println("중앙 점수가 아닙니다.")
	}

}
