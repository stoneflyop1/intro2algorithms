package sorts

// QuickSort sort
func QuickSort(a []int, p, r int) {
	if p < r {
		i := partitionHoare(a, p, r)
		QuickSort(a, p, i-1)
		QuickSort(a, i+1, r)
	}
}

// SortK get kth ordered element, kth will be the final index,
// the array will be partitioned for kth element
func SortK(a []int, k int) int {
	n := len(a)
	if k < 0 || k >= n {
		panic("k < 0 || k >= n")
	}
	index := selection4K(a, 0, n-1, k)
	return a[index]
}

// selection4K select the kth order index of the array
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

// partitionS LomutoPartition for selection
func partitionS(a []int, p, r int) int {
	i := p
	x := a[r]
	for j := p; j < r; j++ { // j must begin with p
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

// partitionQ LomutoPartition for quick sort
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

// partitionHoare partition from both ends
func partitionHoare(a []int, l, r int) int {
	p := a[l]
	i := l
	j := r + 1
	for i < j {
		i++
		for i < r && a[i] < p {
			i++
		}
		j--
		for j > l && a[j] > p {
			j--
		}
		a[i], a[j] = a[j], a[i]
	}
	a[i], a[j] = a[j], a[i]
	a[l], a[j] = a[j], a[l]
	return j
}
