package services

import (
	"testing"

	"api/utils/sort"
)

func TestSort(t *testing.T) {
	// Init
	elements := sort.GetElements(10)

	// Execution
	Sort(elements)
	// Validation
	if elements[0] != 0 {
		// t.Error does not stop executing remaining test case!
		t.Error("First element of sorted array should 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Last element of sorted array should 9")
	}
}

func TestSort20K(t *testing.T) {
	// Init
	elements := sort.GetElements(20001)

	// Execution
	Sort(elements)
	// Validation
	if elements[0] != 0 {
		// t.Error does not stop executing remaining test case!
		t.Error("First element of sorted array should 0")
	}

	if elements[len(elements)-1] != 20000 {
		t.Error("Last element of sorted array should 9")
	}
}

// This sort algorithm performs better for small data set
func BenchmarkBubbleSort10K(b *testing.B) {
	// Init
	elements := sort.GetElements(10000)

	// Execution
	for i :=0; i < b.N; i++ {
		Sort(elements)
	}
}


// This sort algorithm performs better for small data set
func BenchmarkBubbleSort100K(b *testing.B) {
	// Init
	elements := sort.GetElements(100000)

	// Execution
	for i :=0; i < b.N; i++ {
		Sort(elements)
	}
}
