/*
 * @lc app=leetcode.cn id=160 lang=javascript
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @return {ListNode}
 */
var getIntersectionNode = function(headA, headB) {
    let pA = headA
    let pB = headB

    while (pA != null && pB != null) {
        pA = pA.next;
        pB = pB.next;
    }

    if (pA == null) {
        pA = headB;
        while(pB != null) {
            pA = pA.next;
            pB = pB.next;
        }
        pB = headA;
    } else if (pB == null) {
        pB = headA;
        while(pA != null) {
            pA = pA.next;
            pB = pB.next;
        }
        pA = headB;
    }

    while (pA != null && pB != null) {
        if (pA == pB) {
            return pA;
        }

        pA = pA.next;
        pB = pB.next;
    }

    return null;
};
// @lc code=end

