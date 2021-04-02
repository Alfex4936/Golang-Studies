package main

type Mediator interface {
	isFull(students int)
	IncreaseTuition(fee int)
}
