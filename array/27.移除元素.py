#
# @lc app=leetcode.cn id=27 lang=python3
#
# [27] 移除元素
#

# @lc code=start
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        # 双指针
        # 1. i指针指向当前元素，j指针指向下一个元素
        # 2. 如果nums[i] == val, 则将nums[i]与nums[j]交换
        # 3. 如果nums[i] != val, 则i += 1
        # 4. 重复2-3步骤，直到i指针到达数组末尾
        i = 0
        n = len(nums)
        for j in range(n):
            if nums[j] != val:
                nums[i] = nums[j]
                i += 1
        return i
        
# @lc code=end
