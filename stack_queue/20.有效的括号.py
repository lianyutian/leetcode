#
# @lc app=leetcode.cn id=20 lang=python3
#
# [20] 有效的括号
#


# @lc code=start
class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        m = {}
        m[")"] = "("
        m["]"] = "["
        m["}"] = "{"

        for c in s:
            if c == "(" or c == "{" or c == "[":
                stack.append(c)
            else:
                if len(stack) == 0:
                    return False
                pop = stack.pop()
                if pop != m[c]:
                    return False
        return len(stack) == 0


# @lc code=end
