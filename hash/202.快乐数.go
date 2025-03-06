/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start

func isHappy(n int) bool {
	// 循环终止条件 （无限循环，那么也就是说求和的过程中，sum会重复出现，这对解题很重要！）
	tempMap := make(map[int]bool)

	for n != 1 && !tempMap[n] {
		n, tempMap[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

// @lc code=end
