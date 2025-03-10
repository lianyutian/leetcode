/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
// @lc code=start
func twoSum(nums []int, target int) []int {
	tempMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if preIndex, ok := tempMap[target-v]; ok {
			return []int{i, preIndex}
		} else {
			tempMap[v] = i
		}
	}
	return []int{}
}

// @lc code=end
