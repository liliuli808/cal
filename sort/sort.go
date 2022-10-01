package sort

func QuickSort(q []int, l int, r int) []int {
	if l >= r {
		return nil
	}
	i, j := l, r-1
	k := (l + r) >> 1

	for {
		if q[i] >= q[k] {
			break
		}
		i++
	}

	for {
		if q[j] < q[k] {
			break
		}
		j--
	}

	if i < j {
		q[i] = q[j]
	}
	QuickSort(q, l, j)
	QuickSort(q, j+1, r)
	return q
}
