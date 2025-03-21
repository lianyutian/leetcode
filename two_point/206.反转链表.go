/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var left *ListNode
	cur := head

	for cur != nil {
		temp := cur.Next
		cur.Next = left

		left = cur
		cur = temp
	}

	return left
}

// @lc code=end

