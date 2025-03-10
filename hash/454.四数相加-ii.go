/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	map1 := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			map1[num1+num2]++
		}
	}

	res := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			sum := num3 + num4
			if v, ok := map1[0-sum]; ok {
				res += v
			}
		}
	}

	return res
}

// @lc code=end

