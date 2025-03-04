#
# @lc app=leetcode.cn id=209 lang=python3
#
# [209] 长度最小的子数组
#

# @lc code=start
class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        left = 0
        right = 0
        n = len(nums) - 1
        sum = 0
        res = sys.maxsize   

        while right <= n:
            sum += nums[right]
            while sum >= target:
                temp = right - left + 1
                res = min(res, temp) 
                sum -= nums[left]
                left += 1
            right += 1
        return res if res != sys.maxsize else 0
# @lc code=end

