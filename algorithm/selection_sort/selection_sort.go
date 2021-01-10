package main

import (
	"fmt"
)

func SelectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		smallest := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[smallest] {
				smallest = j
			}
		}
		array[i], array[smallest] = array[smallest], array[i]
	}

	return array
}

func main() {
	fmt.Println(SelectionSort([]int{-1, 0, 2, 2, 33, -292992}))
}
