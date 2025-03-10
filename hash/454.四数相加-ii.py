#
# @lc app=leetcode.cn id=454 lang=python3
#
# [454] 四数相加 II
#


# @lc code=start
class Solution:
    def fourSumCount(
        self, nums1: List[int], nums2: List[int], nums3: List[int], nums4: List[int]
    ) -> int:
        map = dict()
        for num1 in nums1:
            for num2 in nums2:
                sum = num1 + num2
                if sum in map:
                    map[sum] += 1
                else:
                    map[sum] = 1

        res = 0
        for num3 in nums3:
            for num4 in nums4:
                sum = num3 + num4
                if 0 - sum in map:
                    res += map[0 - sum]
        return res


# @lc code=end
