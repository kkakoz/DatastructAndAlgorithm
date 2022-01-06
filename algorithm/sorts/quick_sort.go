package sorts

import "constraints"

func QuickSort[T constraints.Ordered](arr []T) {
	quickSort2(arr, 0, len(arr)-1)
}

func quickSort[T constraints.Ordered](arr []T, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

func partition[T constraints.Ordered](arr []T, l, r int) int {
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

func quickSort2[T constraints.Ordered](arr []T, start, end int) {
	if start >= end {
		return
	}
	left, right := start, end
	pivot := arr[(start+end)/2]

	// 小于等于 ,不然的话[1,2]的区间接下来调用还是[1,2]
	for left <= right {
		// arr[left] < pivot
		for left <= right && arr[left] < pivot {
			left++
		}
		for left <= right && arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	quickSort2(arr, start, right)
	quickSort2(arr, left, end)
}
