package main

import (
	"fmt"
)

func MergeSort(array []int) []int {
	aux := make([]int, len(array))
	copy(aux, array)

	MergeSortHelper(array, 0, len(array)-1, aux)
	return array
}

func MergeSortHelper(main []int, lo, hi int, aux []int) {
	if lo == hi {
		return
	}

	mid := (lo + hi) / 2
	MergeSortHelper(aux, lo, mid, main)
	MergeSortHelper(aux, mid+1, hi, main)
	Merge(main, lo, mid, hi, aux)
}

func Merge(main []int, lo, mid, hi int, aux []int) {
	k := lo
	i := lo
	j := mid + 1

	for i <= lo && j <= hi {
		if aux[i] < aux[j] {
			main[k] = aux[i]
			i++
		} else {
			main[k] = aux[j]
			j++
		}
		k++
	}

	for i <= mid {
		main[k] = aux[i]
		i++
		k++
	}
	for j <= hi {
		main[k] = aux[j]
		j++
		k++
	}
}

func main() {
	fmt.Println(MergeSort([]int{-1, 0, 2, 3, 2, 33, 349, -292992}))
}
