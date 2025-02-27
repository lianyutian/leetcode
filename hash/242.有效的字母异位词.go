/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	record := [26]int{}
	for _, v := range s {
		record[v-'a']++
	}
	for _, v := range t {
		record[v-'a']--
	}
	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

