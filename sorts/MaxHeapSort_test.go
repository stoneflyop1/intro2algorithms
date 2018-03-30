package sorts

import "testing"

func TestMaxHeapSort(t *testing.T) {
	arr := []int{5, 2, 4, 7, 1, 3, 2, 6}
	MaxHeapSort(arr)
	for i:=0;i<len(arr)-1;i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("MaxHeapsort order is incorrect, got %v", arr)
		}
	}
}