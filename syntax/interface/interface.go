package main

import "fmt"

type Duck interface {
	Say()
	Move()
}

type A struct{}
type B struct{}

func (a A) Say() {
	fmt.Println("A called")
}
func (a A) Move() {}

func (b B) Say() {
	fmt.Println("B called")
}
func (b B) Move() {}

func Call(d ...Duck) {
	for _, duck := range d {
		duck.Say()
	}
}

func main() {
	a := A{}
	b := B{}

	Call(a, b)
}
