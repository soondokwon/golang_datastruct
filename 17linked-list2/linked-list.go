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
	var pTail *Node // nil pointing
	pRoot = &Node{num: 0, name: "benny"}
	pTail = pRoot

	for i := 1; i < 10; i++ {
		pTail = AddNode(pTail, i, "john"+strconv.Itoa(i))
	}

	fmt.Println("-------------------------------")
	PrintNode(pRoot)

	fmt.Println("-------------------------------")
	pRoot, pTail = RemoveNode(2, pRoot, pTail)

	fmt.Println("-------------------------------")
	PrintNode(pRoot)

	fmt.Println("-------------------------------")
	pRoot, pTail = RemoveNode(4, pRoot, pTail)

	fmt.Println("-------------------------------")
	PrintNode(pRoot)

	fmt.Println("-------------------------------")
	pRoot, pTail = RemoveNode(0, pRoot, pTail)

	fmt.Println("-------------------------------")
	PrintNode(pRoot)

	fmt.Println("-------------------------------")
	pRoot, pTail = RemoveNode(8, pRoot, pTail)

	fmt.Println("-------------------------------")
	PrintNode(pRoot)

}

func AddNode(pInput *Node, newNum int, newName string) *Node {
	// 새 노드 추가
	node := &Node{num: newNum, name: newName}
	pInput.next = node

	return node
}

func getCurNodeByNum(num int, pRoot *Node) *Node {
	pCur := pRoot
	var pResult *Node

	for { // 끝 까지
		if pCur.num == num {
			fmt.Printf("pCur=[%v]\n", pCur)
			pResult = pCur
			break
		}

		if pCur.next == nil { // 무한 루트 종료 조건
			break
		}

		pCur = pCur.next
	}

	return pResult
}

func RemoveNode(num int, pRoot *Node, pTail *Node) (*Node, *Node) {
	pCur := getCurNodeByNum(num, pRoot)

	// 1. root를 지울 떄
	if pCur == pRoot {
		pRoot = pRoot.next
		if pRoot == nil { // node가 1개 일때
			pTail = nil
		}

		return pRoot, pTail
	}

	//-------------------------------
	// prev 노드 찾기
	pPrev := pRoot
	for pPrev.next != pCur {
		pPrev = pPrev.next
	}

	// 2. tail을 지울 때
	if pCur == pTail {
		pPrev.next = nil
		pTail = pPrev

		return pRoot, pTail
	}

	// 3. root, tail을 제외한 노드 지울 때
	pPrev.next = pPrev.next.next

	return pRoot, pTail
}

func PrintNode(pCur *Node) {
	for { // 끝 까지
		fmt.Printf("pCur.num=[%v], pCur.name=[%v]\n", pCur.num, pCur.name)
		//fmt.Printf("pCur=[%v]\n", pCur)
		if pCur.next == nil {
			break
		}

		pCur = pCur.next
	}
}
