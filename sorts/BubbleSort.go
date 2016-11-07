package sorts

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// BubbleSort sort
// 1. for i <- 1 to length[arr]
// 2.     do for j <- length[arr] downto i+1
// 3.         do if arr[j] < arr[j-1]
// 4.             then exchange arr[j] <-> arr[j-1]
func BubbleSort2(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
