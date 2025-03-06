#
# @lc app=leetcode.cn id=202 lang=python3
#
# [202] 快乐数
#


# @lc code=start
class Solution:

    def isHappy(self, n: int) -> bool:
        list = []

        def getSum(n: int) -> int:
            sum = 0
            while n > 0:
                sum += (n % 10) ** 2
                n //= 10
            return sum

        while n != 1 and (n not in list):
            list.append(n)
            n = getSum(n)
        return n == 1


# @lc code=end
