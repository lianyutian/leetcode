/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
package array

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		midValue := nums[mid]
		if midValue == target {
			return mid
		} else if target < midValue {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// @lc code=end
