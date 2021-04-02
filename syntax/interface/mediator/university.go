package main

import (
	"fmt"
)

// University ...
type University interface {
	MakePublic()
	GetMoreStudents()
	IncreaseTuition()
}

type A_UNIV struct {
	Mediator  Mediator
	students  int
	fee       int
	isPrivate bool
}

type B_UNIV struct {
	Mediator  Mediator
	students  int
	fee       int
	isPrivate bool
}

func (a *A_UNIV) MakePublic() {
	if !a.Mediator.satisfy(a) {
		fmt.Println("There are too many universities now.")
	} else {
		fmt.Println("A university becomes an official university.")
	}
}

// GetMoreStudents 학생 수 +
func (a *A_UNIV) GetMoreStudents(students int) {
	if !a.Mediator.isFull(a, students) {
		fmt.Println("A university can accept more students.")
		a.students += students
	} else {
		fmt.Println("A university cannot accept more students.")
	}
}

// 등록금 +
func (a *A_UNIV) IncreaseTuition(money int) {
	if !a.Mediator.isExpensive(a, money) {
		fmt.Sprintln("A university increases tuition by %d won", money)
		a.fee += money
	} else {
		fmt.Println("A university must maintain current tuition.")
	}

}

func (b *B_UNIV) MakePublic() {
	if !b.Mediator.satisfy(b) {
		fmt.Println("There are too many universities now.")
	} else {
		fmt.Println("A university becomes an official university.")
	}
}

// IncreaseTuition 등록금 +
func (b *B_UNIV) IncreaseTuition(money int) {
	if !b.Mediator.isExpensive(b, money) {
		fmt.Sprintln("B university increases tuition by %d won", money)
		b.fee += money
	} else {
		fmt.Println("B university must maintain current tuition.")
	}

}

// GetMoreStudents 학생 수 +
func (b *B_UNIV) GetMoreStudents(students int) {
	if !b.Mediator.isFull(b, students) {
		fmt.Println("B university can accept more students.")
		b.students += students
	} else {
		fmt.Println("B university cannot accept more students.")
	}
}
