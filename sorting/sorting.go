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
