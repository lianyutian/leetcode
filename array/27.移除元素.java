/*
 * @lc app=leetcode.cn id=27 lang=java
 *
 * [27] 移除元素
 */

// @lc code=start
class Solution {
    public int removeElement(int[] nums, int val) {
        int slow = 0;
        int res = 0;

        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != val) {
                res++; 
                nums[slow++] = nums[i];
            }

        }

        return res;
    }
}
// @lc code=end

