/*
 * @lc app=leetcode.cn id=27 lang=javascript
 *
 * [27] 移除元素
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function(nums, val) {
    let length = nums.length
    let left = 0

    for (let right = 0; right < length; right++) {
        if (nums[right] != val) {
            nums[left++] = nums[right]
        }
    }

    return left
};
// @lc code=end

