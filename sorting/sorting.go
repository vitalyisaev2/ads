package sorting

import (
	"sort"
)

// Skiena 2.5.1
func SelectionSort(data sort.Interface) {

	var length = data.Len()

	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}

// Skiena 2.5.2
func InsertionSort(data sort.Interface) {

	var length = data.Len()

	for i := 0; i < length-1; i++ {
		j := i + 1
		for j > 0 && data.Less(j, j-1) {
			data.Swap(j, j-1)
			j -= 1
		}
	}
}
