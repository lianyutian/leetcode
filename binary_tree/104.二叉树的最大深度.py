#
# @lc app=leetcode.cn id=104 lang=python3
#
# [104] 二叉树的最大深度
#


# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        res = 0

        def getDepth(node: TreeNode, depth: int):
            nonlocal res
            res = res if res > depth else depth
            if not node.left and not node.right:
                return
            if node.left:
                depth += 1
                getDepth(node.left, depth)
                depth -= 1  # 回朔
            if node.right:
                depth += 1
                getDepth(node.right, depth)
                depth -= 1  # 回朔
            return

        if not root:
            return res
        getDepth(root, 1)
        return res


# @lc code=end
