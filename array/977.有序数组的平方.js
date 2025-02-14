/*
 * @lc app=leetcode.cn id=977 lang=javascript
 *
 * [977] 有序数组的平方
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortedSquares = function(nums) {
    let l = nums.length
    let left = 0
    let right = l - 1
    let res = []
    let i = l - 1

    while (left <= right) {
        squaresLeft = nums[left] * nums[left]
        squaresRight = nums[right] * nums[right]

        if (squaresLeft < squaresRight) {
            res[i--] = squaresRight
            right--
        } else {
            res[i--] = squaresLeft
            left++
        }
    }
    
    return res
};
// @lc code=end

