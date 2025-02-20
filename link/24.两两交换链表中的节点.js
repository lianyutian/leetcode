/*
 * @lc app=leetcode.cn id=24 lang=javascript
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var swapPairs = function(head) {
    let dummyNode = new ListNode()
    dummyNode.next = head
    let pre = dummyNode
    let cur = head

    while (cur != null && cur.next != null) {
        let next = cur.next
        let nextnext = next.next

        pre.next = next
        next.next = cur
        cur.next = nextnext

        pre = cur
        cur = nextnext
    }

    return dummyNode.next
};
// @lc code=end

