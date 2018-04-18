package main

import (
	"fmt"
)

//----------------------------------------
// Person is a structure.
type Person struct {
	name  string
	age   int
	phone string
}

//----------------------------------------
// ShowPerson is a print function
func (p *Person) ShowPerson() {
	fmt.Printf("p.name=%v\n", p.name)
	fmt.Printf("p.age=%v\n", p.age)
	fmt.Printf("p.phone=%v\n", p.phone)
}

type Student struct {
	name  string
	class int

	subject Subject // has-a
}

type Subject struct {
	name  string
	grade string
}

func (s *Student) ViewSubject() {
	fmt.Println(s.subject)
}

//func (s Student) InputSubject(name string, grade string) {
func (s *Student) InputSubject(name string, grade string) {
	s.subject.name = name
	s.subject.grade = grade
}

func ViewSubject(s *Student) { // 주소 복사
	fmt.Println(s.subject)
}

func ViewSubjectVal(s Student) { // 객체 전체 복사
	fmt.Println(s.subject)
}

func main() {
	var p Person
	p1 := Person{"홍길동", 25, "01024586594"}
	p2 := Person{name: "김영수",
		age: 30}
	p3 := Person{}

	fmt.Println("------------------------")
	p1.ShowPerson()
	fmt.Println("------------------------")
	fmt.Println(p, p1, p2, p3)

	fmt.Println("------------------------")
	var st Student
	st.name = "재인"
	st.class = 3
	st.subject.name = "영어"
	st.subject.grade = "A"

	st.ViewSubject()   // 주소 전달
	ViewSubject(&st)   // 주소 전달
	ViewSubjectVal(st) // 객체 전체 복사해서 전달

	fmt.Println("------------------------")
	st.InputSubject("과학", "B")
	st.ViewSubject()
}
