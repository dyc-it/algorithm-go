package sort

import (
	"testing"
)

var slice = []int{5, 3, 9, 0, 7, 8, 1, 6, 4, 2}
var supposedSlice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestBubble_sort(t *testing.T) {
	bubble_sort(slice)
	for i := range slice {
		if slice[i] != supposedSlice[i] {
			t.Errorf("Testing fails")
		}
	}
}

func BenchmarkBubble_sort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubble_sort(slice)
	}
}
