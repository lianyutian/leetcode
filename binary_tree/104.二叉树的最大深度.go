/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	queue := make([]*TreeNode, 0)
// 	queue = append(queue, root)
// 	depth := 0
// 	for len(queue) > 0 {
// 		size := len(queue)
// 		for i := 0; i < size; i++ {
// 			node := queue[0]
// 			queue = queue[1:]
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 		}
// 		depth++
// 	}
// 	return depth
// }
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	return getDepth(root, 1, res)

}

func getDepth(root *TreeNode, depth int, res int) int {
	if depth > res {
		res = depth
	}
	if root.Left == nil && root.Right == nil {
		return res
	}
	if root.Left != nil {
		depth++
		res = getDepth(root.Left, depth, res)
		depth--
	}
	if root.Right != nil {
		depth++
		res = getDepth(root.Right, depth, res)
		depth--
	}
	return res
}

// @lc code=end

