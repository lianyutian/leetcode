/*
 * @lc app=leetcode.cn id=59 lang=javascript
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
/**
 * @param {number} n
 * @return {number[][]}
 */
var generateMatrix = function(n) {
    let top = 0
    let right = n - 1
    let bottom = n - 1
    let left = 0
    let res = Array.from({length: n}, () => Array(n).fill(0));
    let k = 1

    while (k < (n * n)) {
        // 左到右
        for (let i = left; i < right; i++) {
            res[top][i] = k
            k++
        }
        // 上到下
        for (let i = top; i < bottom; i++) {
            res[i][right] = k
            k++
        }
        // 右到左
        for (let i = right; i > left; i--) {
            res[bottom][i] = k
            k++
        }
        // 下到上
        for (let i = bottom; i > top; i--) {
            res[i][left] = k
            k++
        }

        // 向内缩
        top++
        right--
        bottom--
        left++
    }

    // 中心点
    res[left][right] = n * n

    return res

};
// @lc code=end

