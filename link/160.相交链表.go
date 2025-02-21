/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB

	for pA != nil && pB != nil {
		pA = pA.Next
		pB = pB.Next
	}

	if pA == nil {
		pA = headB
		for pB != nil {
			pA = pA.Next
			pB = pB.Next
		}
		pB = headA
	} else if pB == nil {
		pB = headA
		for pA != nil {
			pA = pA.Next
			pB = pB.Next
		}
		pA = headB
	}

	for pA != nil && pB != nil {
		if pA == pB {
			return pA
		}

		pA = pA.Next
		pB = pB.Next
	}

	return nil
}

// @lc code=end

