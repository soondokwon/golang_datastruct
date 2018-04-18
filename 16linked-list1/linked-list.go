package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	num  int
	name string
	next *Node
}

func main() {
	var pRoot *Node // nil pointing
	pRoot = &Node{num: 0, name: "benny"}

	for i := 1; i < 10; i++ {
		AddNode(pRoot, i, "john"+strconv.Itoa(i))
	}

	PrintNode(pRoot)
}

func AddNode(pInput *Node, newNum int, newName string) {
	//--------------------------------
	// root 노드를 가르 킨다.
	var pTail *Node // nil pointing
	pTail = pInput

	//--------------------------------
	// 마지막 노드로 이동
	for pTail.next != nil {
		pTail = pTail.next
	}

	//--------------------------------
	// 새 노드 추가
	node := &Node{num: newNum, name: newName}
	pTail.next = node
}

func PrintNode(pCur *Node) {
	for pCur.next != nil { // 끝 까지
		fmt.Printf("pCur.num=[%v], pCur.name=[%v]\n", pCur.num, pCur.name)
		//fmt.Printf("pCur=[%v]\n", pCur)

		pCur = pCur.next
	}
}
