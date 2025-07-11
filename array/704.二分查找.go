/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start

func search(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1

	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// @lc code=end
