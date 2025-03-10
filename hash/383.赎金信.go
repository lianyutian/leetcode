/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	record := [26]int{}
	for _, c := range magazine {
		record[c-'a']++
	}

	for _, c := range ransomNote {
		record[c-'a']--
		if record[c-'a'] < 0 {
			return false
		}
	}

	return true
}

// @lc code=end

