/*
 * @lc app=leetcode.cn id=202 lang=javascript
 *
 * [202] 快乐数
 */

// @lc code=start
/**
 * @param {number} n
 * @return {boolean}
 */
var isHappy = function(n) {
    const set = new Set()
    while (n != 1 && !set.has(n)) {
        set.add(n)
        n = getSum(n)
    }
    return n === 1
};

var getSum = function(n) {
    let sum = 0
    while (n > 0) {
        sum += (n % 10) ** 2
        n = Math.floor(n / 10)
    }
    return sum
}
// @lc code=end

