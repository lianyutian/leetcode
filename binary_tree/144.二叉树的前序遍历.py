#
# @lc app=leetcode.cn id=144 lang=python3
#
# [144] 二叉树的前序遍历
#


# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        res = []

        def dfs(root: Optional[TreeNode]) -> List[int]:
            if root is None:
                return res
            res.append(root.val)
            dfs(root.left)
            dfs(root.right)
            return res

        dfs(root)

        return res


# @lc code=end
