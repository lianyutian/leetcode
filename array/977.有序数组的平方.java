/*
 * @lc app=leetcode.cn id=977 lang=java
 *
 * [977] 有序数组的平方
 */

// @lc code=start
class Solution {
    public int[] sortedSquares(int[] nums) {
        int l = nums.length;
        int left = 0;
        int right = l - 1;
        int[] res = new int[l];
        int i = l - 1;
        
        while (left <= right) {
            int squaresLeft = nums[left] * nums[left];
            int squaresRight = nums[right] * nums[right];

            if (squaresLeft < squaresRight) {
                res[i--] = squaresRight;
                right--;
            } else {
                res[i--] = squaresLeft;
                left++;
            }
        }

        return res;
    }
}
// @lc code=end

