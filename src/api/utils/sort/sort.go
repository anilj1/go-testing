package sort

import (
	"sort"
)

const (
	privateConst = "private"
)

func GetElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n-1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}

// This sort algorithm performs better for small data set
func BubbleSort(elements []int) {
	keepWorking := true
	for keepWorking {
		keepWorking = false

		for i := 0; i < len(elements) - 1 ; i++ {
			if elements[i] > elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

// This sort algorithm performs better for large data set
func Sort(elements []int) {
	sort.Ints(elements)
}
