package sorts

// MaxHeapSort Sort
// 0. heapSize = len(a)
// 1. buildMaxHeap(a, heapSize)
// 2. for i = len(a) downto 2
// 3.     exchange a[1] with a[i]
// 4.     heapSize--
// 5.     maxHeapify(a, 1, heapSize)
func MaxHeapSort(a []int) {
	size := len(a)
	heapSize := size - 1
	buildMaxHeap(a, heapSize)
	for i := heapSize; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapSize--
		maxHeapify(a, 0, heapSize)
	}
}

func parentHeap(i int) int {
	return i / 2
}

func leftHeap(i int) int {
	return (i+1)*2 - 1
}

func rightHeap(i int) int {
	return (i + 1) * 2
}

// maxHeapify maintain max heap
// 0. maxHeapify(a, largest, heapSize)
// 1. l = leftHeap(i)
// 2. r = rightHeap(i)
// 3. if l <= heapSize and a[l] > a[i]
// 4.    largest = l
// 5. else largest = i
// 6. if r <= heapSize and a[r] > a[largest]
// 7.     largest = r
// 8. if largest != i
// 9.     exchange a[i] with a[largest]
// 10.    maxHeapify(a, largest, heapSize)
func maxHeapify1(a []int, i int, heapSize int) {
	l := leftHeap(i)
	r := rightHeap(i)
	largest := i
	if l <= heapSize && a[l] > a[i] {
		largest = l
	}
	if r <= heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}

func maxHeapify(a []int, i int, heapSize int) {
	for {
		l := leftHeap(i)
		r := rightHeap(i)
		largest := i
		if l <= heapSize && a[l] > a[i] {
			largest = l
		}
		if r <= heapSize && a[r] > a[largest] {
			largest = r
		}

		if largest == i {
			break
		}
		a[i], a[largest] = a[largest], a[i]
		i = largest

	}

}

// buildMaxHeap
// 0. heapSize = len(a)
// 1. for i = len(a)/2 downto 1
// 2.     maxHeapify(a, i, heapSize)
func buildMaxHeap(a []int, heapSize int) {
	for i := len(a) / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}
