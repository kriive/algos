package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 6, 7, 8, 4, 5, 2, 3, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr = Bubble(arr)

	if !equals(arr, expected) {
		t.Errorf("sort.Bubble({9,6,7,8,4,5,2,3,1}) = %d, expected {1,2,3,4,5,6,7,8,9}", arr)
	}
}

func TestBubbleSortAlreadySorted(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr = Bubble(arr)

	if !equals(arr, expected) {
		t.Errorf("sort.Bubble({1,2,3,4,5,6,7,8,9}) = %d, expected {1,2,3,4,5,6,7,8,9}", arr)
	}
}

func equals(arr []int, arr2 []int) bool {
	if len(arr) != len(arr2) {
		return false
	}

	for i, v := range arr {
		if v != arr2[i] {
			return false
		}
	}
	return true
}
