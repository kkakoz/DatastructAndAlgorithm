package search

func BinarySearch(arr []int, n int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == n {
			return mid
		} else if n > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
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
