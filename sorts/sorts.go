package sorts

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// TestSorts test sort algorithms
func Tests() {
	fmt.Println("*****************Sorting******************")
	arr := []int{5, 2, 4, 6, 1, 3}
	fmt.Println("-----------SortInt--------------")
	fmt.Println(arr)
	SortInt(arr)
	fmt.Println(arr)

	i := LinearSearch(arr, 3)
	fmt.Println(i)

	a := []int{1, 0, 0, 1, 1, 1, 0}
	b := []int{1, 1, 1, 1, 0, 1, 1}
	c := BinaryAdd(a, b)
	n := len(a) - 1
	aa := 0
	bb := 0
	ii := 0
	for ; n >= 0; n-- {
		p := int(math.Pow(2, float64(ii)))
		aa += a[n] * p
		bb += b[n] * p
		ii++
	}
	n0 := len(c) - 1
	cc := 0
	ii = 0
	for ; n0 >= 0; n0-- {
		p := int(math.Pow(2, float64(ii)))
		cc += c[n0] * p
		ii++
	}
	fmt.Println(aa, " ", bb, " ", cc, " ", c)

	fmt.Println("--------------selection sort-----------------")

	arr2 := []int{3, 6, 1, 9, 4}
	fmt.Println(arr2)
	SelectionSort(arr2)
	fmt.Println(arr2)

	fmt.Println("--------------merge sort--------------")
	arr3 := []int{5, 2, 4, 7, 1, 3, 2, 6}
	fmt.Println(arr3)
	MergeSort(arr3, 0, len(arr3))
	fmt.Println(arr3)

	fmt.Println("--------------insert sort--------------")
	arr3 = []int{5, 2, 4, 7, 1, 3, 2, 6}
	fmt.Println(arr3)
	InsertionSort(arr3, len(arr3)-1)
	fmt.Println(arr3)

	fmt.Println("--------------bubble sort--------------")
	arr3 = []int{5, 2, 4, 7, 1, 3, 2, 6}
	fmt.Println(arr3)
	BubbleSort(arr3)
	fmt.Println(arr3)

	fmt.Println("--------------max heap sort--------------")
	arr3 = []int{5, 7, 4, 2, 1, 3, 2, 6}
	fmt.Println(arr3)
	MaxHeapSort(arr3)
	fmt.Println(arr3)

	fmt.Println("--------------quick sort--------------")
	arr3 = []int{5, 7, 4, 2, 1, 3, 2, 6}
	fmt.Println(arr3)
	QuickSort(arr3, 0, len(arr3)-1)
	fmt.Println(arr3)

	fmt.Println("--------------sort k--------------")
	arr3 = []int{4, 7, 0, 2, 1, 3, 9, 6}
	fmt.Println(arr3)
	k := SortK(arr3, 5)
	fmt.Println(arr3)
	fmt.Println(k)

	fmt.Println("--------------dec2binary--------------")
	for _, n = range []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19} {
		bit := dec2binary(n)

		fmt.Println(n, bit)
	}

	fmt.Println("--------------compare sort--------------")
	arr0 := []int{60, 35, 81, 98, 14, 47}
	fmt.Println(arr0)
	arr00 := ComparisonSort(arr0)
	fmt.Println(arr00)

	fmt.Println("--------------exp--------------")
	rand.Seed(time.Now().UTC().UnixNano())
	startIndex := 1000
	step := 500
	for ii := 0; ii < 20; ii++ {
		aSize := startIndex + step*ii
		aArr := make([]int, aSize)
		for jj := 0; jj < aSize; jj++ {
			aArr[jj] = rand.Intn(aSize)
		}
		aCount := InsertionSortCount(aArr)
		fmt.Println(aSize, aCount)
	}

	fmt.Println("--------------find max Index--------------")
	fmt.Println(arr0)
	maxIndex := findMaxIndex(arr0, 0, len(arr0))
	fmt.Println(maxIndex, arr0[maxIndex])
	min, max := findMinMaxValue(arr0, 0, len(arr0))
	fmt.Println(min, max)

	fmt.Println("--------------merge sort 2--------------")
	arrarr := []int{6, 9, 7, 8, 5, 1, 3}
	fmt.Println(arrarr)
	MergeSort2(arrarr)
	fmt.Println(arrarr)

}

func findMaxIndex(arr []int, startIndex, endIndex int) int {

	mid := (startIndex + endIndex) / 2
	if startIndex == mid {
		return startIndex
	}
	index1 := findMaxIndex(arr, startIndex, mid)
	index2 := findMaxIndex(arr, mid, endIndex)
	if arr[index1] < arr[index2] {
		return index2
	}
	return index1
}

func findMinMaxValue(arr []int, startIndex, endIndex int) (int, int) {
	mid := (startIndex + endIndex) / 2
	if mid == startIndex {
		return arr[startIndex], arr[startIndex]
	} else if startIndex+1 == mid {
		if arr[startIndex] < arr[mid] {
			return arr[startIndex], arr[mid]
		}
		return arr[mid], arr[startIndex]
	}
	minL, maxL := findMinMaxValue(arr, startIndex, mid)
	minR, maxR := findMinMaxValue(arr, mid, endIndex)
	min := minL
	if min > minR {
		min = minR
	}
	max := maxL
	if max < maxR {
		max = maxR
	}
	return min, max
}

func pow2(i int) int {
	if i == 0 {
		return 1
	}
	v := 2
	for j := 1; j < i; j++ {
		v *= 2
	}
	return v
}

func dec2binary(n int) []int {
	var bit []int

	for n > 1 {
		i := 0
		n0 := n
		for n0 > 1 {
			//fmt.Println(n0)
			n0 = n0 / 2
			i++
		}
		n -= pow2(i)
		bit = append(bit, i)
	}
	if n == 1 {
		bit = append(bit, 0)
	}
	return bit
}

// SortInt InsertionSort
func SortInt(arr []int) {
	for j := 1; j < len(arr); j++ {
		a := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > a {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = a
	}
}

// LinearSearch (A, v)
//1. for i= 1 to A.length
//2.    if v == A[i]
//3.        return i
//4. return NIL
func LinearSearch(arr []int, v int) int {
	for i, v0 := range arr {
		if v0 == v {
			return i
		}
	}
	return -1
}

// BinaryAdd add two binary number
//1. for i = A.length downto 1
//2.     C[i+1] = C[i+1] + A[i] + B[i]
//3.     if C[i+1] >= 2
//4.         C[i] = C[i] + 1
//5.         C[i+1] = C[i+1] - 2
func BinaryAdd(a []int, b []int) []int {
	c := make([]int, len(a)+1)
	for i := len(a) - 1; i >= 0; i-- {
		cc := a[i] + b[i]
		// if cc >= 2 {
		//     c[i+1]++
		//     c[i] = cc - 2
		// } else {
		//     c[i] = cc
		// }
		c[i+1] += cc
		if c[i+1] >= 2 {
			c[i]++
			c[i+1] -= 2
		}
	}
	return c
}
