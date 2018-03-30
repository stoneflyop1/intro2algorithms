package sorts

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 2, 4, 7, 1, 3, 2, 6}
	MergeSort(arr, 0, len(arr))
	for i:=0;i<len(arr)-1;i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("MergeSort order is incorrect, got %v", arr)
		}
	}
}

func TestMergeSort2(t *testing.T) {
	arr := []int{5, 2, 4, 7, 1, 3, 2, 6}
	MergeSort2(arr)
	for i:=0;i<len(arr)-1;i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("MergeSort2 order is incorrect, got %v", arr)
		}
	}
}