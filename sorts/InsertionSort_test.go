package sorts


import "testing"

func TestInsertionSort(t *testing.T) {
	arr := []int{5, 2, 4, 7, 1, 3, 2, 6}
	InsertionSort(arr, len(arr)-1)
	for i:=0;i<len(arr)-1;i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("InsertionSort Order is incorrect, got %v", arr)
		}
	}
}