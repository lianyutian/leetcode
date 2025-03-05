/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
    tempMap := make(map[int]int, 0)
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := tempMap[v]; !ok {
			tempMap[v] = 1
		}
	}
	for _, v := range nums2 {
		if _, ok := tempMap[v]; ok {
			res = append(res, v)
			delete(tempMap, v)
		}
	}
	return res
}
// @lc code=end

