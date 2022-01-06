package leetcode

// 704二分查找
func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return binarySearch(nums, start, mid-1, target)
	}
	return binarySearch(nums, mid+1, end, target)
}
