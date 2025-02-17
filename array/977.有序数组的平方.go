/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
package array

func sortedSquares(nums []int) []int {
	l := len(nums)
	left := 0
	right := l - 1
	res := make([]int, l)
	i := l - 1

	for left <= right {
		squareLeft := nums[left] * nums[left]
		squareRight := nums[right] * nums[right]

		if squareLeft < squareRight {
			res[i] = squareRight
			right--
		} else {
			res[i] = squareLeft
			left++
		}
		i--
	}

	return res
}

// @lc code=end
