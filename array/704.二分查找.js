/*
 * @lc app=leetcode.cn id=704 lang=javascript
 *
 * [704] 二分查找
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
    let len = nums.length
    let left = 0
    let right = len - 1

    while (left <= right) {
        let mid = Math.floor((right - left) / 2) + left
        if (target === nums[mid]) {
            return mid
        } else if (target > nums[mid]) {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1
};
// @lc code=end

