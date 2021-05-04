package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getElements(n int) []int {
	result := make([]int, n)

	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}

	return result
}

func TestBubbleSort(t *testing.T) {
	elements := []int{9, 8, 7, 6}

	elements = BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 4, len(elements))
	assert.EqualValues(t, 6, elements[0])
	assert.EqualValues(t, 7, elements[1])
	assert.EqualValues(t, 8, elements[2])
	assert.EqualValues(t, 9, elements[3])
}

func TestBubbleSortNilSlice(t *testing.T) {
	BubbleSort(nil)
}

func TestGetElements(t *testing.T) {
	elements := getElements(4)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 4, len(elements))
	assert.EqualValues(t, 3, elements[0])
	assert.EqualValues(t, 2, elements[1])
	assert.EqualValues(t, 1, elements[2])
	assert.EqualValues(t, 0, elements[3])
}

func BenchmarkBubbleSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}

func BenchmarkSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}

func BenchmarkSort100000(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}
