package leetcode

// 977 有序数组的平方
func sortedSquares(nums []int) []int {
	left := 0
	right, index := len(nums)-1, len(nums)-1
	newNums := make([]int, len(nums))
	for left <= right {
		l := nums[left]
		if nums[left] < 0 {
			l = -nums[left]
		}
		r := nums[right]
		if nums[right] < 0 {
			r = -nums[right]
		}
		if l > r {
			newNums[index] = l * l
			left++
		} else {
			newNums[index] = r * r
			right--
		}
		index--
	}
	return newNums
}
