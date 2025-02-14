/*
 * @lc app=leetcode.cn id=209 lang=javascript
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
/**
 * @param {number} target
 * @param {number[]} nums
 * @return {number}
 */
var minSubArrayLen = function(target, nums) {
    let len = nums.length
    let left = 0
    let sum = 0
    let res = Number.MAX_VALUE

    for (let right = 0; right < len; right++) {
        sum += nums[right]

        while (sum >= target) {
            let temp = right - left + 1
            if (temp < res) {
                res = temp
            }
            sum -= nums[left]
            left++
        }
    }

    return res == Number.MAX_VALUE ? 0 : res
};
// @lc code=end

