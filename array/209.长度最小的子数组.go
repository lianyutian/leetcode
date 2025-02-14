/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	left := 0
	res := math.MaxInt
	sum := 0

	for right := 0; right < l; right++ {
		sum += nums[right]
		for sum >= target {
			temp := right - left + 1
			if temp < res {
				res = temp
			}
			sum -= nums[left]
			left++
		}
	}

	if res == math.MaxInt {
		res = 0
	}

	return res
}

// @lc code=end

