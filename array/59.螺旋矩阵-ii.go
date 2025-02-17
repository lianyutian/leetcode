/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
package array

func generateMatrix(n int) [][]int {
	top := 0
	bottom := n - 1
	left := 0
	right := n - 1
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n) // Initialize inner slices
	}
	k := 1
	for k < (n * n) {
		for i := left; i < right; i++ {
			res[top][i] = k
			k++
		}
		for i := top; i < bottom; i++ {
			res[i][right] = k
			k++
		}
		for i := right; i > left; i-- {
			res[bottom][i] = k
			k++
		}
		for i := bottom; i > top; i-- {
			res[i][left] = k
			k++
		}
		top++
		right--
		bottom--
		left++
	}

	res[left][right] = n * n

	return res
}

// @lc code=end
