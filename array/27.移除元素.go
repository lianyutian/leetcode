/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
package array

func removeElement(nums []int, val int) int {
	length := len(nums)
	left := 0

	for right := 0; right < length; right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}

	return left
}

// @lc code=end
