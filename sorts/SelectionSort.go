package sorts

// SelectionSort arr
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if min > arr[j] {
				min = arr[j]
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
