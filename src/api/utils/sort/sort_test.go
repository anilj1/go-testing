package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T)  {
	if privateConst != "private" {
		t.Error("privateConst should be 'private")
	}
}

// Using the Assert API instead of t.Error()
func TestBubbleSortAscending(t *testing.T) {
	// Init
	elements := GetElements(10)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0], "First element of unsorted array should 9")
	assert.EqualValues(t, 0, elements[len(elements)-1], "Last element of unsorted array should 0")

	// Execution
	BubbleSort(elements)

	// Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0], "First element of sorted array should 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "Last element of sorted array should 9")

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
	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0], "First element of unsorted array should 9")
	assert.EqualValues(t, 0, elements[len(elements)-1], "Last element of unsorted array should 0")

	// Execution
	Sort(elements)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 0, elements[0], "First element of sorted array should 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "Last element of sorted array should 9")

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

