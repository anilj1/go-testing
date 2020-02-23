package services

import (
	"api/utils/sort"
)

func init() {
}

func Sort(elements []int) {
	if len(elements) <= 20000 {
		sort.BubbleSort(elements)
	} else {
		sort.Sort(elements)
	}
}