#
# @lc app=leetcode.cn id=977 lang=python3
#
# [977] 有序数组的平方
#

# @lc code=start
class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        left, right = 0, len(nums) - 1
        res = [0] * (right + 1)
        k = right
        while left <= right:
            squareLeft = nums[left] ** 2
            squareRight = nums[right] ** 2
            
            if squareLeft < squareRight:
                res[k] = squareRight
                right -= 1
            else:
                res[k] = squareLeft
                left += 1
            k -= 1
        return res
# @lc code=end

