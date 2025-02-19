/*
 * @lc app=leetcode.cn id=203 lang=javascript
 *
 * [203] 移除链表元素
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
 * @param {number} val
 * @return {ListNode}
 */
var removeElements = function(head, val) {
    let dummyNode = new ListNode(-1, head)
    let pre = dummyNode
    let cur = head

    while (cur != undefined) {
        if (cur.val === val) {
            pre.next = cur.next
        } else {
            pre = cur
        }
        cur = cur.next
    }

    return dummyNode.next
};
// @lc code=end

