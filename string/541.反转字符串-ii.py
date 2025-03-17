#
# @lc app=leetcode.cn id=541 lang=python3
#
# [541] 反转字符串 II
#


# @lc code=start
def rever(s: List[str]) -> List[str]:
    left = 0
    right = len(s) - 1
    while left <= right:
        s[left], s[right] = s[right], s[left]
        left += 1
        right -= 1
    return s


class Solution:

    def reverseStr(self, s: str, k: int) -> str:
        l = len(s)
        res = list(s)

        for i in range(0, l, 2 * k):
            res[i : i + k] = rever(res[i : i + k])
        return "".join(res)


# @lc code=end
