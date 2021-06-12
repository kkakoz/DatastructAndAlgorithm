package sorts

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

func partition(arr []int, l, r int) int {
	val := arr[l]
	j := l
	i := l + 1
	for i <= r {
		if arr[i] > val {
			i++
		} else {
			j++
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}
