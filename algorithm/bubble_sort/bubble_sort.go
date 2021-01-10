package main

import (
	"fmt"
)

func BubbleSort(array []int) []int {
	isSorted := false
	count := 0
	n := len(array) - 1

	for !isSorted {
		isSorted = true
		for i := 0; i < n-count; i++ {
			if array[i] > array[i+1] {
				temp := array[i]
				array[i] = array[i+1]
				array[i+1] = temp
				isSorted = false
			}
		}
		count++
	}
	return array
}

func main() {
	fmt.Println(BubbleSort([]int{-1, 0, 2, 2, 33, -292992}))
}
