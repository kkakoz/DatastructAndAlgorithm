package leetcode

func isBadVersion(version int) bool {
	if version <= 4 {
		return false
	}
	return true
}

// 278. 第一个错误的版本
func firstBadVersion(n int) int {
	left, right := 0, n
	for left+1 < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	if !isBadVersion(right) {
		return right + 1
	}
	if !isBadVersion(left) {
		return left + 1
	}
	return 0
}
