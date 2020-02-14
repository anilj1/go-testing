package services

import (
	"api/utils/sort"
)

func Sort(elements []int) {
	if len(elements) <= 20000 {
		sort.BubbleSort(elements)
	} else {
		sort.Sort(elements)
	}
}