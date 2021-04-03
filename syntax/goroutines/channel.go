package main

import (
	"fmt"
)

func main() {
	var numbers []int // nil
	done := make(chan bool)

	// start a goroutine to initialise array
	go func() {
		numbers = make([]int, 5000)
		done <- true
	}()

	<-done // one buffer
	numbers[0] = 5000
	fmt.Println(numbers[0]) // 5000
}
