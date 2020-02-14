package sort

import (
	"testing"
)

func TestConstants(t *testing.T)  {
	if privateConst != "private" {
		t.Error("privateConst should be 'private")
	}
}

func TestBubbleSortAscending(t *testing.T) {
	// Init
	elements := GetElements(10)

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 0 {
		// t.Error does not stop executing remaining test case!
		t.Error("First element of sorted array should 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Last element of sorted array should 9")
	}
}

func TestSortLibraryAPI(t *testing.T) {
	// Init
	elements := GetElements(10)

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

func TestBubbleSortSorted(t *testing.T) {
	// Init
	elements := []int {9,8,7,6,5,4}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 4 {
		t.Error("First element of sorted array should 4")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Last element of sorted array should 9")
	}
}

func BenchmarkBubbleSort10(b *testing.B) {
	// Init
	elements := GetElements(10)

	// Execution
	for i :=0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort10(b *testing.B) {
	// Init
	elements := GetElements(10)

	// Execution
	for i :=0; i < b.N; i++ {
		Sort(elements)
	}
}

// This sort algorithm performs better for small data set
func BenchmarkBubbleSort10K(b *testing.B) {
	// Init
	elements := GetElements(10000)

	// Execution
	for i :=0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

// This sort algorithm performs better for large data set
func BenchmarkSort10K(b *testing.B) {
	// Init
	elements := GetElements(10000)

	// Execution
	for i :=0; i < b.N; i++ {
		Sort(elements)
	}
}

// This sort algorithm performs better for small data set
func BenchmarkBubbleSort100K(b *testing.B) {
	// Init
	elements := GetElements(100000)

	// Execution
	for i :=0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

// This sort algorithm performs better for large data set
func BenchmarkSort100K(b *testing.B) {
	// Init
	elements := GetElements(100000)

	// Execution
	for i :=0; i < b.N; i++ {
		Sort(elements)
	}
}

