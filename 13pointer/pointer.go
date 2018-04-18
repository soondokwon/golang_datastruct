package main

import (
	"fmt"
)

type Student struct {
	number string // 학번
	name   string // 이름
	age    int    // 나이
	score  int    // 학점
	grade  string // 학년
}

func (this *Student) PrintStudent() {
	fmt.Printf("this.number=[%v]\n", this.number)
	fmt.Printf("this.name=[%v]\n", this.name)
	fmt.Printf("this.age=[%v]\n", this.age)
	fmt.Printf("this.score=[%v]\n", this.score)
	fmt.Printf("this.grade=[%v]\n", this.grade)
}

func (this *Student) UpdateScore(score int) {
	this.score = score
}

func (this *Student) UpdateGrade(grade string) {
	this.grade = grade
}

func main() {
	var iA int = 20
	iB := 30
	var pTmp *int // default : nil

	fmt.Println("----------------------------")
	fmt.Printf("pTmp=[%v]\n", pTmp)

	fmt.Println("----------------------------")
	pTmp = &iA
	fmt.Printf("pTmp=[%v] *pTmp=[%v]\n", pTmp, *pTmp)

	fmt.Println("----------------------------")
	pTmp = &iB
	fmt.Printf("pTmp=[%v] *pTmp=[%v]\n", pTmp, *pTmp)

	fmt.Println("---------------------------------")
	var s Student = Student{
		number: "9501007", name: "benny",
		age: 20, score: 80,
		grade: "2학년"}

	s2 := Student{
		number: "9301007", name: "soondo",
		age: 22, score: 95,
		grade: "4학년"}

	s.PrintStudent()
	fmt.Println("---------------------------------")

	s.UpdateScore(90)
	s.UpdateGrade("3학년")
	s.PrintStudent()
	fmt.Println("---------------------------------")

	s2.PrintStudent()
	fmt.Println("---------------------------------")
}
