package main

import (
	"fmt"
)

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 && array[j] < array[j-1] {
			array[j], array[j-1] = array[j-1], array[j]
			j--
		}
	}

	return array
}

func main() {
	fmt.Println(InsertionSort([]int{-1, 0, 2, 2, 33, -292992}))
}
