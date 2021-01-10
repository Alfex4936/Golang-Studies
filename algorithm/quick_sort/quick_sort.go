package main

import (
	"fmt"
)

func QuickSort(array []int) []int {
	return QuickSortHelper(array, 0, len(array)-1)
}

func QuickSortHelper(array []int, lo, hi int) []int {
	for hi > lo {
		p := partition(array, lo, hi)
		QuickSortHelper(array, lo, p)
		lo = p + 1
	}
	return array
}

func partition(array []int, lo, hi int) int {
	i := lo - 1
	j := hi + 1
	pivot := array[(lo+hi)/2]

	for {
		i++
		for array[i] < pivot {
			i++
		}
		j--
		for array[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}

		array[i], array[j] = array[j], array[i]
	}
}

func main() {
	fmt.Println(QuickSort([]int{-1, 0, 2, 2, 33, -292992}))
}
