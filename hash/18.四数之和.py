#
# @lc app=leetcode.cn id=18 lang=python3
#
# [18] 四数之和
#


# @lc code=start
class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        res = list()
        l = len(nums)
        nums.sort()

        for left in range(l - 3):
            nl = nums[left]
            # 去重
            if left > 0 and nums[left - 1] == nl:
                continue
            for left2 in range(left + 1, l - 2):
                nl2 = nums[left2]
                # 去重
                if left2 > left + 1 and nums[left2 - 1] == nl2:
                    continue
                index, right = left2 + 1, l - 1
                while index < right:
                    ni, nr = nums[index], nums[right]
                    sum = nl + nl2 + ni + nr
                    if sum == target:
                        res.append([nl, nl2, ni, nr])
                        while index < right and nums[index] == ni:
                            index += 1
                        while index < right and nums[right] == nr:
                            right -= 1
                    elif sum < target:
                        index += 1
                    else:
                        right -= 1

        return res


# @lc code=end
