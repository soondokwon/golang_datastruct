package main

import (
	"fmt"
)

func printIncStar(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func printDescStar(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size-i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func main() {
	fmt.Println("----------------------------")

	for i := 2; i <= 9; i++ {
		/*
			if i == 5 { // 5단 건너 뛰기
				continue
			}
		*/

		for j := 1; j <= 9; j++ {

			fmt.Printf("%v x %v = %v\n", i, j, i*j)

		}
		fmt.Println("----------------------------")
	}

	printIncStar(10)
	fmt.Println("----------------------------")
	printDescStar(10)
}
