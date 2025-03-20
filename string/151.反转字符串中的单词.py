#
# @lc app=leetcode.cn id=151 lang=python3
#
# [151] 反转字符串中的单词
#


# @lc code=start
class Solution:
    def reverseWords(self, s: str) -> str:
        # split()函数能够自动忽略多余的空白字符
        words = s.split()
        words = words[::-1]
        return " ".join(words)


# @lc code=end
