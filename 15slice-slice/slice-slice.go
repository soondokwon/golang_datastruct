package main

import (
	"fmt"
)

func main() {
	// 분분 선택 slice는 범위 포인터(pointer)로 봐야 한다.
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("a=[%v]\n", a)

	b := a[4:8] // index번호 4부터 7까지 : pointer
	fmt.Printf("b=[%v]\n", b)

	c := a[4:] // index번호 4부터 끝 까지 : pointer
	fmt.Printf("c=[%v]\n", c)

	d := a[:4] // index번호 처음부터 3까지 : pointer
	fmt.Printf("d=[%v]\n", d)

	for i := 0; i < len(a); i++ {
		a[i] *= 10
	}

	fmt.Println("--------------------------")

	fmt.Printf("b=[%v]\n", b)
	fmt.Printf("c=[%v]\n", c)
	fmt.Printf("d=[%v]\n", d)

	fmt.Println("--------------------------")

	for i := 0; i < 5; i++ {
		var tail int
		a, tail = RemoveTail(a)
		fmt.Printf("tail=[%v]\n", tail)
	}

	fmt.Println("--------------------------")
	fmt.Printf("a=[%v]\n", a)

	fmt.Println("--------------------------")
	for i := 0; i < 3; i++ {
		var front int
		a, front = RemoveFront(a)
		fmt.Printf("front=[%v]\n", front)
	}
	fmt.Println("--------------------------")
	fmt.Printf("a=[%v]\n", a)

}

func RemoveTail(in []int) ([]int, int) {
	return in[:len(in)-1], in[len(in)-1]
}

func RemoveFront(in []int) ([]int, int) {
	return in[1:], in[0]
}
