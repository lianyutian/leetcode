/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start

import (
	"strings"
)

func reverseWords(s string) string {
	res := make([]string, 0)
	splitS := strings.Split(s, " ")
	for _, str := range splitS {
		if str == "" {
			continue
		}
		res = append(res, str)
	}
	return reverse(res)
}

func reverse(s []string) string {
	left := 0
	right := len(s) - 1
	for left <= right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return strings.Join(s[:], " ")
}

// @lc code=end
