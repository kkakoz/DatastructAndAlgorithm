package search

import "constraints"

func BinarySearch[T constraints.Ordered](arr []T, n T) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == n {
			return mid
		}
		if n > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func binarySearchModel(arr []int, target int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	start, end := 0, len(arr)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if arr[mid] == target {
			end = mid
		} else if arr[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if arr[start] == target {
		return start
	}
	if arr[end] == target {
		return end
	}
	return -1
}

func BinarySearchRecursion(arr []int, n int) int {
	return binarySearchRecursion(arr, n, 0, len(arr)-1)
}

func binarySearchRecursion(arr []int, n, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == n {
		return mid
	}
	if n > arr[mid] {
		return binarySearchRecursion(arr, n, mid+1, right)
	}
	return binarySearchRecursion(arr, n, left, mid-1)
}
