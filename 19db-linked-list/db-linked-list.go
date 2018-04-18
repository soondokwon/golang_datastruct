package main

import "fmt"

type Node struct {
	num  int
	name string

	prev *Node
	next *Node
}

type LinkedList struct {
	pRoot *Node
	pTail *Node
}

func (this *LinkedList) AddNode(newNum int, newName string) {
	// 첫번째 노드 일떄...
	if this.pRoot == nil {
		this.pRoot = &Node{num: newNum, name: newName}
		this.pTail = this.pRoot

		return
	}

	// 그 외의 경우
	this.pTail.next = &Node{num: newNum, name: newName}
	temp := this.pTail
	this.pTail = this.pTail.next
	this.pTail.prev = temp
}

func (this *LinkedList) getCurNodeByNum(num int) *Node {
	pCur := this.pRoot
	var pResult *Node

	for { // 무한 루프
		if pCur.num == num {
			fmt.Printf("pCur=[%v]\n", pCur)
			pResult = pCur
			break
		}

		if pCur.next == nil { // 무한 루프 종료 조건
			break
		}

		pCur = pCur.next
	}

	return pResult
}

func (this *LinkedList) RemoveNode(num int) {
	pCur := this.getCurNodeByNum(num)

	// 1. root를 지울 떄
	if pCur == this.pRoot {
		this.pRoot = this.pRoot.next
		this.pRoot.prev = nil
		if this.pRoot == nil { // node가 1개 일때
			this.pTail = nil
		}

		return
	}

	//-------------------------------
	// prev 노드 찾기
	pPrev := pCur.prev

	// 2. tail을 지울 때
	if pCur == this.pTail {
		pCur.prev = nil
		pPrev.next = nil
		this.pTail = pPrev

		return
	}

	// 3. root, tail을 제외한 노드 지울 때
	temp := pCur.prev
	pPrev.next = pCur.next
	pCur.next.prev = temp
	pCur.prev = nil
	pCur.next = nil
}

func (this *LinkedList) PrintNode() {
	pCur := this.pRoot

	for { // 끝 까지
		fmt.Printf("num=[%v], name=[%v]\n", pCur.num, pCur.name)
		//fmt.Printf("pCur=[%v]\n", pCur)
		if pCur.next == nil {
			break
		}

		pCur = pCur.next
	}
}

func (this *LinkedList) PrintNodeReverse() {
	pCur := this.pTail

	for { // 끝 까지
		fmt.Printf("num=[%v], name=[%v]\n", pCur.num, pCur.name)
		//fmt.Printf("pCur=[%v]\n", pCur)
		if pCur.prev == nil {
			break
		}

		pCur = pCur.prev
	}

}

func main() {
	linkedList := &LinkedList{}

	linkedList.AddNode(1, "benny")
	linkedList.AddNode(2, "john")
	linkedList.AddNode(3, "tommy")
	linkedList.AddNode(4, "jack")
	linkedList.AddNode(5, "allice")
	linkedList.AddNode(6, "jordan")
	linkedList.AddNode(7, "kwon")
	linkedList.AddNode(8, "kim")

	fmt.Println("--------------------------")
	linkedList.PrintNode()

	linkedList.RemoveNode(5)

	fmt.Println("--------------------------")
	linkedList.PrintNode()

	linkedList.RemoveNode(8)

	fmt.Println("--------------------------")
	linkedList.PrintNode()

	linkedList.RemoveNode(1)
	fmt.Println("--------------------------")
	linkedList.PrintNode()

	fmt.Println("--------------------------")
	linkedList.PrintNodeReverse()

}
