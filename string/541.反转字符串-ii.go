/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start

func reverseStr(s string, k int) string {
	sb := []byte(s)
	l := len(sb)
	for i := 0; i < l; i += 2 * k {
		if i+k <= l {
			rever(sb[i : i+k])
		} else {
			rever(sb[i:l])
		}
	}
	return string(sb)
}

func rever(b []byte) {
	left := 0
	right := len(b) - 1

	for left <= right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

// @lc code=end
