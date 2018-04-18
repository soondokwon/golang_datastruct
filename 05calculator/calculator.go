package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcIf(n1 int, n2 int, op string) {
	if op == "+" {
		fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
	} else if op == "-" {
		fmt.Printf("%d - %d = %d\n", n1, n2, n1-n2)
	} else if op == "*" {
		fmt.Printf("%d * %d = %d\n", n1, n2, n1*n2)
	} else if op == "/" {
		fmt.Printf("%d / %d = %d\n", n1, n2, n1/n2)
	} else {
		fmt.Printf("잘못된 연산자입니다.\n")
	}
}

func calcSwitch(n1 int, n2 int, op *string) { // op : call by pointer
	switch *op {
	case "+":
		fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d\n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d\n", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%d / %d = %d\n", n1, n2, n1/n2)
	default:
		fmt.Printf("잘못된 연산자입니다.\n")
	}

	*op += " : call by pointer" // out value
}

func main() {
	fmt.Print("숫자 입력 : ")

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n') // 개행 문자 까지 1줄을 읽는다.
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line) // StrToInt

	fmt.Print("숫자 입력 : ")
	line, _ = reader.ReadString('\n') // 개행 문자 까지 1줄을 읽는다.
	line = strings.TrimSpace(line)    // 비정상 문자 제거

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력한 숫자는 [%d, %d]\n", n1, n2)

	fmt.Print("연산자 입력(+, -, *, /) : ")
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	// calcIf(n1, n2, line)
	calcSwitch(n1, n2, &line)

	fmt.Printf("out value : line=[%v]\n", line)
}
