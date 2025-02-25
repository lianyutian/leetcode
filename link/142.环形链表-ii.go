/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	index := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// 存在环
		if slow == fast {
			for index != fast {
				index = index.Next
				fast = fast.Next
			}
			return index
		}
	}

	return nil
}

// @lc code=end

