#
# @lc app=leetcode.cn id=637 lang=python3
#
# [637] 二叉树的层平均值
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from collections import deque


class Solution:
    def averageOfLevels(self, root: Optional[TreeNode]) -> List[float]:
        res = []
        queue = deque([root])

        while queue:
            size = len(queue)

            level = []

            for i in range(size):
                node = queue.popleft()
                level.append(node.val)
                if node.left:
                    queue.append(node.left)

                if node.right:
                    queue.append(node.right)

            res.append(sum(level) / len(level))
        return res


# @lc code=end
