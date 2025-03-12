#
# @lc app=leetcode.cn id=15 lang=python3
#
# [15] 三数之和
#


# @lc code=start
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        l = len(nums)
        res = list()

        for left, num in enumerate(nums):
            # 如果排序后的元组第一个元素已经大于0，后续怎么组合都不可能等于0
            if num > 0:
                break
            # 第一个元素去重
            if left > 0 and num == nums[left - 1]:
                continue
            index, right = left + 1, l - 1
            while index < right:
                ni, nr = nums[index], nums[right]
                sum = num + ni + nr
                if sum == 0:
                    res.append([num, ni, nr])
                    # 去重第二个元素
                    while index < right and ni == nums[index]:
                        index += 1
                    # 去重第三个元素
                    while index < right and nr == nums[right]:
                        right -= 1
                elif sum < 0:
                    index += 1
                else:
                    right -= 1

        return res


# @lc code=end
