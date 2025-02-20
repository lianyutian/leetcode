/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	pre := dummyNode
	cur := head

	for cur != nil && cur.Next != nil {
		next := cur.Next
		nextnext := next.Next

		pre.Next = next
		next.Next = cur
		cur.Next = nextnext

		pre = cur
		cur = nextnext
	}

	return dummyNode.Next
}

// @lc code=end

