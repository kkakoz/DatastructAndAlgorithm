package sorts

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr) - 1)
}

func mergeSort(arr []int, l, r int) {
	if l == r {
		return
	}
	mid := (l + r) / 2
	mergeSort(arr, l, mid)
	mergeSort(arr, mid + 1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
	temp := make([]int, r - l + 1)
	copy(temp, arr[l:r+1])
	left := 0
	right := mid - l + 1
	for i := l; i <= r; i++ {
		if left + l > mid {
			arr[i] = temp[right]
			right++
		} else if right + l > r {
			arr[i] = temp[left]
			left++
		} else if temp[left] < temp[right] {
			arr[i] = temp[left]
			left++
		} else {
			arr[i] = temp[right]
			right++
		}
 	}
}

func MergeSort2(arr []int) {
	mergeSort2(arr, 0, len(arr) - 1)
}

func mergeSort2(arr []int, l, r int) {
	if l == r {
		return
	}
	mid := (l + r) / 2
	mergeSort(arr, l, mid)
	mergeSort(arr, mid + 1, r)
	if arr[mid] < arr[mid + 1] {
		merge(arr, l, mid, r)
	}
}
