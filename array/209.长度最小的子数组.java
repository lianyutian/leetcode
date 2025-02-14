/*
 * @lc app=leetcode.cn id=209 lang=java
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
class Solution {
    public int minSubArrayLen(int target, int[] nums) {
        int left = 0;
        int l = nums.length;
        int sum = 0;
        int res = Integer.MAX_VALUE;

        for (int right = 0; right < l; right++) {
            sum += nums[right];
            while (sum >= target) {
                int length = right - left + 1;
                if (length < res) {
                    res = length;
                }
                
                sum -= nums[left];
                left++;
            }
        }

        return res == Integer.MAX_VALUE ? 0 : res;
    }
}
// @lc code=end

