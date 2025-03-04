#
# @lc app=leetcode.cn id=59 lang=python3
#
# [59] 螺旋矩阵 II
#

# @lc code=start
class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        left, top, right, bottom = 0, 0, n - 1, n - 1
        res = [[0] * n for _ in range(n)]
        k = 1
        while k < (n * n):
            for i in range(left, right):
                res[top][i] = k
                k += 1
            for i in range(top, bottom):
                res[i][right] = k
                k += 1
            for i in range(right, left, -1):
                res[bottom][i] = k
                k += 1
            for i in range(bottom, top, -1):
                res[i][left] = k
                k += 1

            left += 1
            right -= 1
            bottom -= 1
            top += 1
        
        res[left][right] = n * n

        return res

# @lc code=end

