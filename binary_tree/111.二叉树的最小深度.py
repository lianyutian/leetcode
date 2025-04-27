#
# @lc app=leetcode.cn id=111 lang=python3
#
# [111] 二叉树的最小深度
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
    def minDepth(self, root: Optional[TreeNode]) -> int:
        res = 0
        if not root:
            return res
        queue = deque([root])
        flag = True
        while queue and flag:
            size = len(queue)
            res += 1
            for i in range(size):
                node = queue.popleft()
                if not node.left and not node.right:
                    flag = False
                    break
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        return res


# @lc code=end
