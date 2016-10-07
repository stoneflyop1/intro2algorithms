package sorts

// QuickSort sort
func QuickSort(a []int, p, r int) {
	if p < r {
		i := partitionS(a, p, r)
		QuickSort(a, p, i-1)
		QuickSort(a, i+1, r)
	}
}

// SortK get k sequence index
func SortK(a []int, k int) int {
	n := len(a)
	if k < 0 || k >= n {
		panic("k < 0 || k >= n")
	}
	index := selection4K(a, 0, n-1, k)
	return a[index]
}

func selection4K(a []int, left, right, k int) int {
	if left == right {
		return left
	}
	mid := partitionS(a, left, right)
	if mid == right {
		return mid
	}
	diff := mid - left + 1
	if diff >= k {
		//fmt.Println("diff>k", left, mid, diff, k)
		return selection4K(a, left, mid, k)
	}
	//fmt.Println("diff<=k", mid+1, right, k-diff)
	return selection4K(a, mid+1, right, k-diff)
}

// j must begin with p
func partitionS(a []int, p, r int) int {
	i := p
	x := a[r]
	for j := p; j < r; j++ {
		if a[j] <= x {
			if i != j {
				a[j], a[i] = a[i], a[j]
			}
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}

// for quick sort
func partitionQ(a []int, p, r int) int {
	i := p - 1
	x := a[r]
	for j := p; j < r; j++ {
		if a[j] <= x {
			i++
			if i != j {
				a[j], a[i] = a[i], a[j]
			}

		}
	}
	i++
	a[i], a[r] = a[r], a[i]
	return i
}
