package sorts

import "math"

// MergeSort Sort
func MergeSort(a []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		if p+1 < q {
			MergeSort(a, p, q)
		}
		if q+1 < r {
			MergeSort(a, q, r)
		}

		merge2(a, p, q, r)
	}

}

// mergeWithSentinel 带有限位器的合并算法
func mergeWithSentinel(a []int, p, q, r int) {
	ll := make([]int, q+1-p)
	for ii := p; ii < q+1; ii++ {
		ll[ii-p] = a[ii]
	}
	rr := make([]int, r-q)
	for ii := q + 1; ii < r+1; ii++ {
		rr[ii-q-1] = a[ii]
	}
	i := 0
	j := 0
	var sentinelMax = math.MaxInt32 //限位器，取一个最大值
	ll = append(ll, sentinelMax)
	rr = append(rr, sentinelMax)
	//fmt.Println("ll\t", ll)
	//fmt.Println("rr\t", rr)
	for k := p; k <= r; k++ {
		if ll[i] <= rr[j] {
			a[k] = ll[i]
			//fmt.Println("i\t", i, "\t", ll[i], "\t", rr[j])
			i++
		} else {
			a[k] = rr[j]
			//fmt.Println("j\t", j)
			j++
		}
	}
}
// merge2 不带限位器
func merge2(a []int, p, q, r int) {
	n1, n2 := q-p, r-q
	ll := make([]int, n1)
	for i := p; i < q; i++ {
		ll[i-p] = a[i]
	}
	rr := make([]int, n2)
	for i := q; i < r; i++ {
		rr[i-q] = a[i]
	}
	ii, jj := 0, 0
	for k := p; k < r; k++ {
		if ii < n1 && jj < n2 {
			if ll[ii] <= rr[jj] {
				a[k] = ll[ii]
				ii++
			} else {
				a[k] = rr[jj]
				jj++
			}
		} else if ii < n1 {
			a[k] = ll[ii]
			ii++
		} else if jj < n2 {
			a[k] = rr[jj]
			jj++
		}
	}
}

// MergeSort2 iterative merge sort
func MergeSort2(arr []int) {
	num := 1                   //number of elements in a group
	number := num * 2          // total number of elements of two adjacent group
	count := len(arr) / number // group count
	for count > 0 {
		for i := 0; i < count; i++ {
			b := number * i
			k1 := b               // start index of the first group
			end1 := b + num       // end index of the first group, not included
			if end1 >= len(arr) { // if already the end, break the loop
				break
			}

			arr0 := make([]int, number)
			k := 0
			k2 := end1          // start index of the second group
			end := b + number   // end index of the second group, not included
			if end > len(arr) { // the element index of the second group cannot be greater than the array length
				end = len(arr)
			}
			// merge algorithm
			for k1 < b+num && k2 < end {
				if arr[k1] <= arr[k2] {
					arr0[k] = arr[k1]
					k1++
				} else {
					arr0[k] = arr[k2]
					k2++
				}
				k++
			}
			if k1 == b+num {
				for k3 := k2; k3 < end; k3++ {
					arr0[k] = arr[k3]
					k++
				}
			} else {
				for k3 := k1; k3 < b+num; k3++ {
					arr0[k] = arr[k3]
					k++
				}
			}
			for k4 := b; k4 < end; k4++ { // write to the original array after merging
				arr[k4] = arr0[k4-b]
			}

		}
		if count == 1 {
			break
		}
		// group merge
		if count%2 == 0 {
			count /= 2
		} else { // include the second group which has less elements than the first group
			count = count/2 + 1
		}
		// group elements doubled, may be less than doubled
		num = number
		number *= 2
	}

}
