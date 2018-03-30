package sorts

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{5, 2, 4, 7, 1, 3, 2, 6}
	QuickSort(arr, 0, len(arr)-1)
	for i:=0;i<len(arr)-1;i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("QuickSort order is incorrect, got %v", arr)
		}
	}
}

func TestSortK(t *testing.T) {
	arr3 := []int{4, 7, 0, 2, 1, 3, 9, 6}
	k := SortK(arr3, 5)
	if k != 6 {
		t.Errorf("SortK expect 6 , got %d", k)
	}
}