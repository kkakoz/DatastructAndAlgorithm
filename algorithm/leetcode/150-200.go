package leetcode

import "fmt"

// 167 两组织和 2 输入有序数组
func twoSum2(numbers []int, target int) []int {
	res := make([]int, 0, 2)
	for i, j := 0, len(numbers)-1; i < j; {
		cur := numbers[i] + numbers[j]
		if cur == target {
			res = append(res, i+1, j+1)
			return res
		}
		if cur < target {
			i++
		}
		if cur > target {
			j--
		}
	}
	return res
}

// 189 轮转数组
func rotate(nums []int, k int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	length := len(nums)
	shiftPlace := k % length
	left := make([]int, shiftPlace)
	copy(left, nums[len(nums)-shiftPlace:])
	fmt.Println(left)
	for i := len(nums) - 1; i >= shiftPlace; i-- {
		nums[i] = nums[i-shiftPlace]
	}
	for i, v := range left {
		nums[i] = v
	}
}
