package sorts

// InsertionSort recursive
func InsertionSort(arr []int, endIndex int) {
	if endIndex > 0 {
		InsertionSort(arr, endIndex-1)
		insertSortSort(arr, endIndex)
	}

}

func insertSortSort(arr []int, endIndex int) {
	key := arr[endIndex]
	i := endIndex - 1
	for ; i >= 0 && key < arr[i]; i-- {
		arr[i+1] = arr[i]
	}
	arr[i+1] = key
}

// InsertionSortCount iterative with count
func InsertionSortCount(arr []int) int {
	count := 0
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		j := i - 1
		count++
		for j > 0 && arr[j] > v {
			count++
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = v
	}
	return count
}

// ComparisonSort comp
func ComparisonSort(arr []int) []int {
	countArr := make([]int, len(arr))
	s := make([]int, len(arr))
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] <= arr[j] { // if there is no equal sign, it's unstable
				countArr[j]++
			} else {
				countArr[i]++
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		s[countArr[i]] = arr[i]
	}
	return s
}
