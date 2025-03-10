#
# @lc app=leetcode.cn id=1 lang=python3
#
# [1] 两数之和
#


# @lc code=start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        map = dict()
        for i, num in enumerate(num):
            if target - num in map:
                return [i, map[target - num]]
            else:
                map[num] = i
        return []


# @lc code=end
