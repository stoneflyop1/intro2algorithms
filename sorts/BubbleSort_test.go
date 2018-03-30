package sorts

import "testing"

func TestBubbleSort(t *testing.T) {
	arr3 := []int{5, 2, 4, 7, 1, 3, 2, 6}
	BubbleSort(arr3)
	for i := 0; i < len(arr3)-1; i++ {
		if (arr3[i] > arr3[i+1]) {
			t.Errorf("BubbleSort order is incorrect, got: %v", arr3)
		}
	}
}