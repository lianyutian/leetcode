#
# @lc app=leetcode.cn id=349 lang=python3
#
# [349] 两个数组的交集
#

# @lc code=start
class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        tempSet = set()
        res = []

        for num in nums1:
            tempSet.add(num)

        for num in nums2:
            if num in tempSet:
                res.append(num)
                tempSet.discard(num)

        return res
# @lc code=end

