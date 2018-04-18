package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	clone := [5]int{}
	// 1. 배열 복사
	for i := 0; i < len(arr); i++ {
		clone[i] = arr[i]
	}

	fmt.Printf("arr=%v\n", arr)
	fmt.Printf("clone=%v\n", clone)

	fmt.Println("------------------------------")
	// 2. temp 변수를 사용해서 array reverse 처리 : 2n
	temp := [5]int{}
	for i := 0; i < len(arr); i++ {
		temp[i] = arr[len(arr)-1-i]
	}
	fmt.Printf("temp=%v\n", temp)

	fmt.Println("------------------------------")
	for i := 0; i < len(temp); i++ {
		arr[i] = temp[i]
	}
	fmt.Printf("arr=%v\n", arr)

	fmt.Println("------------------------------")
	// 3. 위치 교환으로 array reverse 처리 : n/2
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Printf("arr=%v\n", arr)

	fmt.Println("------------------------------")
	// 4. radix sort : 정수로 크기가 정해져 있어야 한다. 너무 크면 안됨.
	arrSortSrc := [15]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 4, 5, 1, 3, 7, 8}
	arrMap := [10]int{}

	for i := 0; i < len(arrSortSrc); i++ {
		arrMap[arrSortSrc[i]]++
	}

	idx := 0
	for i := 0; i < len(arrMap); i++ {
		for j := 0; j < arrMap[i]; j++ {
			arrSortSrc[idx] = i
			idx++
		}
	}

	fmt.Printf("arrSortSrc=%v\n", arrSortSrc)
}
